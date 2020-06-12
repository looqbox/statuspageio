package statuspageio

import (
	"encoding/json"
	"log"
	"net/http"
)

// IncidentBody saves information in a format required by statuspage.io API to
// interact with incidents at the site
type IncidentBody struct {
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	Impact        string    `json:"impact_override"`
	Notification  bool      `json:"deliver_notifications"`
	Body          string    `json:"body"`
	Components    Component `json:"components"`
	ComponentsIds []string  `json:"component_ids"`
}

// Component mounts a nested json to be used at IncidentBody
type Component struct {
	ComponentId string `json:"component_id"`
}

// Request is a general interface used for requests
type Request interface {
	exec() *http.Response
}

// RequestJson saves information in a format ready to do a http request with
// `Content-Type: application/json` Header
type RequestJson struct {
	Url     string
	Method  string
	Headers []Header
	Body    []byte
	Client  *http.Client
}

// RequestParam saves information in a format ready to do a http request with
// query parameters
type RequestParam struct {
	Url     string
	Method  string
	Headers []Header
	Body    string
	Client  *http.Client
}

// ListIncidents executes a request and return a list of reported incidents at Statuspage
func (request BaseRequest) ListIncidents(searchQuery string) *http.Response {
	body := ""
	if searchQuery != "" {
		body = "q=" + searchQuery
	}

	finalRequest := RequestParam{
		Url:     request.Url + "/incidents",
		Method:  "GET",
		Headers: request.Headers,
		Body:    body,
		Client:  request.Client,
	}

	res := finalRequest.exec()

	return res
}

// GetIncident retrieve information about a specific incident
func (request BaseRequest) GetIncident(incidentId string) *http.Response {

	finalRequest := RequestParam{
		Url:     request.Url + "/incidents/" + incidentId,
		Method:  "GET",
		Headers: request.Headers,
		Body:    "",
		Client:  request.Client,
	}

	res := finalRequest.exec()

	return res
}

// UpdateIncident Updates one incident information
func (request BaseRequest) UpdateIncident(incidentId string, incidentBody IncidentBody) *http.Response {
	bodyJson, err := json.Marshal(incidentBody)
	if err != nil {
		log.Fatal(err)
	}

	finalRequest := RequestJson{
		Url:     request.Url + "/incidents/" + incidentId,
		Method:  "PATCH",
		Headers: request.Headers,
		Body:    bodyJson,
		Client:  request.Client,
	}

	res := finalRequest.exec()

	return res
}

// func CreateIncident(baseRequest BaseRequest, incidentBody IncidentBody) RequestJson {
//         request := RequestJson{
//                 Url:     baseRequest.Url + "/incidents",
//                 Method:  "POST",
//                 Headers: baseRequest.Headers,
//                 Body:    incidentBody,
//         }

//         return request
// }

func DeleteIncident(baseRequest BaseRequest, incidentId string) RequestParam {
	request := RequestParam{
		Url:     baseRequest.Url + "/incidents/" + incidentId,
		Method:  "DELETE",
		Headers: baseRequest.Headers,
		Body:    "",
	}

	return request
}
