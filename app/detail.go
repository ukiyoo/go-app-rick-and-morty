package main

import (
	"encoding/json"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Detail struct {
	app.Compo

	ID             int
	loader         bool
	Response       CharacterDetail
	Episode        EpisodeDetail
	Location       LocationDetail
	OriginLocation LocationDetail
}

func (d *Detail) getLocation(url string) {
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

	d.updateLocation(data)
}

func (d *Detail) getOriginLocation(url string) {
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

	d.updateLocation(data)
}

func (d *Detail) getEpisode(url string) {
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

	d.updateEpisode(data)
}

func (d *Detail) getApi(url string) {
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
	time.AfterFunc(1*time.Second, d.loaderOff) // delete this
	d.updateResponse(data)
}

func (d *Detail) getFirstLocation() {
	for i, s := range d.Response.Episode {
		if i == 0 {
			d.getEpisode(s)
		}
	}
}

func (d *Detail) getFirstEpisode() {
	for i, s := range d.Response.Episode {
		if i == 0 {
			d.getEpisode(s)
		}
	}
}

func (d *Detail) updateOriginLocation(data LocationDetail) {
	d.OriginLocation = data
	d.Update()
}

func (d *Detail) updateLocation(data LocationDetail) {
	d.Location = data
	d.Update()
}

func (d *Detail) updateEpisode(data EpisodeDetail) {
	d.Episode = data
	d.Update()
}

func (d *Detail) updateResponse(data CharacterDetail) {
	d.Response = data
	d.Update()
}

func (d *Detail) OnMount(ctx app.Context) {
	d.getApi("https://rickandmortyapi.com/api/character/"+strconv.Itoa(d.ID))
	d.getLocation(d.Response.Origin.URL)
	d.getOriginLocation(d.Response.Location.URL)
	d.getFirstEpisode()
}

func (d *Detail) loaderOff() {
	d.loader = true
	d.Update()
}

func (d *Detail) loaderOn() {
	d.loader = false
	d.Update()
}

func (d *Detail) Render() app.UI {
	return app.If(!d.loader,
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
										app.Img().Src(d.Response.Image),
									),
								),
							),
							app.Div().Class("tile is-parent").Body(
								app.Article().Class("tile is-child box").Body(
									app.P().Class("subtitle").Text(d.Response.Name),
									app.Text(d.Response.Species+" - "),
									app.Text(d.Response.Gender+" - "),
									newStatusTag().Text(d.Response.Status),
									app.Br(),

									app.Br(),
									app.Small().Class("has-text-grey-light").Text("Origin"),
									app.P().Text(d.Response.Origin.Name),

									app.Br(),
									app.Small().Class("has-text-grey-light").Text("Location"),
									app.P().Text(d.Location.Name),
									app.P().Text(d.Location.Type),
									app.P().Text(d.Location.Dimension),

									app.Br(),
									app.Small().Class("has-text-grey-light").Text("First seen in:"),
									app.P().Text(d.Episode.Name),
									app.P().Text(d.Episode.Episode),
								),
							),
						),
					),
				),
			),
		),
	)
}
