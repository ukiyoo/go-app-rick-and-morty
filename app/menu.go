package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type menu struct {
	app.Compo
}

func (c *menu) Render() app.UI {
	return app.Div().Body(
		app.Aside().Class("menu is-hedden-mobile").Body(
			app.P().Class("menu-label").Text("Characters"),
			app.Ul().Class("menu-list").Body(
				app.Li().Body(
					app.A().Text("Team Settings"),
				),
			),
			app.P().Class("menu-label").Text("Episodes"),
			app.P().Class("menu-label").Text("Locations"),
		),
	)
}
