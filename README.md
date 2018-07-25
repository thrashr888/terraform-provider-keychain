# Terraform Provider Keychain

This provider is used to manage the local macOS Keychain. Ideal for syncing a set of application or wifi passwords.

Note that this is **macOS only**!


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)


## Building The Provider

Clone repository to: `$GOPATH/src/github.com/thrashr888/terraform-provider-keychain`

```sh
$ mkdir -p $GOPATH/src/github.com/thrashr888; cd $GOPATH/src/github.com/thrashr888
$ git clone git@github.com:thrashr888/terraform-provider-keychain
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/thrashr888/terraform-provider-keychain
$ go install
```


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


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-keychain
...
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


## TODO

- [ ] Make the data source usable
- [ ] Allow the data source to return multiple items (??? might not want to)
- [ ] Add a separate provider for specifically wifi passwords
- [X] Add API docs
- [X] Add instructions on using with Terraform
- [X] Add travis config
