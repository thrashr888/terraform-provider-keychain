resource "keychain" "test-ssid" {
  account = "ssid name"
  data    = "wifi password 2"
}

resource "keychain" "test-ssid-2" {
  account = "ssid name 2"
  data    = "wifi password 3"
}

output "password" {
  value = "${keychain.test-ssid.data}"
}
