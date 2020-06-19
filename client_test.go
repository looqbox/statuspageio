package statuspageio

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	pageID := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
	apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"

	expectedResult := BaseRequest{
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

	result := Connect(pageID, apiKey)

	assert.Equal(t, expectedResult, result)
}
