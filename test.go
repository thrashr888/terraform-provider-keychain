package main

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

func main() {
	// test()

	testtwo()
}

func test() {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService("MyService")
	item.SetAccount("gabriel")
	item.SetLabel("A label")
	item.SetAccessGroup("A123456789.group.com.mycorp")
	item.SetData([]byte("toomanysecrets"))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	trustedApplications := []string{"/Applications/Mail.app"}
	item.SetAccess(&keychain.Access{Label: "AirPort", TrustedApplications: trustedApplications})

	err := keychain.AddItem(item)

	if err == keychain.ErrorDuplicateItem {
		// Duplicate
	}
}

func testtwo() {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService("AirPort")
	item.SetAccount("SSID NAME")
	item.SetLabel("SSID NAME")
	item.SetAccessGroup("AirPort")
	item.SetDescription("test description")
	item.SetData([]byte("wifi password"))
	item.SetSynchronizable(keychain.SynchronizableAny)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	// trustedApplications := []string{"/usr/libexec/airportd", "group://AirPort"}
	// item.SetAccess(&keychain.Access{Label: "AirPort", TrustedApplications: trustedApplications})

	trustedApplications := []string{"/usr/libexec/airportd"}
	item.SetAccess(&keychain.Access{Label: "AirPort", TrustedApplications: trustedApplications})

	err := keychain.AddItem(item)

	if err == keychain.ErrorDuplicateItem {
		// Duplicate
		fmt.Printf("DUPLICATE!", item)
	}

	fmt.Printf("Added %s\n", item)
}
