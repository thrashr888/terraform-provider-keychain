package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	keychain "github.com/keybase/go-keychain"
)

//
// Setup
//

func resourceKeychain() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeychainCreate,
		Read:   resourceKeychainRead,
		Update: resourceKeychainUpdate,
		Delete: resourceKeychainDelete,
		Exists: resourceKeychainExists,

		Schema: map[string]*schema.Schema{
			// Where
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AirPort",
				Description: "The type of service (default is `Airport`) - Where",
			},
			// Account | SSID
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name - Account / SSID",
			},
			"access-group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AirPort",
				Description: "The access group name (default is `Airport`)",
			},
			// Name | SSID
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The record's label (defaults to using the Account name) - Name / SSID",
			},
			// Password
			"data": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password - Password",
			},
			// Kind
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AirPort network password",
				Description: "A helpful description - Kind",
			},
			"synchronizable": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Should this sync across devices? (default is `true`)",
			},
			"accessible": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "The lock setting (default is `true`)",
			},
		},
	}
}

//
// CRUD
//

func resourceKeychainCreate(d *schema.ResourceData, meta interface{}) error {
	service := d.Get("service").(string)
	account := d.Get("account").(string)
	label := d.Get("label").(string)
	data := d.Get("data").(string)
	description := d.Get("description").(string)
	// TODO:
	// accessGroup := d.Get("access-group").(string)
	// accessible := d.Get("accessible").(bool)
	// synchronizable := d.Get("synchronizable").(bool)

	err := addItem(service, account, label, data, description)

	if err == keychain.ErrorDuplicateItem {
		// Duplicate
		d.SetId(createID(service, account))
		return nil
	}

	if err != nil {
		return err
	}

	d.SetId(createID(service, account))
	return nil
}

func resourceKeychainRead(d *schema.ResourceData, meta interface{}) error {

	// get by id or service & account
	id := d.Id()
	service := ""
	account := ""

	if id != "" {
		service, account = getID(id)
	} else {
		service = d.Get("service").(string)
		account = d.Get("account").(string)
	}

	results, err := queryItem(service, account)

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		// Error
		d.SetId("")
		return nil
	} else if len(results) != 1 {
		// Not found
		d.SetId("")
		return nil
	}

	obj := results[0]

	d.Set("service", obj.Service)
	d.Set("account", obj.Account)
	if obj.Label == "" {
		d.Set("label", obj.Account)
	} else {
		d.Set("label", obj.Label)
	}
	d.Set("data", obj.Data)
	d.Set("description", obj.Description)

	return nil
}

func resourceKeychainUpdate(d *schema.ResourceData, meta interface{}) error {
	err := deleteByID(d.Id())
	if err != nil {
		return err
	}

	err = resourceKeychainCreate(d, meta)
	if err != nil {
		return err
	}

	return nil
}

func resourceKeychainDelete(d *schema.ResourceData, meta interface{}) error {
	err := deleteByID(d.Id())

	if err != nil {
		return err
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceKeychainExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	results, err := queryByID(d.Id())

	// If the resource does not exist, inform Terraform
	if err != nil {
		// Error
		return false, err
	} else if len(results) != 1 {
		// Not found
		return false, nil
	}

	return true, nil
}

//
// Internal API
//

func addItem(service string, account string, label string, data string, description string) error {
	item := newItem(service, account, label, data, description)
	err := keychain.AddItem(item)
	return err
}

func queryItem(service string, account string) ([]keychain.QueryResult, error) {
	queryItem := newItem(service, account, "", "", "")
	queryItem.SetMatchLimit(keychain.MatchLimitOne)
	queryItem.SetReturnData(true)
	queryItem.SetReturnAttributes(true)
	return keychain.QueryItem(queryItem)
}

func queryByID(id string) ([]keychain.QueryResult, error) {
	service, account := getID(id)
	return queryItem(service, account)
}

func deleteItem(service string, account string) error {
	item := newItem(service, account, "", "", "")
	err := keychain.DeleteItem(item)
	return err
}

func deleteByID(id string) error {
	service, account := getID(id)
	return deleteItem(service, account)
}

//
// Helpers
//

func newItem(service string, account string, label string, data string, description string) keychain.Item {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetAccessGroup("AirPort")

	if label != "" {
		item.SetLabel(label)
	} else {
		item.SetLabel(account)
	}
	if data != "" {
		item.SetData([]byte(data))
	}
	if description != "" {
		item.SetDescription(description)
	}

	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	trustedApplications := []string{"/usr/libexec/airportd"}
	item.SetAccess(&keychain.Access{Label: "AirPort", TrustedApplications: trustedApplications})

	return item
}

//
// ID Generation
//

func createID(service string, account string) string {
	id := []string{service, account}
	serializedID, _ := json.Marshal(id)
	return string(serializedID)
}

func getID(id string) (string, string) {
	var parts []string

	if err := json.Unmarshal([]byte(id), &parts); err != nil {
		panic(err)
		fmt.Printf("Parts missing: %v", id)
	}

	return parts[0], parts[1]
}
