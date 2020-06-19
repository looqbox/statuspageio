package statuspageio

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func addHeaders(headers []Header, req *http.Request) {
	for i := 0; i < len(headers); i++ {
		req.Header.Add(headers[i].Name, headers[i].Value)
	}
}

func (request RequestFormat) exec() (*http.Response, error) {
	var body io.Reader
	switch request.Body.(type) {
	default:
		return nil, errors.New("Body type not recognized")
	case string:
		body = strings.NewReader(request.Body.(string))
	case IncidentBody:
		bodyJSON, err := json.Marshal(request.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = bytes.NewBuffer(bodyJSON)
	}

	req, err := http.NewRequest(request.Method, request.URL, body)
	if err != nil {
		log.Fatal(err)
	}

	addHeaders(request.Headers, req)

	// Execute http request
	res, err := request.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res, nil
}
