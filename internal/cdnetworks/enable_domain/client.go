package enable_domain

import (
	"cdnetwork/internal/cdnetworks/enable_domain/models"
	"cdnetwork/pkg/auth"
	"cdnetwork/pkg/log"
	"fmt"
)

func EnableDomain(domainId string) {
	request := models.EnableSingleDomainServiceRequest{}
	log.LogInfo(fmt.Sprintf("domain id: %s", domainId))

	apiuri := fmt.Sprintf("/api/domain/%s/enable", domainId)
	fmt.Println(apiuri)
	config := auth.BasicConfig{
		Uri:    apiuri,
		Method: "PUT",
	}
	log.LogInfo(fmt.Sprintf(config.Uri))
	response := auth.Invoke(config, request.String())
	log.LogInfo(fmt.Sprintf("response body is %#v\n", response))
}
