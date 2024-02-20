package controlGroup

import (
	"cdnetwork/common/auth" // Import your auth package
	"cdnetwork/controlGroup/models"
	"fmt"
	"strings"
)

// ControlGroupManager represents a manager for control groups
type ControlGroupManager struct {
	ControlGroupName string
	LoginNames       []*models.LoginName
	AuthConfig       auth.BasicConfig
}

// NewControlGroupManager creates a new instance of ControlGroupManager
func NewControlGroupManager(controlGroupName, userName, apiKey, endPoint, uri, method string) *ControlGroupManager {
	config := auth.BasicConfig{
		UserName: userName,
		ApiKey:   apiKey,
		EndPoint: endPoint,
		Uri:      uri,
		Method:   method,
	}

	return &ControlGroupManager{
		ControlGroupName: controlGroupName,
		AuthConfig:       config,
	}
}

// SetLoginNameList sets the login name list
func (c *ControlGroupManager) SetLoginNameList(loginNames []*models.LoginName) {
	c.LoginNames = loginNames
}

// UpdateControlGroup updates the control group for a given domain
func (c *ControlGroupManager) UpdateControlGroup(domain string) {
	request := models.EditControlGroupRequest{}
	request.SetDomainList([]*string{&domain})
	request.SetAccountList(c.LoginNames)
	request.SetIsAdd(true)
	request.SetControlGroupName(c.ControlGroupName)

	response := auth.Invoke(c.AuthConfig, request.String())
	fmt.Printf("Response body for domain %s is %#v\n", domain, response)
}

// CreateDomains is responsible for creating domains
func CreateDomains(domainName string) {
	// 使用逗号分割多个域名
	domainNames := strings.Split(domainName, ",")

	loginNameList := []*models.LoginName{}
	loginName1 := models.LoginName{}
	loginName1.SetLoginName("raypro")
	loginNameList = append(loginNameList, &loginName1)

	controlGroupName := "Raypro"

	manager := NewControlGroupManager(
		controlGroupName,
		"raypro",
		"98514TVVFKeNJ6dfQ8Ho6xBn5OTvph",
		"api.cdnetworks.com",
		"/user/control-groups/fdb99358-9e43-4197-95c7-7c2213a71a6f",
		"PUT",
	)
	manager.SetLoginNameList(loginNameList)

	for _, createDomainName := range domainNames {
		manager.UpdateControlGroup(createDomainName)
	}
}
