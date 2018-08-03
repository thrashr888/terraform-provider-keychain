---
layout: "keychain"
page_title: "keychain: keychain_wifi"
sidebar_current: "docs-keychain-resource-keychain-wifi"
description: |-
  Provides a macOS keychain resource specifically for WiFi passwords.
---

# keychain

Provides a macOS keychain resource specifically for WiFi passwords.

## Example Usage

```hcl
resource "keychain_wifi" "test_ssid_1" {
  account = "SSID Name 1"
  data    = "Wifi Password 1"
}
```

## Argument Reference

The following arguments are supported:

* `account` - (Required) The SSID name.
* `data` - (Required) The WiFi password.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the key. A combination of `service` and `account`.