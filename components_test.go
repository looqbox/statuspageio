package statuspageio

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	componentID = "PHrzHysLcKu0axEWLyg2tsaW7xKWwbRC"
)

func TestListComponents(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "OAuth "+apiKey, r.Header.Get("Authorization"))
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(okResponse))
	})
	testServer := httptest.NewServer(h)

	statuspage := Connect(pageID, apiKey)
	statuspage.URL = "http://" + testServer.Listener.Addr().String()

	status, _ := statuspage.ListComponents()

	assert.Equal(t, "200 OK", status)
}
