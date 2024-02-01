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

	config := auth.BasicConfig{UserName: "{userName}", ApiKey: "{apiKey}", EndPoint: "api.cdnetworks.com", Uri: "/api/domain/11", Method: "DELETE"}
	response := auth.Invoke(config, deleteApiDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
