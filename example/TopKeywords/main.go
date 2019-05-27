package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mrboom141/apptweak-go/apptweak"
)

type input struct {
	appID               int
	markets             []string
	languages           []string
	devices             []string
	topCompetitorIds    []int
	topKeywords         []string
	topNegativeKeywords []string
}

func main() {
	app := input{
		appID:     1100673977,
		markets:   []string{"de"},
		languages: []string{"de"},
		devices:   []string{"ipad"},
	}

	opts := apptweak.Options{
		Country:  app.markets[0],
		Language: app.languages[0],
		Device:   app.devices[0],
	}
	token := os.Getenv("APPTWEAK_TOKEN")
	client := apptweak.NewAuthClient(token, &http.Client{})
	resp, err := client.TopKeywords(app.appID, opts)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, keyword := range resp.KeywordList {
		fmt.Println(keyword.Keyword)
	}
}
