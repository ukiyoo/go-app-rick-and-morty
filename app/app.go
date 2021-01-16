package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"strconv"
)

type home struct {
	app.Compo

	ID           int
	CategorySlug string
}

func (h *home) Render() app.UI {
	c := api.GetCategory(h.CategorySlug)

	return app.Div().Body(
		app.Shell().
			Menu(app.Div().
				Body(&menu{CurrentCategory: c}),
			).
			OverlayMenu(
				app.Div().
					Style("background", "linear-gradient(#2e343a, rgba(0, 0, 0, 0.9))").
					Style("height", "100%").
					Class("overlay-menu").
					Body(&menu{CurrentCategory: c}),
			).
			Content(app.If(h.ID == 0,
				app.Div().
					Style("height", "100%").
					Style("overflow-x", "hidden").
					Style("overflow-y", "auto").
					Body(&Content{CurrentCategory: c}),
			).Else(
				app.Div().
					Style("height", "100%").
					Style("overflow-x", "hidden").
					Style("overflow-y", "auto").
					Body(&Detail{ID: h.ID}),
			),
			),
	)
}

func main() {
	for _, s := range api.Slugs() {
		app.Route("/"+s, &home{CategorySlug: s})
	}

	for i := 1; i <= 671; i++ {
		app.Route("/"+CHARACTER+"/"+strconv.Itoa(i), &home{ID: i})
	}
	//for i := 1; i <= 108; i++ {
	//	app.Route("/"+LOCATION+"/"+strconv.Itoa(i), &home{ID: i})
	//}
	//for i := 1; i <= 41; i++ {
	//	app.Route("/"+EPISODE+"/"+strconv.Itoa(i), &home{ID: i})
	//}

	app.Route("/", &home{})
	app.Run()
}
