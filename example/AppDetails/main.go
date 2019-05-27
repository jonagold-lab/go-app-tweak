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
		appID:     1,
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

	options := apptweak.Options{
		Country:  app.markets[0],
		Language: app.languages[0],
		Device:   app.devices[0],
	}
	token := os.Getenv("APPTWEAK_TOKEN")
	hc := &http.Client{}
	client := apptweak.NewAuthClient(token, hc)
	resp, err := client.AppDetails(app.appID, options)
	if err != nil {
		if err.Error() == "Response Error" {
			fmt.Printf("This is a %v: %v", err)
		}
		fmt.Println("error:", err)
		return
	}
	fmt.Println(resp.AD.Title)

}
