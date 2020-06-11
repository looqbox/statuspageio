// Package create to interact with statuspage.io API, created mainly to
// generate requests to manage incidents creation/update.
package statuspageio

import (
)

// Basic struct to store values used for every request at Statuspage.io
// API, like the base url and the authorization header
type BaseRequest struct {
        Url     string
        Headers []Header
}

// Struct to save Headers
type Header struct {
        Name  string
        Value string
}

// Connect mounts a base request to connect with the Statuspage.io API.
// It needs a pageId and a apiKey that you can retrieve at the statuspage
// manage page.
func Connect(pageId string, apiKey string) BaseRequest {
        request := root.BaseRequest{
                Url: "https://api.statuspage.io/v1/pages/" + pageId,
                Headers: []root.Header{
                        root.Header{
                                Name:  "Authorization",
                                Value: "OAuth " + apiKey,
                        },
                },
        }

        return request
}
