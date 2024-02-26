package cloudflare

import (
	"cdnetwork/pkg/log"
	"fmt"
	"net/http"
)

func sendRequest(apiToken, httpMethod, endpoint string) (*http.Response, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest(httpMethod, endpoint, nil)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Error create request: %s, method: %s, err: %v", endpoint, httpMethod, err))
		return nil, err
	}

	// Add the necessary headers for Cloudflare API
	req.Header.Add("Authorization", "Bearer "+apiToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Error in send request: %s, method: %s, err: %v", endpoint, httpMethod, err))
		return nil, err
	}

	return resp, nil
}
