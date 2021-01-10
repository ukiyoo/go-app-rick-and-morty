package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Div().Class("container").Body(
		app.Div().Class("columns").Body(
			app.Div().Class("column is-3").Body(
				app.Aside().Class("menu is-hedden-mobile").Body(
					app.P().Class("menu-label").Text("General"),
					app.Ul().Class("menu-list").Body(
						app.Li().Body(
							app.A().Text("Team Settings"),
						),
					),
					app.P().Class("menu-label").Text("Dashboard"),
					app.P().Class("menu-label").Text("Administration"),
				),
			),
			app.Div().Class("column is-9").Body(
				app.Section().Class("hero is-info welcome is-small").Body(
					app.Div().Class("hero-body").Body(
						app.Div().Class("container").Body(
							app.H1().Text("Hello, Admin."),
						),
					),
				),
				app.Section().Class("info-tiles").Body(
					app.Div().Class("tile is-ancestor has-text-centered").Body(
						app.Div().Class("tile is-parent").Body(
							app.Article().Class("tile is-child box").Body(
								app.P().Class("title").Text("439k"),

							),
						),
					),
				),

				app.Div().Class("columns").Body(
					app.Div().Class("column is-6").Body(
						app.Div().Class("card events-card").Body(
							app.Header().Class("card-header").Body(
								app.P().Class("card-header-title").Text("Events"),
								app.Div().Class("card-table").Body(
									app.Div().Class("content").Body(),
								),
							),
						),
					),
					app.Div().Class("column is-one-quarter").Body(
						app.Div().Class("card events-card").Body(
							app.Header().Class("card-header").Body(
								app.P().Class("card-header-title").Text("Events"),
								app.Div().Class("card-table").Body(
									app.Div().Class("content").Body(),
								),
							),
						),
					),
				),
				&character{},
			),
		),
		//&character{},
	)
}

func main() {
	app.Route("/", &home{})
	app.Run()
}
