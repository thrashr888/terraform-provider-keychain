package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider provides a terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"keychain_item": resourceKeychainItem(),
			"keychain_wifi": resourceKeychainWifi(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"keychain_item": dataSourceKeychainItem(),
			"keychain_wifi": dataSourceKeychainWifi(),
		},
	}
}
