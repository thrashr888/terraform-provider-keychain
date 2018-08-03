---
layout: "keychain"
page_title: "keychain: keychain_wifi"
sidebar_current: "docs-keychain-datasource-keychain-wifi
description: |-
  Get a keychain wifi password by SSID / account name.
---

# keychain

Use this data source to get a [keychain item][1], specifically a wifi password.

## Example Usage

```hcl
data "keychain_wifi" "account-key" {
  account = "${var.service_username}"
}

provider "keychain_wifi" "my-service" {
  account = "${var.service_username} 2"
  data = "${data.keychain.account-key.data}"
}
```

## Attributes Reference

- `account` - The name of the WiFi's SSID. Ex. "ATTWIFI-btf678ads9"

- `data` - The password from the keychain.

[1]: https://developer.apple.com/documentation/security/keychain_services/keychain_items
