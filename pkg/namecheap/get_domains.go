package namecheap

import (
	"cdnetwork/internal/log"
	"cdnetwork/internal/util"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	apiUrl  = "https://api.namecheap.com/xml.response"
	command = "namecheap.domains.getList"
)

type DomainGetListResult struct {
	Domains []Domain `xml:"Domain"`
}

type Paging struct {
	TotalItems  int `xml:"TotalItems"`
	CurrentPage int `xml:"CurrentPage"`
	PageSize    int `xml:"PageSize"`
}

type CommandResponse struct {
	DomainGetListResult DomainGetListResult `xml:"DomainGetListResult"`
	Paging              Paging              `xml:"Paging"`
}

type ApiResponse struct {
	XMLName          xml.Name        `xml:"ApiResponse"`
	Status           string          `xml:"Status,attr"`
	RequestedCommand string          `xml:"RequestedCommand"`
	CommandResponse  CommandResponse `xml:"CommandResponse"`
}

type Domain struct {
	ID         string `xml:"ID,attr"`
	Name       string `xml:"Name,attr"`
	User       string `xml:"User,attr"`
	Created    string `xml:"Created,attr"`
	Expires    string `xml:"Expires,attr"`
	IsExpired  string `xml:"IsExpired,attr"`
	IsLocked   string `xml:"IsLocked,attr"`
	AutoRenew  string `xml:"AutoRenew,attr"`
	WhoisGuard string `xml:"WhoisGuard,attr"`
	IsPremium  string `xml:"IsPremium,attr"`
	IsOurDNS   string `xml:"IsOurDNS,attr"`
}

type FilteredDomain struct {
	Name    string `xml:"Name,attr"`
	Created string `xml:"Created,attr"`
	Expires string `xml:"Expires,attr"`
}

func GetExpiredDomains() []FilteredDomain {
	config := util.GetConfig()
	userName := config.Namecheap.NamecheapUsername
	apiKey := config.Namecheap.NamecheapApiKey
	clinetIp := config.Namecheap.NamecheapClientIp
	pageSize := 100

	url := fmt.Sprintf("%s?ApiUser=%s&ApiKey=%s&UserName=%s&Command=%s&ClientIp=%s&PageSize=%d&Page=", apiUrl, userName, apiKey, userName, command, clinetIp, pageSize)

	var domains []Domain
	index := 1
	for index < 5 {
		fmt.Println(index)
		_domains, totolCount, shouldReturn := getDomains(url, index)
		if !shouldReturn {
			domains = append(domains, _domains...)
		}
		index++
		fmt.Println("totolCount:", totolCount)
		fmt.Println("index_pageSize:", index*pageSize)
		if totolCount > index*pageSize {
			break
		}
	}

	return filterDomainsWithExpired(domains)
}

func filterDomainsWithExpired(domains []Domain) []FilteredDomain {
	currentDate := time.Now()
	nextMonth := currentDate.AddDate(0, 1, 0)
	firstDayOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())
	lastDayOfNextMonth := firstDayOfNextMonth.AddDate(0, 1, -1)
	var result []FilteredDomain
	for _, domain := range domains {
		expireDate, err := parseDate(domain.Expires)
		if err != nil {
			fmt.Println("Error parsing expiration date:", err)
			continue
		}
		if isBetweenDates(expireDate, firstDayOfNextMonth, lastDayOfNextMonth) {
			result = append(result, FilteredDomain{
				Name:    domain.Name,
				Created: domain.Created,
				Expires: domain.Expires,
			})
		}
	}

	for i, domain := range result {
		fmt.Printf("%d, Name: %s, Created: %s, Expires: %s\n", i+1, domain.Name, domain.Created, domain.Expires)
	}
	return result
}

func parseDate(dateStr string) (time.Time, error) {
	// Assuming dateStr is in the format MM/DD/YYYY
	return time.Parse("01/02/2006", dateStr)
}

func isBetweenDates(date time.Time, start time.Time, end time.Time) bool {
	return date.After(start) && date.Before(end)
}
func getDomains(url string, index int) ([]Domain, int, bool) {

	url = url + strconv.Itoa(index)
	fmt.Println("url:", url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching domain list:", err)
		return []Domain{}, -1, true
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []Domain{}, -1, true
	}

	log.LogInfo(string(body))
	var apiResponse ApiResponse
	if err := xml.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return []Domain{}, -1, true
	}
	return apiResponse.CommandResponse.DomainGetListResult.Domains, apiResponse.CommandResponse.Paging.TotalItems, false
}
