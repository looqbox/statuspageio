package statuspageio

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ListComponents executes a request that returns a list of created components
// at your Statuspage
func (request BaseRequest) ListComponents() (string, []ComponentsResponse) {
	finalRequest := RequestParam{
		Url:     request.Url + "/components",
		Method:  "GET",
		Headers: request.Headers,
		Body:    "",
		Client:  request.Client,
	}

	res := finalRequest.exec()

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse []ComponentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}
