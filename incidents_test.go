package statuspageio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pageId = "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
var apiKey = "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
var incidentId = "PHrzHysLcKu0axEWLyg2tsaW7xKWwbRC"

const (
	okResponse = `{
                "status": "ok"
        }`
)

func TestListIncidents(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(okResponse))
	})
	testServer := httptest.NewServer(h)

	statuspage := Connect(pageId, apiKey)
	statuspage.Url = "http://" + testServer.Listener.Addr().String()

	result := statuspage.ListIncidents("")

	assert.Equal(t, "200 OK", result.Status)
}

func TestGetIncidents(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(okResponse))
	})
	testServer := httptest.NewServer(h)

	statuspage := Connect(pageId, apiKey)
	statuspage.Url = "http://" + testServer.Listener.Addr().String()

	result := statuspage.GetIncident("")

	assert.Equal(t, "200 OK", result.Status)
}

func TestUpdateIncident(t *testing.T) {
	exampleBody := IncidentBody{
		Name:         "Test Incident",
		Status:       "Investigating",
		Impact:       "Minor",
		Notification: false,
		Body:         "Test Body",
		Components: Component{
			ComponentId: "nd83293he9hr3",
		},
		ComponentsIds: []string{"nfoeriu038d", "fnwe8789f"},
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
		assert.Equal(t, "PATCH", r.Method)
		var body IncidentBody
		responseBody, _ := ioutil.ReadAll(r.Body)
		json.NewDecoder(bytes.NewBuffer(responseBody)).Decode(&body)
		log.Println(body)
		log.Println(exampleBody)
		assert.Equal(t, exampleBody, body)
		w.Write([]byte(okResponse))
	})
	testServer := httptest.NewServer(h)

	statuspage := Connect(pageId, apiKey)
	statuspage.Url = "http://" + testServer.Listener.Addr().String()

	result := statuspage.UpdateIncident(incidentId, exampleBody)

	assert.Equal(t, "200 OK", result.Status)
}
