package statuspageio

import (
	"bytes"
	"log"
	"net/http"
	"strings"
)

func addHeaders(headers []Header, req *http.Request) {
	for i := 0; i < len(headers); i++ {
		req.Header.Add(headers[i].Name, headers[i].Value)
	}
}

func (request RequestParam) exec() *http.Response {
	body := strings.NewReader(request.Body)

	req, err := http.NewRequest(request.Method, request.Url, body)
	if err != nil {
		log.Fatal(err)
	}

	addHeaders(request.Headers, req)

	// Execute http request
	res, err := request.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (request RequestJson) exec() *http.Response {
	req, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(request.Body))
	if err != nil {
		log.Fatal(err)
	}

	addHeaders(request.Headers, req)

	// Execute http request
	res, err := request.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
