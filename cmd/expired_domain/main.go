package main

import (
	"cdnetwork/internal/google_sheet"
	"cdnetwork/internal/namecheap"
	"cdnetwork/internal/postgresql"
)

func main() {
	domains := namecheap.GetExpiredDomains()
	domainsForExcel := postgresql.GetAgents(domains)
	google_sheet.CreateExpiredDomainExcel(domainsForExcel)
}
