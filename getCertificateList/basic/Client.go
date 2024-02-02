package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"strings"
	"os"
)

type SslCertificate struct {
	CertificateID string `xml:"certificate-id"`
	Name          string `xml:"name"`
	// 省略其他字段...
}

type SslCertificates struct {
	XMLName        xml.Name         `xml:"ssl-certificates"`
	SslCertificate []SslCertificate `xml:"ssl-certificate"` // 這裡的字段名改為與XML標籤相匹配
}

// 定義JSON結構體
type CertificateJSON struct {
	CertificateID string `json:"certificate_id"`
	DomainName    string `json:"domain_name"`
}

func main() {
	//call api
	queryCertificateListRequest := client.QueryCertificateListRequest{}

	config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: "/api/ssl/certificate", Method: "GET"}
	response := auth.Invoke(config, queryCertificateListRequest.String())
	//使用response結果解析xml數據
	xmlData := response
	var sslCerts SslCertificates
	err := xml.Unmarshal([]byte(xmlData), &sslCerts)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	// 處理並轉換數據
	var certsJSON []CertificateJSON
	for _, cert := range sslCerts.SslCertificate { // 注意這裡的字段名與結構體中的字段名相匹配
		splitName := strings.Split(cert.Name, "_")
		domainName := splitName[len(splitName)-1] // 取"_"後的部分

		certsJSON = append(certsJSON, CertificateJSON{
			CertificateID: cert.CertificateID,
			DomainName:    domainName,
		})
	}

	// 將處理後的數據轉換为JSON
	jsonData, err := json.MarshalIndent(certsJSON, "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
		return
	}
	// 輸出JSON
	fmt.Println(string(jsonData))
	targetDomainName := os.Getenv("certname")
	for _, cert := range certsJSON {
		if cert.DomainName == targetDomainName  {
			fmt.Println(cert.CertificateID)
			break
		}
	}
}
