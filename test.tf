# resource "keychain" "pault" {
#   # SecClass       = "GenericPassword"
#   service = "MyService"
#   account = "gabriel"
#   label   = "A label"

#   # AccessGroup    = "A123456789.group.com.mycorp"
#   data = "toomanysecrets"

#   # Synchronizable = true
#   # Accessible     = true
# }

resource "keychain" "pault" {
  service = "AirPort"
  account = "ssid name"
  label   = "ssid name"

  data = "wifi password"

  description = "test description"

  # synchronizable = true
  # accessible     = true
}
