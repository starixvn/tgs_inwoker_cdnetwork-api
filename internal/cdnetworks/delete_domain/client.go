package delete_domain

import (
	"cdnetwork/internal/cdnetworks/delete_domain/models"
	"cdnetwork/pkg/auth"
	"fmt"
)

func DeleteDomain(domainName string, deletedid string) {
	request := &models.DeleteApiDomainServiceRequest{}
	request.SetDomainName(domainName)
	request.SetDomainId(deletedid)

	apiuri := fmt.Sprintf("/api/domain/%s", deletedid)
	config := auth.BasicConfig{
		Uri:    apiuri,
		Method: "DELETE"}
	response := auth.Invoke(config, request.String())
	fmt.Printf("response body is %#v\n", response)
}
