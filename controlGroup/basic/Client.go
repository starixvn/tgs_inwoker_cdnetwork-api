package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
)

func main() {

	editControlGroupRequest := client.EditControlGroupRequest{}
	var subeditControlGroupRequest0 = " "
	editControlGroupRequest.SetDomainList([]*string{&subeditControlGroupRequest0})
	loginNameList := []*LoginName{}
	var i = 0
	loginName1 := client.LoginName{}
	loginName1.SetLoginName("raypro")
	loginNameList[i] = &loginName1
	i++
	editControlGroupRequest.SetAccountList(loginNameList)
	editControlGroupRequest.SetIsAdd(true)
	editControlGroupRequest.SetControlGroupName("Raypro")

	config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: "/user/control-groups/fdb99358-9e43-4197-95c7-7c2213a71a6f", Method: "PUT"}
	response := auth.Invoke(config, editControlGroupRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
