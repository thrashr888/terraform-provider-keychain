package main

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceKeychainWifi() *schema.Resource {
	return &schema.Resource{
		Read: resourceKeychainRead,

		Schema: map[string]*schema.Schema{
			// Where
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AirPort",
				Description: "The type of service - Where",
			},
			// Account | SSID
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name - Account / SSID",
			},
			// Name
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The record's label (defaults to using the Account name) - Name",
			},
			// Password
			"data": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The password output - Password",
			},
			// Kind
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A helpful description - Kind",
			},
		},
	}
}
