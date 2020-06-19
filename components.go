package statuspageio

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ListComponents executes a request that returns a list of created components
// at your Statuspage
func (request BaseRequest) ListComponents() (string, []ComponentsResponse) {
	finalRequest := RequestFormat{
		URL:     request.URL + "/components",
		Method:  "GET",
		Headers: request.Headers,
		Body:    "",
		Client:  request.Client,
	}

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse []ComponentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}
