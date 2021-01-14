package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Content struct {
	app.Compo

	Url string
}

func (c *Content) Render() app.UI {

	return app.Section().Class("section").Body(

		app.Div().Class("box").Body(

			app.Article().Class("media").Body(
				app.Div().Class("media-left").Body(
					app.Figure().Class("image is-128x128").Body(),
				),

				app.Div().Class("media-content").Body(
					app.Div().Class("content").Body(
						app.P().Body(

							app.Br(),
							app.Text("Hello"),
							app.Br(),
						),
					),
				),
			),
		),
	)
}
