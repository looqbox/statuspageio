package root

import (
	"encoding/json"
)

type BaseRequest struct {
	Url     string
	Headers []Header
}

type RequestJson struct {
	Url     string
	Method  string
	Headers []Header
	Body    *json.RawMessage
}

type RequestParam struct {
	Url     string
	Method  string
	Headers []Header
	body    string
}

type Header struct {
	Name  string
	Value string
}

type IncidentBody struct {
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	Impact        string    `json:"impact_override"`
	Notification  bool      `json:"deliver_notifications"`
	Body          string    `json:"body"`
	Components    Component `json:"components"`
	ComponentsIds []string  `json:"component_ids"`
}

type Component struct {
	ComponentId string `json:"component_id"`
}
