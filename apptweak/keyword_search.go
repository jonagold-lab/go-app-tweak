package apptweak

import (
	"encoding/json"
	"errors"
	"net/http"
)

type KeywordTopAppsResponse struct {
	AppList []SearchRequestApp `json:"content"`
	MD      MetaData           `json:"metadata"`
}

type SearchRequestApp struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Icon      string   `json:"icon"`
	Developer string   `json:"developer"`
	Price     string   `json:"price"`
	Genres    []string `json:"-"` //TODO: Both []int and []string threw error, so currently not using Genres here. Needs fixture!
	Devices   []string `json:"devices"`
	Slug      string   `json:"slug"`
	Rating    float64  `json:"rating"`
	InApps    bool     `json:"in_apps"`
}

// KeywordSearch takes in an url in the format /ios/searches.json?term=<keyword> and gives back a list of top apps on that keyword
func (c *Client) KeywordSearch(o Options) (*KeywordTopAppsResponse, error) {
	uri := defaultBaseURL + "ios/searches.json"
	if o.Term == "" {
		return nil, errors.New("Term can't be an empty string")
	}
	if o.Num == 0 {
		return nil, errors.New("Num can't be nil")
	}
	u, err := addOptions(uri, o)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var resp KeywordTopAppsResponse
	err = json.Unmarshal([]byte(b), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
