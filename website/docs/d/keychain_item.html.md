---
layout: "keychain"
page_title: "keychain: keychain_item"
sidebar_current: "docs-keychain-datasource-keychain-item
description: |-
  Get a keychain item by service and account names.
---

# keychain

Use this data source to get a [keychain item][1], likely containing a password or secret key.

## Example Usage

```hcl
data "keychain_item" "account-key" {
  service = "My Service"
  account = "${var.service_username}"
}

provider "keychain_item" "my-service" {
  service = "My Second Service"
  account = "${var.service_username}"
  data = "${data.keychain.account-key.data}"
}
```

## Attributes Reference

- `service` - The name of the key's service. Ex. "AIM".

- `account` - The name of the key's account, or username. Ex. "myaccount@example.com"

- `data` - The password from the keychain.

[1]: https://developer.apple.com/documentation/security/keychain_services/keychain_items
