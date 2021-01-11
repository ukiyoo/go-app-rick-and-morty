package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Section().Class("section").Body(

		app.Div().Class("container").Body(
			app.Div().Class("columns").Body(
				app.Div().Class("column is-3").Body(
					&menu{},
				),
				app.Div().Class("column is-9").Body(
					app.Section().Class("info-tiles").Body(
						app.H1().Class("is-size-4 pb-2").Text("CHARACHTERS"),
					),
					&character{},
				),
			),
		),
	)
}

func main() {
	app.Route("/", &home{})
	app.Run()
}
