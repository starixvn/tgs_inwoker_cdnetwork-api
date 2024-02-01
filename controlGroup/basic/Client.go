package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
	"strings"
)

func main() {
    createDomainNames := os.Getenv("createdomainname")

    // 使用逗号分割多个域名
    domainNames := strings.Split(createDomainNames, ",")

    loginNameList := []*client.LoginName{}
    loginName1 := client.LoginName{}
    loginName1.SetLoginName("raypro")
    loginNameList = append(loginNameList, &loginName1)

    controlGroupName := "Raypro"

    config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: "/user/control-groups/fdb99358-9e43-4197-95c7-7c2213a71a6f", Method: "PUT"}

    for _, createDomainName := range domainNames {
        editControlGroupRequest := client.EditControlGroupRequest{}
        editControlGroupRequest.SetDomainList([]*string{&createDomainName})
        editControlGroupRequest.SetAccountList(loginNameList)
        editControlGroupRequest.SetIsAdd(true)
        editControlGroupRequest.SetControlGroupName(controlGroupName)

        response := auth.Invoke(config, editControlGroupRequest.String())
        fmt.Printf("Response body for domain %s is %#v\n", createDomainName, response)
    }
}
