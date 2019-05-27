package apptweak

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppDetails(t *testing.T) {
	tests := []struct {
		description   string
		token         string
		appID         int
		urlPath       string
		options       Options
		expectedError error
		responseCode  int
		fixture       string
		apptitle      string
	}{
		{
			description:   "apptweak api success",
			token:         "12345x",
			appID:         1414415906,
			urlPath:       "/ios/applications/1414415906/metadata.json",
			options:       Options{},
			expectedError: nil,
			responseCode:  200,
			fixture:       "./fixtures/app_meta.json",
			apptitle:      "Playbook: Learn Skills Faster",
		},
		{
			description:   "with options",
			token:         "12345x",
			appID:         1414415906,
			urlPath:       "/ios/applications/1414415906/metadata.json",
			options:       Options{Country: "us", Language: "en", Device: "iphone"},
			expectedError: nil,
			responseCode:  200,
			fixture:       "./fixtures/app_meta.json",
			apptitle:      "Playbook: Learn Skills Faster",
		},
		{
			description:   "Error: wrong appid",
			token:         "12345x",
			appID:         0,
			urlPath:       "/ios/applications/0/metadata.json",
			options:       Options{Device: "iphone", Language: "us", Country: "us"},
			expectedError: &ErrorResponse{"app unavailable", 0, "iphone", "us", "us"},
			responseCode:  404,
			fixture:       "./fixtures/app_meta_app_unavailable.json",
			apptitle:      "",
		},

		{
			description:   "Error: unknown token",
			token:         "",
			appID:         1414415906,
			urlPath:       "/ios/applications/1414415906/metadata.json",
			options:       Options{},
			expectedError: errors.New("Unknown token"),
			responseCode:  403,
			fixture:       "./fixtures/unknown_token.html",
			apptitle:      "",

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
			resp, err := client.AppDetails(tc.appID, tc.options)


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
			assert.Equal(t, tc.apptitle, resp.AD.Title, tc.description)

			// If case of given options, testing if Options are correctly represented in request params
			if tc.description == "with options" {
				assert.Equal(t, tc.options.Country, resp.MD.Req.Params.Country, tc.description)
				assert.Equal(t, tc.options.Language, resp.MD.Req.Params.Language, tc.description)
				assert.Equal(t, tc.options.Device, resp.MD.Req.Params.Device, tc.description)
			}

		})

	}

}
