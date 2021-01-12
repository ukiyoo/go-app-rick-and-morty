package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type search struct {
	app.Compo
}

// <div class="field is-grouped">
//   <p class="control is-expanded">
//     <input class="input" type="text" placeholder="Find a repository">
//   </p>
//   <p class="control">
//     <a class="button is-info">
//       Search
//     </a>
//   </p>
// </div>

func (s *search) Render() app.UI {
	return app.Div().Class("field is-grouped").Body(
		app.P().Class("control is-expanded").Body(
			app.Input().Class("input").Type("text"),
		),
		app.P().Class("control").Body(
			app.A().Class("button is-info").Text("Search"),
		),
	)
}
