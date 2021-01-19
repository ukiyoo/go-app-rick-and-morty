package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type home struct {
	app.Compo

	ID              int
	CategorySlug    string
	UpdateAvailable bool
}

func (h *home) OnAppUpdate(ctx app.Context) {
	h.UpdateAvailable = ctx.AppUpdateAvailable
	h.Update()
}

func (h *home) onUpdateClick(ctx app.Context, e app.Event) {
	app.Reload()
}

func (h *home) Render() app.UI {
	c := api.GetCategory(h.CategorySlug)

	return app.Div().Body(
		app.Shell().
			Menu(app.Div().
				Body(&menu{CurrentCategory: c, UpdateStatus: h.UpdateAvailable}),
			).
			OverlayMenu(
				app.Div().
					Style("background", "linear-gradient(#2e343a, rgba(0, 0, 0, 0.9))").
					Style("height", "100%").
					Class("overlay-menu").
					Body(&menu{CurrentCategory: c, UpdateStatus: h.UpdateAvailable}),
			).
			Content(app.Main().
				Style("height", "calc(100% - 45px - 12px - 5px)").
				Style("margin", "45px 0 12px").
				// Style("padding", "42px 0 0").
				Style("overflow-x", "hidden").
				Style("overflow-y", "auto").
				Body(
					app.If(c.Slug == CHARACTER,
						&Character{CurrentCategory: c},
					).ElseIf(c.Slug == LOCATION,
						&Location{CurrentCategory: c},
					).ElseIf(c.Slug == EPISODE,
						&Episode{CurrentCategory: c},
					).Else(
						&Search{},
					),
				),
			),
	)
}

func main() {
	for _, s := range api.Slugs() {
		app.Route("/"+s, &home{CategorySlug: s})
	}

	// for i := 1; i <= 671; i++ {
	// 	app.Route("/"+CHARACTER+"/"+strconv.Itoa(i), &home{ID: i})
	// }
	//for i := 1; i <= 108; i++ {
	//	app.Route("/"+LOCATION+"/"+strconv.Itoa(i), &home{ID: i})
	//}
	//for i := 1; i <= 41; i++ {
	//	app.Route("/"+EPISODE+"/"+strconv.Itoa(i), &home{ID: i})
	//}

	app.Route("/", &home{})
	app.Run()
}
