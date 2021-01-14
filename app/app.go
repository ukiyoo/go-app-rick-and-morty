package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo

	CategorySlug string
}

func (h *home) Render() app.UI {
	c := cat.Get(h.CategorySlug)

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
			Content(app.Div().
				Style("height", "100%").
				Style("overflow-x", "hidden").
				Style("overflow-y", "auto").
				Body(&Content{Url: c.URL, Slug: c.Slug})),
	)
}

func main() {
	for _, s := range cat.Slugs() {
		app.Route("/"+s, &home{CategorySlug: s})
	}

	// for i := 1; i < 600; i++ {
	// 	app.Route("/character/"+strconv.Itoa(i), &home{id: i, isSingle: true})
	// }
	app.Route("/", &home{})
	app.Run()
}
