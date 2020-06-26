package statuspageio

import (
	"net/http"
	"time"
)

type Incident struct {
	Incident IncidentBody `json:"incident"`
}

// IncidentBody saves information in a format required by statuspage.io API to
// execute requests to interact with incidents at the site
type IncidentBody struct {
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	Impact        string    `json:"impact_override"`
	Notification  bool      `json:"deliver_notifications"`
	Body          string    `json:"body"`
	Components    Component `json:"components"`
	ComponentsIDs []string  `json:"component_ids"`
}

// Component mounts a nested json to be used at IncidentBody
type Component struct {
	ComponentID string `json:"component_id"`
}

// Request is a general interface used for requests
type Request interface {
	exec() *http.Response
}

// RequestFormat saves information in a format reade to do a http request to
// Statuspage IO
type RequestFormat struct {
	URL     string
	Method  string
	Headers []Header
	Body    interface{}
	Client  *http.Client
}

// IncidentsResponse stores response information for Incidents requests
type IncidentsResponse struct {
	Name                    string    `json:"name"`
	Status                  string    `json:"status"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	MonitoringAt            time.Time `json:"monitoring_at"`
	ResolvedAt              time.Time `json:"resolved_at"`
	Impact                  string    `json:"impact"`
	Shortlink               string    `json:"shortlink"`
	ScheduledFor            time.Time `json:"scheduled_for"`
	ScheduledUntil          time.Time `json:"scheduled_until"`
	ScheduledRemindPrior    bool      `json:"scheduled_remind_prior"`
	ScheduledRemindedAt     time.Time `json:"scheduled_reminded_at"`
	ImpactOverride          string    `json:"impact_override"`
	ScheduledAutoInProgress bool      `json:"scheduled_auto_in_progress"`
	ScheduledAutoCompleted  bool      `json:"scheduled_auto_completed"`
	Metadata                struct {
	} `json:"metadata"`
	StartedAt       time.Time `json:"started_at"`
	ID              string    `json:"id"`
	PageID          string    `json:"page_id"`
	IncidentUpdates []struct {
		Status               string    `json:"status"`
		Body                 string    `json:"body"`
		CreatedAt            time.Time `json:"created_at"`
		WantsTwitterUpdate   bool      `json:"wants_twitter_update"`
		TwitterUpdatedAt     time.Time `json:"twitter_updated_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		DisplayAt            time.Time `json:"display_at"`
		DeliverNotifications bool      `json:"deliver_notifications"`
		TweetID              string    `json:"tweet_id"`
		ID                   string    `json:"id"`
		IncidentID           string    `json:"incident_id"`
		CustomTweet          string    `json:"custom_tweet"`
		AffectedComponents   []struct {
			Code      string `json:"code"`
			Name      string `json:"name"`
			OldStatus string `json:"old_status"`
			NewStatus string `json:"new_status"`
		} `json:"affected_components"`
	} `json:"incident_updates"`
	PostmortemBody                string    `json:"postmortem_body"`
	PostmortemBodyLastUpdatedAt   time.Time `json:"postmortem_body_last_updated_at"`
	PostmortemIgnored             bool      `json:"postmortem_ignored"`
	PostmortemPublishedAt         bool      `json:"postmortem_published_at"`
	PostmortemNotifiedSubscribers bool      `json:"postmortem_notified_subscribers"`
	PostmortemNotifiedTwitter     bool      `json:"postmortem_notified_twitter"`
	Components                    []struct {
		Status             string    `json:"status"`
		Name               string    `json:"name"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Position           int       `json:"position"`
		Description        string    `json:"description"`
		Showcase           bool      `json:"showcase"`
		ID                 string    `json:"id"`
		PageID             string    `json:"page_id"`
		GroupID            string    `json:"group_id"`
		Group              bool      `json:"group"`
		OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		AutomationEmail    string    `json:"automation_email"`
	} `json:"components"`
}
