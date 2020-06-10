package incidents

import (
	"github.com/statuspage/pkg/root"
)

func ListIncidents(baseRequest root.BaseRequest, searchQuery string) root.RequestParam {

	body := ""
	if searchQuery != "" {
		body = "q=" + searchQuery
	}

	request := root.RequestParam{
		Url:     baseRequest.Url + "/incidents",
		Method:  "GET",
		Headers: baseRequest.Headers,
		Body:    body,
	}

	return request
}

func GetIncident(baseRequest root.BaseRequest, incidentId string) root.RequestParam {

	request := root.RequestParam{
		Url:     baseRequest.Url + "/incidents/" + incidentId,
		Method:  "GET",
		Headers: baseRequest.Headers,
		Body:    "",
	}

	return request
}

func UpdateIncident(baseRequest root.BaseRequest, incidentId string, incidentBody root.IncidentBody) root.RequestJson {

	request := root.RequestJson{
		Url:     baseRequest.Url + "/incidents/" + incidentId,
		Method:  "PATCH",
		Headers: baseRequest.Headers,
		Body:    incidentBody,
	}

	return request
}
