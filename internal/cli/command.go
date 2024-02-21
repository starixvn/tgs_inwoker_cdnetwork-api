package cli

import (
	"cdnetwork/internal/cdnetworks/controlGroup"
	"cdnetwork/internal/cdnetworks/createDomain"
	"cdnetwork/internal/cdnetworks/deleteDomain"
	"cdnetwork/internal/cdnetworks/disableDomain"
	"cdnetwork/internal/cdnetworks/enableDomain"
	"cdnetwork/internal/cdnetworks/getCertificate"
	"cdnetwork/internal/cdnetworks/getDomainId"
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
}

func (c *CdnCommand) UpdateControlGroup() {
	controlGroup.UpdateControlGroup(c.DomainName)
}

func (c *CdnCommand) CreateDomain() {
	createDomain.CreateDomains(c.DomainName, c.OriginSet)
}

func (c *CdnCommand) DeleteDomain() {
	deleteDomain.DeleteDomain(c.DomainName, c.DomainId)
}

func (c *CdnCommand) DisableDomain() {
	disableDomain.DisableDomain(c.DomainName)
}

func (c *CdnCommand) EnableDomain() {
	enableDomain.EnableDomain(c.DomainId)
}

func (c *CdnCommand) GetCertificate() {
	getCertificate.GetCertificate(c.DomainName)
}

func (c *CdnCommand) GetDomainId() {
	getDomainId.GetDomainId(c.DomainName)
}
