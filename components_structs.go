package statuspageio

import (
	"time"
)

// ComponentsResponse stores response information for Components requests
type ComponentsResponse struct {
	ID                 string    `json:"id"`
	PageID             string    `json:"page_id"`
	GroupID            string    `json:"group_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Group              bool      `json:"group"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Position           int       `json:"position"`
	Status             string    `json:"status"`
	Showcase           bool      `json:"showcase"`
	OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
	AutomationEmail    string    `json:"automation_email"`
}
