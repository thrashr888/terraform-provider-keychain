resource "keychain" "test-ssid" {
  service     = "AirPort"
  account     = "ssid name"
  label       = "ssid name"
  data        = "wifi password"
  description = "test description"
}
