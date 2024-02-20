package main

import (
	"cdnetwork/controlGroup"
)

func main() {

	domainName := "example.com,example.net" // Example domain names
	controlGroup.CreateDomains(domainName)
}
