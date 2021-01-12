package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Loader struct {
	app.Compo
	loader bool
}

func (l *Loader) isLoader() {
	l.loader = true
	l.Update()
}

func (l *Loader) Render() app.UI {
	return app.Section().Class("hero is-fullheight").Body(
		app.Div().Class("hero-body has-text-centered").Body(
			app.Div().Class("container").Body(
				app.Div().Class("lds-dual-ring"),
			),
		),
	)
}
