package main

import (
"fmt"
"openApi-authentication/common/auth"
"openApi-authentication/api/client"
)

func main() {

    enableSingleDomainServiceRequest := client.EnableSingleDomainServiceRequest{}
        
    config := auth.BasicConfig{UserName: "{userName}", ApiKey: "{apiKey}", EndPoint: "api.cdnetworks.com", Uri: "/api/domain/*/enable", Method: "PUT"}
    response := auth.Invoke(config, enableSingleDomainServiceRequest.String())
    fmt.Printf("response body is %#v\n", response)
}

