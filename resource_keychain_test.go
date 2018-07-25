package main

import (
	"testing"
)

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

// func TestQueryItem(t *testing.T) {
// 	items, err := queryItem("AirPort", "SSID name 1")

// 	if err != nil {
// 		t.Errorf("queryItem errored, got: %s", err)
// 	}

// 	item := items[0]
// 	fmt.Printf("Item: %#v\n", item)

// 	fmt.Printf("Service: %s\n", item.Service)
// 	fmt.Printf("Account: %s\n", item.Account)
// 	fmt.Printf("AccessGroup: %s\n", item.AccessGroup)
// 	fmt.Printf("Label: %s\n", item.Label)
// 	fmt.Printf("Description: %s\n", item.Description)
// 	fmt.Printf("Data: %s\n", item.Data)
// }
