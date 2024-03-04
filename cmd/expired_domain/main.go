package main

import (
	"cdnetwork/pkg/google_sheet"
	"cdnetwork/pkg/namecheap"
	"cdnetwork/pkg/postgresql"
)

func main() {
	domains := namecheap.GetExpiredDomains()
	domainsForExcel := postgresql.GetAgents(domains)
	google_sheet.CreateExpiredDomainExcel(domainsForExcel)
}
