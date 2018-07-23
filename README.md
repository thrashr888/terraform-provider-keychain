
## Example

```terraform

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
