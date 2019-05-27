package apptweak

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type AppTopKeywordResponse struct {
	KeywordList []Keyword `json:"content"`
	MD          MetaData  `json:"metadata"`
}

type Keyword struct {
	Keyword string `json:"keyword"`
	Ranking int    `json:"ranking"`
}

// TopKeywords takes in an url in the format ios/applications/<appID>/keywords/top.json and gives back a list of keywords with rankning for the app
func (c *Client) TopKeywords(appID int, o Options) (*AppTopKeywordResponse, error) {
	uri := defaultBaseURL + "ios/applications/" + strconv.Itoa(appID) + "/keywords/top.json"
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
	var resp AppTopKeywordResponse
	err = json.Unmarshal([]byte(b), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, err
}
