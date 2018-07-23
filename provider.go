package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider provides a terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"keychain": resourceKeychain(),
		},
	}
}
