package incidents

import (
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestListIncidents(t *testing.T) {
        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"

        expectedResult := root.RequestParam{
                Url:    "https://api.statuspage.io/v1/pages/" + pageId + "/incidents",
                Method: "GET",
                Headers: []Header{
                        Header{
                                Name:  "Authorization",
                                Value: "Oauth " + apiKey,
                        },
                },
                Body: "",
        }

        baseRequest := connector.Connect(pageId, apiKey)
        result := ListIncidents(baseRequest, "")

        assert.Equals(t, expectedResult, result)
}

func TestGetIncident(t *testing.T) {

        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
        incidentId := "4FmLlc4e6iQiR7kB6fpDqi1nEZdDuBCo"

        expectedResult := root.RequestParam{
                Url:    "https://api.statuspage.io/v1/pages/" + pageId + "/incidents/" + incidentId,
                Method: "GET",
                Headers: []Header{
                        Header{
                                Name:  "Authorization",
                                Value: "Oauth " + apiKey,
                        },
                },
                Body: "",
        }

        baseRequest := connector.Connect(pageId, apiKey)
        result := GetIncident(baseRequest, incidentId)

        assert.Equals(t, expectedResult, result)
}

func TestUpdateIncident(t *testing.T) {

        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
        testBody := []byte(`{
  "incident": {
    "name": "Test",
    "status": "monitoring",
    "impact_override": "minor",
    "deliver_notifications": true,
    "body": "Incident updated by api",
    "components": {
      "component_id": "operational"
    },
    "component_ids": [
      "5jfwlfsqz7jq"
    ]
  }
}`)

        incidentId := "4FmLlc4e6iQiR7kB6fpDqi1nEZdDuBCo"
        incidentBody := IncidentBody{
                Name: "Test",
                Status: "monitoring",
                Impact: "minor",
                Notification: true,
                Body: "Incident updated by api",
                Components: root.Component{
                        ComponentId: "operational",
                },
                ComponentIds: []string{"5jfwlfsqz7jq"},
        }

        expectedResult := root.RequestJson{
                Url:   "https://api.statuspage.io/v1/pages/" + pageId + "/incidents/" + incidentId,
                Method: "PATCH",
                Headers: []Header{
                        Header{
                                Name:  "Authorization",
                                Value: "Oauth " + apiKey,
                        },
                        Header{
                                Name:  "Content-Type",
                                Value: "application/json",
                        },
                },
                Body: testBody,
        }

        baseRequest := connector.Connect(pageId, apiKey)
        result := UpdateIncident(baseRequest, incidentId, IncidentBody)

        assert.Equals(t, expectedResult, result)

}

func TestCreateIncident(t *testing.T) {

        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
        testBody := []byte(`{
  "incident": {
    "name": "Test",
    "status": "monitoring",
    "impact_override": "minor",
    "deliver_notifications": true,
    "body": "Incident updated by api",
    "components": {
      "component_id": "operational"
    },
    "component_ids": [
      "5jfwlfsqz7jq"
    ]
  }
}`)

        incidentBody := IncidentBody{
                Name: "Test",
                Status: "monitoring",
                Impact: "minor",
                Notification: true,
                Body: "Incident created by api",
                Components: root.Component{
                        ComponentId: "operational",
                },
                ComponentIds: []string{"5jfwlfsqz7jq"}.
        }


        expectedResult := root.RequestJson{
                Url:   "https://api.statuspage.io/v1/pages/" + pageId + "/incidents/" + incidentId,
                Method: "POST",
                Headers: []Header{
                        Header{
                                Name:  "Authorization",
                                Value: "Oauth " + apiKey,
                        },
                        Header{
                                Name:  "Content-Type",
                                Value: "application/json",
                        },
                },
                Body: testBody,
        }

        baseRequest := connector.Connect(pageId, apiKey)
        result := CreatIncident(baseRequest, incidentBody)

        assert.Equals(t, expectedResult, result)
}

func TestDeleteIncident(t *testint.T) {
        pageId := "wiV2d9pz8gdq0xAkNycXVcIEweV8KLw4"
        apiKey := "AsiSTLKioeurneEkdF41Q285y4d5I1sr"
        incidentId := "4FmLlc4e6iQiR7kB6fpDqi1nEZdDuBCo"

        expectedResult := root.RequestParam{
                Url:    "https://api.statuspage.io/v1/pages/" + pageId + "/incidents/" + incidentId,
                Method: "DELETE",
                Headers: []Header{
                        Header{
                                Name:  "Authorization",
                                Value: "Oauth " + apiKey,
                        },
                },
                Body: "",
        }

        baseRequest := connector.Connect(pageId, apiKey)
        result := DeleteIncident(baseRequest, incidentId)

        assert.Equals(t, expectedResult, result)
}
