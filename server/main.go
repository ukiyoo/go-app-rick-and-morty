package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	//os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	http.Handle("/", &app.Handler{
		Author:      "ukiyoo",
		Title:       "Rick and Morty",
		Name:        "Rick and Morty",
		Description: "Rick and Morty API",
		Icon: app.Icon{
			//Default: "/web/rnm-portal.png",
			Default: "/web/giphy.gif",
		},
		ThemeColor:      "#1a1c1f",
		BackgroundColor: "#1a1c1f",
		LoadingLabel: "...Welcome To The Club, Pal.",
		Styles: []string{
			"web/loader.css",
			"web/custom.css",
		},
		Version: "0.1.1",
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
