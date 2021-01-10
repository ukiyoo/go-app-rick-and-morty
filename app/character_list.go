package main

import (
	"encoding/json"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"io/ioutil"
	"net/http"
	"time"
)

type character struct {
	app.Compo

	characters AllCharacters
}

type AllCharacters struct {
	Info struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"info"`
	Results []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Species string `json:"species"`
		Type    string `json:"type"`
		Gender  string `json:"gender"`
		Origin  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"origin"`
		Location struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		Image   string    `json:"image"`
		Episode []string  `json:"episode"`
		URL     string    `json:"url"`
		Created time.Time `json:"created"`
	} `json:"results"`
}


func (c *character) getAllCharacters() {
	r, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
		return
	}

	var all AllCharacters

	err = json.Unmarshal(b, &all)
	if err != nil {
		app.Log(err.Error())
		return
	}

	c.updateAllCharacters(all)
}


func (c *character) updateAllCharacters(data AllCharacters) {
	app.Dispatch(func() {
		c.characters = data
		c.Update()
	})
}

func (c *character) OnMount(ctx app.Context) {
	app.Dispatch(c.getAllCharacters)
}

func (c *character) Render() app.UI {
	return app.Div().Class("columns is-multiline").Body(
		app.Range(c.characters.Results).Slice(func(i int) app.UI {
			return app.Div().Class("column is-3").Body(

				app.Div().Class("card").Body(
					app.Div().Class("card-image").Body(
						app.Figure().Class("image is-4by3").Body(
							app.Img().Src(c.characters.Results[i].Image).
								Alt(c.characters.Results[i].Name),
						),
						app.Div().Class("card-content").Body(
							app.Div().Class("media-content").Body(
								app.P().Class("subtitle").Text(c.characters.Results[i].Name),
							),
						),
					),
				),
			)
		}),
	)
}
