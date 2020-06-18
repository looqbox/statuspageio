package statuspageio_test

import (
	"fmt"

	"github.com/looqbox/statuspageio"
)

func ExampleConnect() {
	pageId := "tLwrSv54NJCJWHYUzLl3n0dMngEc9Cr8"
	apiKey := "5OCMVqruAKqXsZpT5bbrs3ISgsgMu01Z"

	request := statuspageio.Connect(pageId, apiKey)
	fmt.Println(request.Url, request.Headers)
	// Output:
	// https://api.statuspage.io/v1/pages/tLwrSv54NJCJWHYUzLl3n0dMngEc9Cr8 [{Authorization OAuth 5OCMVqruAKqXsZpT5bbrs3ISgsgMu01Z}]
}
