package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	http.Handle("/", &app.Handler{
		Author:      "ukiyoo",
		Title:       "",
		Name:        "",
		Description: "",
		Icon: app.Icon{
			//Default: "/web/rnm-portal.png",
			Default: "/web/giphy.gif",
		},
		ThemeColor:      "#000000",
		BackgroundColor: "#000000",
		Styles: []string{
			"web/loader.css",
			"web/custom.css",
		},
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
