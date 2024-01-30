package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
	"io/ioutil"
)
func main() {
	content, err := ioutil.ReadFile("readdomainid.txt")
    if err != nil {
        fmt.Println("Error reading readdomain.txt:", err)
        return
    }
	DomainID := string(content)
	disableSingleDomainServiceRequest := client.DisableSingleDomainServiceRequest{}
	apiuri := fmt.Sprintf("/api/domain/%s/disable", .DomainID)
	config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "PUT"}
	response := auth.Invoke(config, disableSingleDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
