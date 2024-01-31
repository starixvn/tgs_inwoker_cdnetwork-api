package main

import (
"fmt"
"openApi-authentication/common/auth"
"openApi-authentication/api/client"
"os"
)

func main() {
    createDomainName1 := os.Getenv("createdomainname")
    OriginSet1 := os.Getenv("originset")
    addCdnDomainRequest := &client.AddCdnDomainRequest{}
        addCdnDomainRequest.SetDomainName(createDomainName1)
        fmt.Println(createDomainName1)
        addCdnDomainRequest.SetContractId("40016213")
        addCdnDomainRequest.SetItemId("30")
        originConfig := &client.AddCdnDomainRequestOriginConfig{}
        originConfig.SetOriginIps(OriginSet1)
        addCdnDomainRequest.SetOriginConfig(originConfig)
        addCdnDomainRequest.SetVersion("1.0.0")
        
    config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: "/cdnw/api/domain", Method: "POST"}
    response := auth.Invoke(config, addCdnDomainRequest.String())
    fmt.Printf("response body is %#v\n", response)
    
    
}