// Package statuspageio generate requests used to interact with statuspage.io API
// , created mainly to generate requests to manage incidents creation/update.
package statuspageio

import (
	"net/http"
	"time"
)

// BaseRequest struct store values used for every request at Statuspage.io API,
// like the base url and the authorization header, also includes a http.Client
// to execute requests
type BaseRequest struct {
	URL     string
	Headers []Header
	Client  *http.Client
}

// Header struct save information in Header format
type Header struct {
	Name  string
	Value string
}

// Connect mounts a base request to connect with the Statuspage.io API.
// It needs a pageId and a apiKey that you can retrieve at the statuspage
// manage page.
func Connect(pageID string, apiKey string) BaseRequest {
	request := BaseRequest{
		URL: "https://api.statuspage.io/v1/pages/" + pageID,
		Headers: []Header{
			Header{
				Name:  "Authorization",
				Value: "OAuth " + apiKey,
			},
		},
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	return request
}
