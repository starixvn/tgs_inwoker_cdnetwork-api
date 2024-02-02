package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"strings"
)

type SslCertificate struct {
	CertificateID string `xml:"certificate-id"`
	Name          string `xml:"name"`
	// 省略其他字段...
}

type SslCertificates struct {
	XMLName        xml.Name         `xml:"ssl-certificates"`
	SslCertificate []SslCertificate `xml:"ssl-certificate"` // 注意这里的字段名改为与XML标签相匹配
}

// 定义JSON结构体
type CertificateJSON struct {
	CertificateID string `json:"certificate_id"`
	DomainName    string `json:"domain_name"`
}

func main() {

	queryCertificateListRequest := client.QueryCertificateListRequest{}

	config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: "/api/ssl/certificate", Method: "GET"}
	response := auth.Invoke(config, queryCertificateListRequest.String())
	// fmt.Printf("response body is %#v\n", response)

	xmlData := response
	var sslCerts SslCertificates
	err := xml.Unmarshal([]byte(xmlData), &sslCerts)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	// 处理并转换数据
	var certsJSON []CertificateJSON
	for _, cert := range sslCerts.SslCertificate { // 注意这里的字段名与结构体中的字段名相匹配
		splitName := strings.Split(cert.Name, "_")
		domainName := splitName[len(splitName)-1] // 取"_"后的部分

		certsJSON = append(certsJSON, CertificateJSON{
			CertificateID: cert.CertificateID,
			DomainName:    domainName,
		})
	}

	// 将处理后的数据转换为JSON
	jsonData, err := json.MarshalIndent(certsJSON, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return
	}
	// 输出JSON
	fmt.Println(string(jsonData))
	specifiedDomainName := "globalhttptest.rpgp.cc"
	for _, cert := range certsJSON {
		if cert.DomainName == specifiedDomainName {
			fmt.Printf("Certificate ID for '%s' is: %s\n", specifiedDomainName, cert.CertificateID)
			break // 找到匹配项后退出循环
		}
	}
}
