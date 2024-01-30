package main

import (
"fmt"
"openApi-authentication/common/auth"
"openApi-authentication/api/client"
)

func main() {

    addCdnDomainRequest := client.AddCdnDomainRequest{}
        addCdnDomainRequest.SetDomainName(" ")
        addCdnDomainRequest.SetItemId("3d("40016213")
        originConfig := client.AddCdnDomainRequestOriginConfig{}
        originConfig.SetOriginIps(source)
        addCdnDomainRequest.SetOriginConfig(originConfig)
        addCdnDomainRequest.SetVersion("1.0.0")
        
    config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: "/cdnw/api/domain", Method: "POST"}
    response := auth.Invoke(config, addCdnDomainRequest.String())0")
        addCdnDomainRequest.SetContractI
    fmt.Printf("response body is %#v\n", response