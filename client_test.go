package statuspageio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
	apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"

	expectedResult := BaseRequest{
		Url: "https://api.statuspage.io/v1/pages/" + pageId,
		Headers: []Header{
			Header{
				Name:  "Authorization",
				Value: "OAuth " + apiKey,
			},
		},
	}

	result := Connect(pageId, apiKey)

	assert.Equal(t, expectedResult, result)
}
