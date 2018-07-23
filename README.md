# Terraform Provider Keychain

This provider is used to manage the local macOS Keychain. Ideal for syncing a set of application or wifi passwords.

## Example

```hcl
resource "keychain" "test-ssid" {
  account = "ssid name"
  data    = "wifi password"
}
```

## Testing

```
$ go build -o terraform-provider-keychain
$ terraform init
$ terraform plan
```
