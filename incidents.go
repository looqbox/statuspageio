package statuspageio

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func mountRequest(URL string, method string, headers []Header, body interface{}, client *http.Client) RequestFormat {
	request := RequestFormat{
		URL:     URL,
		Method:  method,
		Headers: headers,
		Body:    body,
		Client:  client,
	}

	return request
}

// ListIncidents executes a request and return a list of reported incidents at Statuspage
func (request BaseRequest) ListIncidents(searchQuery string) (string, []IncidentsResponse) {
	body := ""
	if searchQuery != "" {
		body = "q=" + searchQuery
	}

	URL := request.URL + "/incidents"
	finalRequest := mountRequest(URL, "GET", request.Headers, body, request.Client)

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse []IncidentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}

// GetIncident retrieve information about a specific incident
func (request BaseRequest) GetIncident(incidentID string) (string, IncidentsResponse) {
	URL := request.URL + "/incidents/" + incidentID
	finalRequest := mountRequest(URL, "GET", request.Headers, "", request.Client)

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse IncidentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}

// UpdateIncident Updates one incident information
func (request BaseRequest) UpdateIncident(incidentID string, incidentBody IncidentBody) (string, IncidentsResponse) {
	URL := request.URL + "/incidents/" + incidentID
	finalRequest := mountRequest(URL, "PATCH", request.Headers, incidentBody, request.Client)

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse IncidentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}

// CreateIncident creates a new incident at the page
func (request BaseRequest) CreateIncident(incidentBody IncidentBody) (string, IncidentsResponse) {
	URL := request.URL + "/incidents"
	finalRequest := mountRequest(URL, "POST", request.Headers, incidentBody, request.Client)

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse IncidentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}

// DeleteIncident deletes an incident from the page
func (request BaseRequest) DeleteIncident(incidentID string) (string, IncidentsResponse) {
	URL := request.URL + "/incidents/" + incidentID
	finalRequest := mountRequest(URL, "DELETE", request.Headers, "", request.Client)

	res, err := finalRequest.exec()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var structuredResponse IncidentsResponse
	json.Unmarshal([]byte(jsonData), &structuredResponse)

	return res.Status, structuredResponse
}
