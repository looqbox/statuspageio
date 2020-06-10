package statuspageio

import (
        "testing"

        "github.com/stretchr/testify/assert"

        "github.com/looqbox/statuspage-go/internal/root"
)

func TestConnect(t *testing.T) {
        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"

        expectedResult := root.BaseRequest{
                Url: "https://api.statuspage.io/v1/pages/" + pageId,
                Headers: []root.Header{
                        root.Header{
                                Name:  "Authorization",
                                Value: "OAuth " + apiKey,
                        },
                },
        }

        result := Connect(pageId, apiKey)

        assert.Equal(t, expectedResult, result)
}
