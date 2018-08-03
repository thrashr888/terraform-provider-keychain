package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

//
// Setup
//

func resourceKeychainWifi() *schema.Resource {
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
