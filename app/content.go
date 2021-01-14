package main

import (
	"encoding/json"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"io/ioutil"
	"net/http"
)

type All struct {
	characterList CharacterList
	locationList  LocationList
	episodeList   EpisodeList
}

func (c *Content) getApi(url string) {
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

	var data All

	switch c.Slug {
	case "characters":
		err = json.Unmarshal(b, &data.characterList)
	case "locations":
		err = json.Unmarshal(b, &data.locationList)
	case "episodes":
		err = json.Unmarshal(b, &data.episodeList)
	}

	if err != nil {
		app.Log(err.Error())
		return
	}

	c.updateResponse(data)
}

type Content struct {
	app.Compo

	Url      string
	Slug     string
	Response All
}


func (c *Content) updateResponse(data All) {
	app.Dispatch(func() {
		c.Response = data
		c.Update()
	})
}

func (c *Content) OnMount(ctx app.Context) {
	app.Dispatch(func() {
		c.getApi(c.Url)
	})
}


func (c *Content) Render() app.UI {
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.Div().Class("columns is-multiline").Body(

				app.If(c.Slug == "characters",
					app.Range(c.Response.characterList.Results).Slice(func(i int) app.UI {
						return app.Div().Class("column is-6").Body(
							app.A().Href("http://localhost:8000/chararcter/").Body(
								newCharacterBox().
									Name(c.Response.characterList.Results[i].Name).
									Image(c.Response.characterList.Results[i].Image).
									Species(c.Response.characterList.Results[i].Species).
									Status(c.Response.characterList.Results[i].Status).
									Location(c.Response.characterList.Results[i].Location.Name),
							),
						)
					}),

				).ElseIf(c.Slug == "locations",
					app.Range(c.Response.locationList.Results).Slice(func(i int) app.UI {
						return app.Div().Class("column is-6").Body(
							app.A().Href("http://localhost:8000/chararcter/").Body(
								newLocationBox().
									Name(c.Response.locationList.Results[i].Name).
									Dimension(c.Response.locationList.Results[i].Dimension),
							),
						)
					}),

				).ElseIf(c.Slug == "episodes",
					app.Range(c.Response.episodeList.Results).Slice(func(i int) app.UI {
						return app.Div().Class("column is-6").Body(
							app.A().Href("http://localhost:8000/chararcter/").Body(
								newEpisodeBox().
									Name(c.Response.episodeList.Results[i].Name).
									AirDate(c.Response.episodeList.Results[i].AirDate).
									Episode(c.Response.episodeList.Results[i].Episode),
							),
						)
					}),
				),
			),
		),
	)
}

type characterBox struct {
	app.Compo

	character CharacterDetail
}

func newCharacterBox() *characterBox {
	return &characterBox{}
}

func (c *characterBox) Name(v string) *characterBox {
	c.character.Name = v
	return c
}

func (c *characterBox) Species(v string) *characterBox {
	c.character.Species = v
	return c
}

func (c *characterBox) Image(v string) *characterBox {
	c.character.Image = v
	return c
}

func (c *characterBox) Status(v string) *characterBox {
	c.character.Status = v
	return c
}

func (c *characterBox) Location(v string) *characterBox {
	c.character.Location.Name = v
	return c
}

func (c *characterBox) Render() app.UI {
	return app.Div().Class("box").Body(

		app.Article().Class("media").Body(
			app.Div().Class("media-left").Body(
				app.Figure().Class("image is-128x128").Body(
					app.Img().Class("is-rounded").Src(c.character.Image),
				),
			),

			app.Div().Class("media-content").Body(
				app.Div().Class("content").Body(
					app.P().Body(
						app.Strong().Text(c.character.Name),
						app.Br(),
						app.Small().Text(c.character.Species),
						app.Br(),
						newStatusTag().Text(c.character.Status),
						app.Br(),
						app.Br(),
						app.Small().Class("has-text-grey-light").Text("Last known location: "),
						app.Br(),
						app.Text(c.character.Location),
					),
				),
			),
		),
	)
}

type locationBox struct {
	app.Compo

	location LocationDetail
}

func newLocationBox() *locationBox {
	return &locationBox{}
}

func (l *locationBox) Name(v string) *locationBox {
	l.location.Name = v
	return l
}

func (l *locationBox) Dimension(v string) *locationBox {
	l.location.Dimension = v
	return l
}

func (l *locationBox) Render() app.UI {
	return app.Div().Class("box").Body(

		app.Article().Class("media").Body(
			app.Div().Class("media-content").Body(
				app.Div().Class("content").Body(
					app.P().Body(
						app.Strong().Text(l.location.Name),
						app.Br(),
						app.Small().Class("has-text-grey-light").Text("Description: "),
						app.Br(),
						app.Text(l.location.Dimension),
					),
				),
			),
		),
	)
}

type episodeBox struct {
	app.Compo

	episode EpisodeDetail
}

func newEpisodeBox() *episodeBox {
	return &episodeBox{}
}

func (e *episodeBox) Name(v string) *episodeBox {
	e.episode.Name = v
	return e
}

func (e *episodeBox) AirDate(v string) *episodeBox {
	e.episode.AirDate = v
	return e
}

func (e *episodeBox) Episode(v string) *episodeBox {
	e.episode.Episode = v
	return e
}

func (e *episodeBox) Render() app.UI {
	return app.Div().Class("box").Body(

		app.Article().Class("media").Body(
			app.Div().Class("media-content").Body(
				app.Div().Class("content").Body(
					app.P().Body(
						app.Strong().Text(e.episode.Name),
						app.Br(),
						app.Strong().Text(e.episode.AirDate),
						app.Br(),
						app.Small().Class("has-text-grey-light").Text("Episode: "),
						app.Br(),
						app.Text(e.episode.Episode),
					),
				),
			),
		),
	)
}
