package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	http.Handle("/", &app.Handler{
		Author:      "ukiyoo",
		Title:       "",
		Name:        "",
		Description: "",
		//Icon: app.Icon{
		//	Default: "/web/icon.png",
		//},
		ThemeColor:      "#000000",
		BackgroundColor: "#000000",
		Styles: []string{
			"https://unpkg.com/bulma@0.9.0/css/bulma.min.css",
		},
	})
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
