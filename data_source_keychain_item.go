package main

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceKeychainItem() *schema.Resource {
	return &schema.Resource{
		Read: resourceKeychainRead,

		Schema: map[string]*schema.Schema{
			// Where
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of service - Where",
			},
			// Account
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name - Account",
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
			"synchronizable": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Should this sync across devices? (default is `true`)",
			},
			"accessible": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The lock setting (default is `true`)",
			},
		},
	}
}
