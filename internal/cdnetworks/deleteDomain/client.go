package deleteDomain

import (
	"cdnetwork/internal/cdnetworks/deleteDomain/models"
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
