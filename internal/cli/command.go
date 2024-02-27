package cli

import (
	"cdnetwork/internal/cdnetworks/control_group"
	"cdnetwork/internal/cdnetworks/create_certificate"
	"cdnetwork/internal/cdnetworks/create_domain"
	"cdnetwork/internal/cdnetworks/delete_domain"
	"cdnetwork/internal/cdnetworks/disable_domain"
	"cdnetwork/internal/cdnetworks/enable_domain"
	"cdnetwork/internal/cdnetworks/get_certificate"
	"cdnetwork/internal/cdnetworks/get_domain_id"
	"cdnetwork/internal/cloudflare"
	"time"
)

type CdnCommand struct {
	DomainName string
	DomainId   string
	OriginSet  string
}

// Command interface
type CdnCommander interface {
	UpdateControlGroup()
	CreateDomain()
	DeleteDomain()
	DisableDomain()
	EnableDomain()
	GetCertificate()
	GetDomainId()
	CreateCertificate()
}

func (c *CdnCommand) UpdateControlGroup() {
	control_group.UpdateControlGroup(c.DomainName)
}

func (c *CdnCommand) CreateDomain() {
	create_domain.CreateDomains(c.DomainName, c.OriginSet)

	time.Sleep(1 * time.Second)
	control_group.UpdateControlGroup(c.DomainName)
	time.Sleep(1 * time.Second)
	create_certificate.CreateCertificateByDomain(c.DomainName)

	cloudflare.CreateDNS(c.DomainName)

}

func (c *CdnCommand) DeleteDomain() {
	delete_domain.DeleteDomain(c.DomainName)
}

func (c *CdnCommand) DisableDomain() {
	disable_domain.DisableDomain(c.DomainName)
}

func (c *CdnCommand) EnableDomain() {
	enable_domain.EnableDomain(c.DomainId)
}

func (c *CdnCommand) GetCertificate() string {
	res, err := get_certificate.GetCertificate(c.DomainName)
	if err != nil {
		return err.Error()
	}
	return res
}

func (c *CdnCommand) GetDomainId() int {
	return get_domain_id.GetDomainId(c.DomainName)
}

func (c *CdnCommand) CreateCertificate() {
	create_certificate.CreateCertificateByDomain(c.DomainName)
}
