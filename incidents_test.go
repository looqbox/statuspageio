package statuspageio

import (
        "bytes"
        "encoding/json"
        "io/ioutil"
        "net/http"
        "net/http/httptest"
        "testing"

        "github.com/stretchr/testify/assert"
)

var exampleBody = IncidentBody{
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

const (
        okResponse = `{
                "Status": "ok"
        }`

        pageId     = "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey     = "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
        incidentId = "PHrzHysLcKu0axEWLyg2tsaW7xKWwbRC"
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

        status, _ := statuspage.ListIncidents("")

        assert.Equal(t, "200 OK", status)
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

        status, _ := statuspage.GetIncident("")

        assert.Equal(t, "200 OK", status)
}

func TestUpdateIncident(t *testing.T) {
        h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
                assert.Equal(t, "PATCH", r.Method)
                var body IncidentBody
                responseBody, _ := ioutil.ReadAll(r.Body)
                json.NewDecoder(bytes.NewBuffer(responseBody)).Decode(&body)
                assert.Equal(t, exampleBody, body)
                w.Write([]byte(okResponse))
        })
        testServer := httptest.NewServer(h)

        statuspage := Connect(pageId, apiKey)
        statuspage.Url = "http://" + testServer.Listener.Addr().String()

        status, _ := statuspage.UpdateIncident(incidentId, exampleBody)

        assert.Equal(t, "200 OK", status)
}

func TestCreateIncidents(t *testing.T) {
        h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
                assert.Equal(t, "POST", r.Method)
                var body IncidentBody
                responseBody, _ := ioutil.ReadAll(r.Body)
                json.NewDecoder(bytes.NewBuffer(responseBody)).Decode(&body)
                assert.Equal(t, exampleBody, body)
                w.Write([]byte(okResponse))
        })
        testServer := httptest.NewServer(h)

        statuspage := Connect(pageId, apiKey)
        statuspage.Url = "http://" + testServer.Listener.Addr().String()

        status, _ := statuspage.CreateIncident(exampleBody)

        assert.Equal(t, "200 OK", status)
}

func TestDeleteIncident(t *testing.T) {
        h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
                assert.Equal(t, "DELETE", r.Method)
                w.Write([]byte(okResponse))
        })
        testServer := httptest.NewServer(h)

        statuspage := Connect(pageId, apiKey)
        statuspage.Url = "http://" + testServer.Listener.Addr().String()

        status, _ := statuspage.DeleteIncident("")

        assert.Equal(t, "200 OK", status)
}
