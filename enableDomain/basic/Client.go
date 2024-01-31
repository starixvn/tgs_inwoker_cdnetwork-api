package main

import (
"fmt"
"openApi-authentication/common/auth"
"openApi-authentication/api/client"
)

func main() {
    DomainID := os.Args[1]
    enableSingleDomainServiceRequest := client.EnableSingleDomainServiceRequest{}
	fmt.Println("我是ID",DomainID)
	apiuri := fmt.Sprintf("/api/domain/%s/enable", DomainID)
	fmt.Println(apiuri)
	config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "PUT"}
	fmt.Printf(config.Uri)
	response := auth.Invoke(config, disableSingleDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
