package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Character struct {
	app.Compo

	id     int
	char   Char
	loader Loader
}

type Char struct {
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
}

func (c *Character) getCharacter(id int) {
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/%v", id)
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

	var char Char

	err = json.Unmarshal(b, &char)
	if err != nil {
		app.Log(err.Error())
		return
	}
	// c.isLoader()
	time.AfterFunc(1*time.Second, c.loader.isLoader)
	c.updateChar(char)
}

func (c *Character) updateChar(data Char) {
	app.Dispatch(func() {
		c.char = data
		c.Update()
	})
}

func (c *Character) OnMount(ctx app.Context) {
	app.Dispatch(func() {
		c.getCharacter(c.id)
	})
}

func (c *Character) Render() app.UI {

	return app.Section().Class("section").Body(

		app.Div().Class("box").Body(

			app.Article().Class("media").Body(
				app.Div().Class("media-left").Body(
					app.Figure().Class("image is-128x128").Body(
						app.Img().Class("is-rounded").Src(c.char.Image),
					),
				),

				app.Div().Class("media-content").Body(
					app.Div().Class("content").Body(
						app.P().Body(
							app.Strong().Text(c.char.Name),
							app.Br(),
							app.Small().Text(c.char.Species),
							app.Br(),
							app.Text(c.char.Status),
						),
					),
				),
			),
		),
	)
}
