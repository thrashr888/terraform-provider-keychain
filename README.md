# Terraform Provider Keychain

This provider is used to manage the local macOS Keychain. Ideal for syncing a set of application or wifi passwords.

Note that this is **macOS only**!

## Example

See [test.tf](./test.tf) for more examples.

```hcl
data "keychain" "test_ssid_name" {
  service = "AirPort"
  account = "An ssid name"
}

resource "keychain" "test_ssid" {
  account = "A ssid name"
  data    = "My wifi password"
}
```

## Testing

Current test suite is minimal:

```
$ go test
```

Trying out the provider using `test.tf`:

```
$ go build -o terraform-provider-keychain
$ terraform init
$ terraform plan
$ terraform apply
```

