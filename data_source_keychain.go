package main

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceKeychain() *schema.Resource {
	return &schema.Resource{
		Read: resourceKeychainRead,

		Schema: map[string]*schema.Schema{
			// Where
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of service (default is `Airport`) - Where",
			},
			// Account | SSID
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name - Account / SSID",
			},
			// Password
			"data": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The password output - Password",
			},
		},
	}
}
