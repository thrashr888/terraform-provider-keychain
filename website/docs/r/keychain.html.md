---
layout: "keychain"
page_title: "keychain: keychain"
sidebar_current: "docs-keychain-resource-keychain"
description: |-
  Provides a macOS keychain resource.
---

# keychain

Provides a macOS keychain resource. This sits in front of a number of defined pools of origins and provides various options for geographically-aware load balancing. Note that the load balancing feature must be enabled in your Clouflare account before you can use this resource.

## Example Usage

```hcl
resource "keychain" "test_ssid_1" {
  label   = "SSID Name 1"
  account = "SSID Name 1"
  data    = "WiFi Password 1"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Optional) The key's service name or "Where".
* `account` - (Required) The key's account name or "Account".
* `label` - (Required) The key's label or "Name".
* `data` - (Required) The key's data or "password".
* `description` - (Optional) The key's description or "Kind".
* `synchronizable` - (Optional) The key's synchronizable setting.
* `accessible` - (Optional) The key's accessible setting.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the key. A combination of `service` and `account`.