package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"strconv"
)

type home struct {
	app.Compo

	id       int
	isSingle bool
}

func (h *home) Render() app.UI {

	return app.Div().Body(
		// &menu{},
		// app.Div().Class("container").Body(
		// 	app.Div().Class("columns").Body(
		// 		app.Div().Class("column is-3").Body(),
		// 		app.Div().Class("column is-9").Body(
		// 			// &search{},
		// 			app.Section().Class("info-tiles").Body(
		// 				app.H1().Class("is-size-4 pb-2").Text("CHARACHTERS"),
		// 			),
		// 			app.If(h.isSingle,
		// 				&Character{id: h.id},
		// 			).Else(
		// 				&Characters{},
		// 			),
		// 		),
		// 	),
		// ),
		app.Shell().
			Menu(app.Div().
				Body(&menu{}),
			).
			OverlayMenu(
				app.Div().
					Style("background", "linear-gradient(#2e343a, rgba(0, 0, 0, 0.9))").
					Style("height", "100%").
					Class("overlay-menu").
					Body(&menu{}),
			).
			Content(app.Div().
				Style("height", "100%").
				Style("overflow-x", "hidden").
				Style("overflow-y", "auto").
				Body(&Characters{})),
	)
}

func main() {
	for i := 1; i < 600; i++ {
		app.Route("/character/"+strconv.Itoa(i), &home{id: i, isSingle: true})
	}
	app.Route("/", &home{})
	app.Run()
}
