package getDomainId

import (
	"cdnetwork/internal/cdnetworks/getDomainId/models"
	"cdnetwork/pkg/auth"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type DomainSummary struct {
	DomainID         int    `xml:"domain-id"`
	DomainName       string `xml:"domain-name"`
	ServiceType      int    `xml:"service-type"`
	CName            string `xml:"cname"`
	Status           string `xml:"status"`
	CdnServiceStatus bool   `xml:"cdn-service-status"`
	Enabled          bool   `xml:"enabled"`
}

type DomainList struct {
	XMLName         xml.Name        `xml:"domain-list"`
	DomainSummaries []DomainSummary `xml:"domain-summary"`
}
type Domainkv struct {
	DomainSummaries []DomainSummary `json:"DomainSummaries"`
}

func GetDomainId(domainName string) int {

	request := models.QueryApiDomainListServiceRequest{}

	config := auth.BasicConfig{
		Uri:    "/api/domain",
		Method: "GET",
	}

	response := auth.Invoke(config, request.String())

	xmlData := response
	var domainList DomainList
	err := xml.Unmarshal([]byte(xmlData), &domainList)
	if err != nil {
		fmt.Println("解析 XML 失败:", err)
		return -1
	}

	// 将结构体转换为 JSON
	jsonData, err := json.Marshal(domainList)
	if err != nil {
		fmt.Println("转换为 JSON 失败:", err)
		return -1
	}

	// 打印 JSON 数据
	fmt.Println(string(jsonData))
	var domainkv Domainkv

	err = json.Unmarshal([]byte(jsonData), &domainkv)
	if err != nil {
		fmt.Println("解码JSON失败:", err)
		return -1
	}
	fmt.Println("Target Domain Name:", domainName)
	var DomainID int
	// 遍历DomainSummaries并输出所需的键值对
	for _, summary := range domainkv.DomainSummaries {
		if summary.DomainName == domainName {
			DomainID = summary.DomainID
			fmt.Println("domainidr:", DomainID)
			return DomainID
			break
		}
	}
	return -1
}
