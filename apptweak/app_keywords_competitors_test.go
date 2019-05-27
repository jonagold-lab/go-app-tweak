package apptweak

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppKeywordsCompetitors(t *testing.T) {
	tests := []struct {
		description     string
		token           string
		appID           int
		urlPath         string
		options         Options
		expectedError   error
		responseCode    int
		fixture         string
		firstCompetitor string
	}{
		{
			description:     "happypath",
			token:           "12345x",
			appID:           1414415906,
			urlPath:         "/ios/applications/1414415906/keywords/competitors.json",
			options:         Options{},
			expectedError:   nil,
			responseCode:    200,
			fixture:         "./fixtures/app_keywords_competitors.json",
			firstCompetitor: "GenM - Marketing Courses",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			responseFile, err := ioutil.ReadFile(tc.fixture)
			if err != nil {
				t.Fatal(err)
			}

			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				// Testing if the requested Path matches the expected Path
				assert.Equal(t, r.URL.Path, tc.urlPath, "Url Path should match")

				w.WriteHeader(tc.responseCode)
				w.Write(responseFile)
				w.Header().Set("Content-Type", setContentType(tc.fixture))
			}))

			defer s.Close()
			u, err := url.Parse(s.URL)
			if err != nil {
				log.Fatalln("failed to parse httptest.Server URL:", err)
			}
			hc := &http.Client{}
			hc.Transport = RewriteTransport{URL: u}

			client := NewAuthClient(tc.token, hc)
			resp, err := client.AppKeywordsCompetitors(tc.appID, tc.options)

			// In case of an expected Error, testing if the returned Error matches the expected Error
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err, tc.description)
				return
			}

			//Handling any unexpected Errors and let test fail
			if err != nil && tc.expectedError == nil {
				t.Errorf("Unexpected Error in %v", tc.description)
			}

			// Testing if Unmarshaling of returned JSON works as expected
			assert.Equal(t, tc.firstCompetitor, resp.CompetitorList[0].Title, tc.description)

			// If case of given options, testing if Options are correctly represented in request params
			if tc.description == "with options" {
				assert.Equal(t, tc.options.Country, resp.MD.Req.Params.Country, tc.description)
				assert.Equal(t, tc.options.Language, resp.MD.Req.Params.Language, tc.description)
				assert.Equal(t, tc.options.Device, resp.MD.Req.Params.Device, tc.description)
			}

		})

	}

}
