package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
	"strings"
)

type ControlGroup {
	domain string,
}

func (c ControlGroup) CreateDomain() void {
	domainNames := strings.Split(c.domainName, ",")
	loginNames := []*client.LoginName{{LoginName: "raypro"}}
	controlGroupName := "Raypro"

	config := auth.BasicConfig{
		UserName: "raypro", 
		ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", 
		EndPoint: "api.cdnetworks.com", 
		Uri: "/user/control-groups/fdb99358-9e43-4197-95c7-7c2213a71a6f", 
		Method: "PUT",
	}

	CreateDomainList(domainNames)
}  

func CreateDomains(createDomainNames string)
	// 使用逗号分割多个域名
	domainNames := strings.Split(createDomainNames, ",")
	loginNames := []*client.LoginName{{LoginName: "raypro"}}
	controlGroupName := "Raypro"

	config := auth.BasicConfig{
		UserName: "raypro", 
		ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", 
		EndPoint: "api.cdnetworks.com", 
		Uri: "/user/control-groups/fdb99358-9e43-4197-95c7-7c2213a71a6f", 
		Method: "PUT",
	}

	CreateDomainList(domainNames)
}

func (c *CDNManager) createDomainList(domainNames []string) {
	for _, domainName := range domainNames {
		response := c.createDomain(domainName)
		fmt.Printf("Response body for domain %s and name %s is %#v\n", domainName, c.Name, response)
	}
}

func createDomain(domainName string, *config config) interface{}{
	request := client.EditControlGroupRequest{
        DomainList:      []*string{&domainName},
		AccountList:     c.LoginNames,
		IsAdd:           true,
		ControlGroupName: c.ControlGroupName,
    }

    return auth.Invoke(config, editControlGroupRequest.String())
}
