package apptweak

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type AppKeywordsCompetitorsResponse struct {
	CompetitorList []App    `json:"content"`
	MD             MetaData `json:"metadata"`
}

type App struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Icon    string  `json:"icon"`
	Genres  []int   `json:"genres"`
	Price   string  `json:"price"`
	Rating  float64 `json:"rating"`
	Version string  `json:"version"`
}

// AppKeywordsCompetitors takes in an url in the format /ios/applications/<appID>/keywords/competitors.json?params and gives back list of competitors
func (c *Client) AppKeywordsCompetitors(appID int, o Options) (*AppKeywordsCompetitorsResponse, error) {
	uri := defaultBaseURL + "ios/applications/" + strconv.Itoa(appID) + "/keywords/competitors.json"
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

	var resp AppKeywordsCompetitorsResponse
	err = json.Unmarshal([]byte(b), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
