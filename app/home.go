package main

import (
	"encoding/json"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"io/ioutil"
	"net/http"
	"time"
)

type index struct {
	app.Compo

	loader         bool
	Response       CharacterDetail
	Episode        EpisodeDetail
	Location       LocationDetail
	OriginLocation LocationDetail
}

func (h *index) getLocation(url string) {
	r, err := http.Get(url)
	if err != nil {

	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
	}

	var data LocationDetail
	err = json.Unmarshal(b, &data)
	if err != nil {
		app.Log(err.Error())
	}

	h.updateLocation(data)
}

func (h *index) getOriginLocation(url string) {
	r, err := http.Get(url)
	if err != nil {

	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
	}

	var data LocationDetail
	err = json.Unmarshal(b, &data)
	if err != nil {
		app.Log(err.Error())
	}

	h.updateLocation(data)
}

func (h *index) getEpisode(url string) {
	r, err := http.Get(url)
	if err != nil {

	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
	}

	var data EpisodeDetail
	err = json.Unmarshal(b, &data)
	if err != nil {
		app.Log(err.Error())
	}

	h.updateEpisode(data)
}

func (h *index) getApi(url string) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
		return
	}

	var data CharacterDetail
	err = json.Unmarshal(b, &data)
	if err != nil {
		app.Log(err.Error())
		return
	}
	time.AfterFunc(1*time.Second, h.loaderOff) // delete this
	h.updateResponse(data)
}

func (h *index) getFirstLocation() {
	for i, s := range h.Response.Episode {
		if i == 0 {
			h.getEpisode(s)
		}
	}
}

func (h *index) getFirstEpisode() {
	for i, s := range h.Response.Episode {
		if i == 0 {
			h.getEpisode(s)
		}
	}
}

func (h *index) updateOriginLocation(data LocationDetail) {
	h.OriginLocation = data
	h.Update()
}

func (h *index) updateLocation(data LocationDetail) {
	h.Location = data
	h.Update()
}

func (h *index) updateEpisode(data EpisodeDetail) {
	h.Episode = data
	h.Update()
}

func (h *index) updateResponse(data CharacterDetail) {
	h.Response = data
	h.Update()
}

func (h *index) OnMount(ctx app.Context) {
	h.getApi("https://rickandmortyapi.com/api/character/2")
	h.getLocation(h.Response.Origin.URL)
	h.getOriginLocation(h.Response.Location.URL)
	h.getFirstEpisode()
}

func (h *index) loaderOff() {
	h.loader = true
	h.Update()
}

func (h *index) loaderOn() {
	h.loader = false
	h.Update()
}

func (h *index) Render() app.UI {
	return app.If(!h.loader,
		newLoader(),
	).Else(
		app.Div().Class("section").Body(
			app.Div().Class("container").Body(
				app.Div().Class("tile is-ancestor").Body(
					app.Div().Class("tile is-vertical is-8").Body(
						app.Div().Class("tile").Body(
							app.Div().Class("tile is-parent").Body(
								app.Article().Class("tile is-child card is-primary").Body(
									app.Figure().Class("image").Body(
										app.Img().Src(h.Response.Image),
									),
								),
							),
							app.Div().Class("tile is-parent").Body(
								app.Article().Class("tile is-child box").Body(
									app.P().Class("subtitle").Text("Locations"),
									app.Small().Class("has-text-grey-light").Text("Origin"),
									app.P().Text(h.OriginLocation.Name),
									app.P().Text(h.OriginLocation.Type),
									app.P().Text(h.OriginLocation.Dimension),
									app.Br(),
									app.Small().Class("has-text-grey-light").Text("Now Location"),
									app.P().Text(h.Location.Name),
									app.P().Text(h.Location.Type),
									app.P().Text(h.Location.Dimension),
								),
							),
						),
						app.Div().Class("tile").Body(
							app.Div().Class("tile is-parent").Body(
								app.Article().Class("tile is-child box").Body(
									app.P().Class("subtitle").Text(h.Response.Name),
									app.Text(h.Response.Species+" - "),
									app.Text(h.Response.Gender+" - "),
									newStatusTag().Text(h.Response.Status),
								),
							),
							app.Div().Class("tile is-parent").Body(
								app.Article().Class("tile is-child box").Body(
									app.P().Class("subtitle").Text("First seen in:"),
									app.P().Text(h.Episode.Name),
									app.P().Text(h.Episode.Episode),
									app.P().Text(h.Episode.AirDate),
								),
							),
						),
					),
				),
			),
		),
	)
}
