data "keychain" "test_ssid_0" {
  service = "AirPort"
  account = "SSID Name 1"
}

output "password-data" {
  value = "${data.keychain.test_ssid_0.data}"
}

resource "keychain" "test_ssid_1" {
  label   = "SSID Name 1"
  account = "SSID Name 1"
  data    = "WiFi Password 1"
}

output "password-1" {
  value = "${keychain.test_ssid_1.data}"
}

resource "keychain" "test_ssid_2" {
  label   = "SSID Name 2"
  account = "SSID Name 2"
  data    = "WiFi Password 2"
}

output "password-2" {
  value = "${keychain.test_ssid_2.data}"
}
