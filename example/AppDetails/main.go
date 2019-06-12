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
		appID:            1222530780,
		markets:          []string{"de"},
		languages:        []string{"de"},
		devices:          []string{"iphone"},
		topCompetitorIds: []int{891535485, 654810212},
		topKeywords: []string{
			"hiit",
			"freeletics",
			"high intensity training",
			"weightloss",
		},
		topNegativeKeywords: []string{},
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
			fmt.Printf("This is a %v", err)
		}
		fmt.Println("error:", err)
		return
	}
	fmt.Println(resp.AD.Title)

}
