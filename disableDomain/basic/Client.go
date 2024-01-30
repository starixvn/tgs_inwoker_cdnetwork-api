package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	)
func main() {
	var DomainID string
	disableSingleDomainServiceRequest := client.DisableSingleDomainServiceRequest{}
	fmt.Println(DomainID)
	apiuri := fmt.Sprintf("/api/domain/%s/disable", DomainID)
	config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "PUT"}
	response := auth.Invoke(config, disableSingleDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
