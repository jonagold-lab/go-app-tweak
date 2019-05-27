package apptweak

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.apptweak.com/"
)

type URL struct {
	Url string
}

type Client struct {
	token      string
	httpClient *http.Client
}

type Options struct {
	Country  string `url:"country,omitempty"`
	Language string `url:"language,omitempty"`
	Device   string `url:"device,omitempty"`
	MaxAge   int    `url:"maxage,omitempty"`
	Num      int    `url:"num,omitempty"`
	Term     string `url:"term,omitempty"`
}

// NewAuthClient takes in a token string and generates a Client to use for all requests
func NewAuthClient(token string, httpClient *http.Client) *Client {
	return &Client{
		token:      token,
		httpClient: httpClient,
	}
}

func (c *Client) checkResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	if resp.StatusCode == http.StatusForbidden {
		return errors.New("Unknown token")
	}
	var errResp ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return err
	}
	return &errResp
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("X-Apptweak-Key", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := c.checkResponse(resp); err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func addOptions(s string, o interface{}) (string, error) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(o)
	if err != nil {
		return s, err
	}
	u.RawQuery = qs.Encode()
	return u.String(), nil
}
