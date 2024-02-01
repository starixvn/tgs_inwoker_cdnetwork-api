package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
)

func main() {
	deletedname := os.Getenv("setdomaimname")
	deletedid := os.Args[1]
	deleteApiDomainServiceRequest := client.DeleteApiDomainServiceRequest{}
	deleteApiDomainServiceRequest.SetDomainName(deletedname)
	deleteApiDomainServiceRequest.SetDomainId(deletedid)
	apiuri := fmt.Sprintf("/api/domain/%s", DomainID)
	config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "DELETE"}
	response := auth.Invoke(config, deleteApiDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
