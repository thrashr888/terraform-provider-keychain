package main

import "testing"

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
