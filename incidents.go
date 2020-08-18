package statuspageio

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func (request BaseRequest) mountRequest(URL string, method string, headers []Header, body interface{}) RequestFormat {
	finalRequest := RequestFormat{
		URL:     URL,
		Method:  method,
		Headers: headers,
		Body:    body,
		Client:  request.Client,
	}

	return finalRequest
}

// ListIncidents executes a request and return a list of reported incidents at Statuspage
func (request BaseRequest) ListIncidents(searchQuery string) (string, []IncidentsResponse) {
	body := ""
	if searchQuery != "" {
		body = "q=" + searchQuery
	}

	URL := request.URL + "/incidents"
	finalRequest := request.mountRequest(URL, "GET", request.Headers, body)

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

// ListIncidents executes a request and return a list of unresolved incidents at Statuspage
func (request BaseRequest) ListUnresolvedIncidents() (string, []IncidentsResponse) {
	body := ""

	URL := request.URL + "/incidents/unresolved"
	finalRequest := request.mountRequest(URL, "GET", request.Headers, body)

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
	finalRequest := request.mountRequest(URL, "GET", request.Headers, "")

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
	incident := incident{
		Incident: incidentBody,
	}
	finalRequest := request.mountRequest(URL, "PATCH", request.Headers, incident)

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
	incident := incident{
		Incident: incidentBody,
	}
	finalRequest := request.mountRequest(URL, "POST", request.Headers, incident)

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
	finalRequest := request.mountRequest(URL, "DELETE", request.Headers, "")

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
