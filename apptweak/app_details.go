package apptweak

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AppDetailResponse struct {
	AD AppDetail `json:"content"`
	MD MetaData  `json:"metadata"`
}

type ErrorResponse struct {
	Err           string `json:"error"`
	ApplicationID int    `json:"application_id,omitempty"`
	Device        string `json:"device,omitempty"`
	Country       string `json:"country,omitempty"`
	Language      string `json:"language,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Err
}

type AppDetail struct {
	AppID               int         `json:"id"`
	Dev                 Developer   `json:"developer"`
	Ratings             Rating      `json:"rating"`
	Description         string      `json:"description"`
	Feats               Features    `json:"features"`
	Icon                string      `json:"icon"`
	Genres              []int       `json:"genres"`
	Permissions         []string    `json:"permissions"`
	Price               string      `json:"price"`
	Size                int         `json:"size"`
	PromotionalText     string      `json:"promotionalText"`
	Screens             Screenshots `json:"screenshots"`
	Vids                Videos      `json:"videos"`
	Slug                string      `json:"slug"`
	Title               string      `json:"title"`
	SubTitle            string      `json:"subtitle"`
	Versions            []Version   `json:"versions"`
	ReleaseDate         string      `json:"release_date"`
	Devices             []string    `json:"devices"`
	CustomersAlsoBought []string    `json:"customers_also_bought"`
}
type Developer struct {
	Name        string `json:"name"`
	DeveloperID int    `json:"id"`
}

type Rating struct {
	Average float64 `json:"average"`
}

type Features struct {
	GameCenter bool `json:"game_center"`
	Passbook   bool `json:"passbook"`
	InApps     bool `json:"in:apps"`
}

type Screenshots struct {
	IPhone       []ScreenshotDetail `json:"iphone"`
	IPhone5      []ScreenshotDetail `json:"iphone5"`
	IPhone6      []ScreenshotDetail `json:"iphone6"`
	IPhone6AndUp []ScreenshotDetail `json:"iphone6+"`
	IPad         []ScreenshotDetail `json:"ipad"`
	IPadPro      []ScreenshotDetail `json:"ipadPro"`
	IPhone6Plus  []ScreenshotDetail `json:"iphone6plus"`
	Applewatch   []ScreenshotDetail `json:"appleWatch"`
}

type ScreenshotDetail struct {
	ID            string `json:"id"`
	PathComponent string `json:"path_component"`
	Filename      string `json:"filename"`
	URL           string `json:"url"`
}

// Videos TODO: quickfix use an iterface since iphone6+ can be {} or []
type Videos struct {
	IPhone6AndUp interface{} `json:"iphone6+"`
}

type VideoDetail struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URI    string `json:"uri"`
	Codecs string `json:"codecs"`
	Audio  string `json:"audio"`
}

type Version struct {
	ReleaseDate  string `json:"release_date"`
	ReleaseNotes string `json:"-"` // `json:"release_notes"` TODO: Problem was that release_notes was not a string, prob due to symbols etc. Needs fixture!
	Version      string `json:"version"`
}

type MetaData struct {
	Req     Request         `json:"request"`
	Content MetaDataContent `json:"content"`
}

type Request struct {
	Path        string     `json:"path"`
	Store       string     `json:"store"`
	Params      Parameters `json:"params"`
	PerformedAt string     `json:"performed_at"`
}

type Parameters struct {
	Country  string `json:"country"`
	Language string `json:"language"`
	Device   string `json:"device"`
	ID       string `json:"id"`
	Format   string `json:"format"`
	Num      int    `json:"num"`
	Term     string `json:"term"`
	//TODO: Add MaxAge
}

type MetaDataContent struct {
	Content string `json:"-"`
}

// AppDetails takes appID as string and a struct of type Options such as Country, Device & Language and gives back full Meta Data of the app
func (c *Client) AppDetails(appID int, o Options) (*AppDetailResponse, error) {
	uri := defaultBaseURL + "ios/applications/" + strconv.Itoa(appID) + "/metadata.json"
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

	var resp AppDetailResponse
	err = json.Unmarshal([]byte(b), &resp)
	if err != nil {
		return nil, fmt.Errorf("Error: %v | Response of Request: %s", err, string(b))
	}
	return &resp, nil
}
