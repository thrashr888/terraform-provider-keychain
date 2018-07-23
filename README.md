# Terraform Provider Keychain

This provider is used to manage the local macOS Keychain. Ideal for syncing a set of application or wifi passwords.

## Example

```hcl
resource "keychain" "test-ssid" {
  service     = "AirPort"
  account     = "ssid name"
  label       = "ssid name"
  data        = "wifi password"
  description = "test description"
}
```

## Testing

```
$ go build -o terraform-provider-keychain
$ terraform init
$ terraform plan
```
