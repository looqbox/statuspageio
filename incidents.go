package statuspageio

import (
        "encoding/json"
        "io/ioutil"
        "log"
)

// ListIncidents executes a request and return a list of reported incidents at Statuspage
func (request BaseRequest) ListIncidents(searchQuery string) (string, []IncidentsResponse) {
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

        jsonData, err := ioutil.ReadAll(res.Body)
        if err != nil {
                log.Fatal(err)
        }

        var structuredResponse []IncidentsResponse
        json.Unmarshal([]byte(jsonData), &structuredResponse)

        return res.Status, structuredResponse
}

// GetIncident retrieve information about a specific incident
func (request BaseRequest) GetIncident(incidentId string) (string, IncidentsResponse) {

        finalRequest := RequestParam{
                Url:     request.Url + "/incidents/" + incidentId,
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

        var structuredResponse IncidentsResponse
        json.Unmarshal([]byte(jsonData), &structuredResponse)

        return res.Status, structuredResponse
}

// UpdateIncident Updates one incident information
func (request BaseRequest) UpdateIncident(incidentId string, incidentBody IncidentBody) (string, IncidentsResponse) {

        finalRequest := RequestJson{
                Url:     request.Url + "/incidents/" + incidentId,
                Method:  "PATCH",
                Headers: request.Headers,
                Body:    incidentBody,
                Client:  request.Client,
        }

        res := finalRequest.exec()

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
        finalRequest := RequestJson{
                Url:     request.Url + "/incidents",
                Method:  "POST",
                Headers: request.Headers,
                Body:    incidentBody,
                Client:  request.Client,
        }

        res := finalRequest.exec()

        jsonData, err := ioutil.ReadAll(res.Body)
        if err != nil {
                log.Fatal(err)
        }

        var structuredResponse IncidentsResponse
        json.Unmarshal([]byte(jsonData), &structuredResponse)

        return res.Status, structuredResponse
}

// DeleteIncident deletes an incident from the page
func (request BaseRequest) DeleteIncident(incidentId string) (string, IncidentsResponse) {
        finalRequest := RequestParam{
                Url:     request.Url + "/incidents/" + incidentId,
                Method:  "DELETE",
                Headers: request.Headers,
                Body:    "",
                Client:  request.Client,
        }

        res := finalRequest.exec()

        jsonData, err := ioutil.ReadAll(res.Body)
        if err != nil {
                log.Fatal(err)
        }

        var structuredResponse IncidentsResponse
        json.Unmarshal([]byte(jsonData), &structuredResponse)

        return res.Status, structuredResponse
}
