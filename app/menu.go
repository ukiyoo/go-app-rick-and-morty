package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type menu struct {
	app.Compo
}

func (c *menu) Render() app.UI {
	return app.Div().Class("section").Body(
		app.Aside().Class("menu").Body(
			app.P().Class("menu-label").Text("General"),
			app.Ul().Class("menu-list").Body(
				app.Li().Body(
					app.A().Href("#").Text("CHARACTERS"),
				),
				app.Li().Body(
					app.A().Href("#").Text("LOCATION"),
				),
				app.Li().Body(
					app.A().Href("#").Text("EPISODES"),
				),
			),
		),
	)
}
