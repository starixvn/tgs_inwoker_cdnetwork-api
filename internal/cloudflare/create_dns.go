package cloudflare

import (
	"bytes"
	"cdnetwork/pkg/log"
	"cdnetwork/pkg/util"
	"encoding/json"
	"fmt"
	"net/http"
)

// DNSRecord represents the structure of a DNS record
type DNSRecord struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

func CreateDNS(domain string) {

	config := util.GetConfig()

	apiToken := config.CloudflareToken
	zoneName := extractDomain(domain)
	// Zone ID of your domain
	zoneID := GetZoneId(zoneName)

	// DNS record data
	record := DNSRecord{
		Type:    "CNAME",
		Name:    domain,            // Name of the record
		Content: config.DnsContent, // IP address or content of the record
		TTL:     1,
		Proxied: false, // Whether the record is proxied through Cloudflare
	}

	// Convert record data to JSON
	recordJSON, err := json.Marshal(record)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// API endpoint to create DNS record
	apiEndpoint := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneID)

	// Create HTTP request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(recordJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		log.LogError(fmt.Sprintf("Error sending request: %v", err))
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		log.LogError(fmt.Sprintf("Error: %s", resp.Status))
		return
	}

	fmt.Println("DNS record created successfully!")
}
