package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type loader struct {
	app.Compo

	loader bool
}

func (h *loader) Render() app.UI {
	return app.Div().Class("container").Body(
		app.Div().Class("lds-dual-ring is-vcentered"),
	)
}
