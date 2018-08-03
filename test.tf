// keychain_item data and two tests


resource "keychain_item" "test_key_1" {
  service = "TEST SERVICE"
  label   = "Test Label 1"
  account = "Test Account 1"
  data    = "Test Password 1"
}

output "item-1" {
  value = "${keychain_item.test_key_1.data}"
}

resource "keychain_item" "test_key_2" {
  service = "TEST SERVICE"
  label   = "Test Label 2"
  account = "Test Account 2"
  data    = "Test Password 2"
}

output "item-2" {
  value = "${keychain_item.test_key_2.data}"
}

data "keychain_item" "test_key_0" {
  service = "TEST SERVICE"
  account = "Test Account 1"
}

output "item-data" {
  value = "password: ${data.keychain_item.test_key_0.data}"
}

// keychain_wifi data and two tests

# resource "keychain_wifi" "test_wifi_1" {
#   account = "SSID Name 1"
#   data    = "WiFi Password 1"
# }
# output "wifi-1" {
#   value = "${keychain_wifi.test_wifi_1.data}"
# }

# resource "keychain_wifi" "test_wifi_2" {
#   account = "SSID Name 2"
#   data    = "WiFi Password 2"
# }
# output "wifi-2" {
#   value = "${keychain_wifi.test_wifi_2.data}"
# }

data "keychain_wifi" "test_wifi_0" {
  account = "Burnside Brewing Co."
}
output "wifi-data" {
  value = "${data.keychain_wifi.test_wifi_0.data}"
}