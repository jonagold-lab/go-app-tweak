package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jonagold-lab/go-apptweak/apptweak"
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
		appID:     1414415906,
		markets:   []string{"us"},
		languages: []string{"en"},
		devices:   []string{"iphone"},
		topCompetitorIds: []int{
			1084807225,
			562413829,
			431748264,
			819700936,
			568839295,
		},
		topKeywords: []string{
			"Micro-Learning",
			"learning",
			"mobile learning",
			"learn new skills",
			"learn faster",
			"video learning",
			"video classes",
		},
		topNegativeKeywords: []string{
			"language learning",
			"coding learning",
			"audio",
			"fitness",
			"workout",
			"cooking",
			"acting",
		},
	}

	params := apptweak.Options{
		Country:  app.markets[0],
		Language: app.languages[0],
		Device:   app.devices[0],
		Term:     app.topKeywords[0],
		Num:      10,
	}
	token := os.Getenv("APPTWEAK_TOKEN")
	client := apptweak.NewAuthClient(token, &http.Client{})
	resp, err := client.KeywordSearch(params)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, app := range resp.AppList {
		fmt.Println(app.Title)
	}
}
