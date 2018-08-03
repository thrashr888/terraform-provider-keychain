package main

import (
	"testing"

	keychain "github.com/keybase/go-keychain"
)

var service = "TestService"
var account = "TestAccount"
var label = "Test Label"
var data = "Test Data"
var description = "Test Description"

func TestCreateID(t *testing.T) {
	id := createID("service", "account")
	match := "[\"service\",\"account\"]"
	if id != match {
		t.Errorf("id was incorrect, got: %s, want: %s.", id, match)
	}
}

func TestGetID(t *testing.T) {
	service, account := getID("[\"test-service\",\"test-account\"]")
	if service != "test-service" {
		t.Errorf("service was incorrect, got: %s, want: %s.", service, "test-service")
	}
	if account != "test-account" {
		t.Errorf("account was incorrect, got: %s, want: %s.", account, "test-account")
	}
}

func TestAddItem(t *testing.T) {
	err := addItem(service, account, label, data, description)

	if err == keychain.ErrorDuplicateItem {
		// Duplicate is ok
		t.Logf("Duplicate item at [%s %s]", service, account)
	} else if err != nil {
		t.Errorf("addItem errored, got: %s at [%s %s]", err, service, account)
	}
}

func TestQueryItem(t *testing.T) {
	items, err := queryItem(service, account)

	if err != nil {
		t.Errorf("queryItem errored, got: %s", err)
		return
	} else if len(items) < 1 {
		// Not found
		t.Errorf("Keychain item not found at [%s %s] => %s", service, account, items)
		return
	}

	item := items[0]
	t.Logf("Item found: %#v\n", item)

	if item.Service != service {
		t.Errorf("Service is not correct: expected %s got %s", service, item.Service)
	}
	if item.Account != account {
		t.Errorf("Account is not correct: expected %s got %s", account, item.Account)
	}
	if item.Label != label {
		t.Errorf("Label is not correct: expected %s got %s", label, item.Label)
	}
	if item.Description != description {
		t.Errorf("Description is not correct: expected %s got %s", description, item.Description)
	}
	if item.AccessGroup != "" {
		// AccessGroup is returned unset to ""
		t.Errorf("AccessGroup is not correct: expected %s got %s", "", item.AccessGroup)
	}
	if string(item.Data) != data {
		t.Errorf("Data is not correct: expected %s got %s", data, item.Data)
	}
}

func TestDeleteItem(t *testing.T) {
	err := deleteItem(service, account)

	if err != nil {
		t.Errorf("deleteItem errored, got: %s", err)
	}
}
