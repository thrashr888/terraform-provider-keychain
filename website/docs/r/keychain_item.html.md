---
layout: "keychain"
page_title: "keychain: keychain_item"
sidebar_current: "docs-keychain-resource-keychain-item"
description: |-
  Provides a macOS keychain item resource.
---

# keychain

Provides a macOS keychain resource.

## Example Usage

```hcl
resource "keychain_item" "test_key_1" {
  service = "Key Service 1"
  label   = "Key Name 1"
  account = "Key Name 1"
  data    = "Key Password 1"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) The key's service name or "Where".
* `account` - (Required) The key's account name or "Account".
* `label` - (Required) The key's label or "Name".
* `data` - (Required) The key's data or "password".
* `description` - (Optional) The key's description or "Kind".
* `synchronizable` - (Optional) The key's synchronizable setting.
* `accessible` - (Optional) The key's accessible setting.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the key. A combination of `service` and `account`.