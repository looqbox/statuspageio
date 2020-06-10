package incidents

import (
        "github.com/statuspage/pkg/root"
)

type IncidentBody struct {
        Name          string    `json:"name"`
        Status        string    `json:"status"`
        Impact        string    `json:"impact_override"`
        Notification  bool      `json:"deliver_notifications"`
        Body          string    `json:"body"`
        Components    Component `json:"components"`
        ComponentsIds []string  `json:"component_ids"`
}

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

func UpdateIncident(baseRequest root.BaseRequest, incidentId string, incidentBody IncidentBody) root.RequestJson {

        request := root.RequestJson{
                Url:     baseRequest.Url + "/incidents/" + incidentId,
                Method:  "PATCH",
                Headers: baseRequest.Headers,
                Body:    incidentBody,
        }

        return request
}

func CreateIncident(baseRequest root.BaseRequest, incidentBody IncidentBody) root.RequestJson {
        request := root.RequestJson{
                Url:     baseRequest.Url + "/incidents",
                Method:  "POST",
                Headers: baseRequest.Headers,
                Body:    incidentBody,
        }

        return request
}

func DeleteIncident(baseRequest root.BaseRequest, incidentId string) root.RequestParam {
        request := root.RequestParam{
                Url:     baseRequest.Url + "/incidents/" + incidentId,
                Method:  "DELETE",
                Headers: baseRequest.Headers,
                Body:    "",
        }

        return request
}
