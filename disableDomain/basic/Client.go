package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
	)
func main() {
	DomainID := os.Args[1]
	disableSingleDomainServiceRequest := client.DisableSingleDomainServiceRequest{}
	fmt.Println("我是ID",DomainID)
	apiuri := fmt.Sprintf("/api/domain/%s/disable", DomainID)
	fmt.Printf()
	fmt.Println(apiuri)
	config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "PUT"}
	fmt.Printf(config.Uri)
	response := auth.Invoke(config, disableSingleDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
