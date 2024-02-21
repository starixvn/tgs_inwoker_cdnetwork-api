package enableDomain

import (
	"cdnetwork/internal/cdnetworks/enableDomain/models"
	"cdnetwork/pkg/auth"
	"fmt"
)

func EnableDomain(domainId string) {
	request := models.EnableSingleDomainServiceRequest{}
	fmt.Println("我是ID", domainId)
	apiuri := fmt.Sprintf("/api/domain/%s/enable", domainId)
	fmt.Println(apiuri)
	config := auth.BasicConfig{
		Uri:    apiuri,
		Method: "PUT",
	}
	fmt.Printf(config.Uri)
	response := auth.Invoke(config, request.String())
	fmt.Printf("response body is %#v\n", response)
}
