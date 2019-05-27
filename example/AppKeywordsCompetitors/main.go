package main

import (
	"fmt"
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

	params := apptweak.Parameters{
		Country:  app.markets[0],
		Language: app.languages[0],
		Device:   app.devices[0],
	}
	token := os.Getenv("APPTWEAK_TOKEN")
	client := apptweak.NewAuthClient(token)
	resp, err := client.AppKeywordsCompetitors(app.appID, params)

	if err != nil {
		fmt.Println("error:", err)
	}
	for _, app := range resp.CompetitorList {
		fmt.Println(app.Title)

	}
}
