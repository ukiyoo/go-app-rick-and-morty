package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Characters struct {
	app.Compo

	loader     Loader
	page       int
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

func (c *Characters) getAllCharacters(url string) {
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

	var all AllCharacters

	err = json.Unmarshal(b, &all)
	if err != nil {
		app.Log(err.Error())
		return
	}
	// c.isLoader()
	time.AfterFunc(1*time.Second, c.isLoader)
	c.updateAllCharacters(all)
}

func (c *Characters) updateAllCharacters(data AllCharacters) {
	app.Dispatch(func() {
		c.characters = data
		c.Update()
	})
}

func (c *Characters) isLoader() {
	app.Dispatch(func() {
		c.loader.loader = true
		c.Update()
	})
}

func (c *Characters) OnMount(ctx app.Context) {
	app.Dispatch(func() {
		c.getAllCharacters("https://rickandmortyapi.com/api/character")
	})
}

func (c *Characters) onNext(ctx app.Context, e app.Event) {
	c.loader.loader = false
	c.Update()

	app.Dispatch(func() {
		c.getAllCharacters(c.characters.Info.Next)
	})
}

func (c *Characters) onPrev(ctx app.Context, e app.Event) {
	c.loader.loader = false
	c.Update()

	app.Dispatch(func() {
		c.getAllCharacters(c.characters.Info.Next)
	})
}

func (c *Characters) onPage(ctx app.Context, e app.Event) {
	e.PreventDefault()
	c.loader.loader = false
	c.Update()

	pageInt := ctx.JSSrc.Get("text").String()
	c.page, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/?page=%v", c.page)

	app.Dispatch(func() {
		c.getAllCharacters(url)
	})
}

func (c *Characters) Render() app.UI {
	pages := make([]int, c.characters.Info.Pages)
	return app.If(!c.loader.loader,
		&Loader{},
	).Else(
		app.Div().Class("section").Body(
			app.Div().Class("container").Body(
				app.Div().Class("columns is-multiline").Body(
					app.Range(c.characters.Results).Slice(func(i int) app.UI {
						return app.Div().Class("column is-6").Body(
							app.A().Href("http://localhost:8000/chararcter/").Body(
								app.Div().Class("box").Body(

									app.Article().Class("media").Body(
										app.Div().Class("media-left").Body(
											app.Figure().Class("image is-128x128").Body(
												app.Img().Class("is-rounded").Src(c.characters.Results[i].Image),
											),
										),

										app.Div().Class("media-content").Body(
											app.Div().Class("content").Body(
												app.P().Body(
													app.Strong().Text(c.characters.Results[i].Name),
													app.Br(),
													app.Small().Text(c.characters.Results[i].Species),
													app.Br(),
													app.Text(c.characters.Results[i].Status),
												),
											),
										),
									),
								),
							),
						)
					}),
				),
				app.Nav().Class("pagination is-centered").Body(
					app.A().Href("/").Class("pagination-previous").Text("Prev").OnClick(c.onPrev),
					app.A().Href("/").Class("pagination-next").Text("Next").OnClick(c.onNext),

					app.Ul().Class("pagination-list").Body(
						app.Range(pages).Slice(func(i int) app.UI {
							i++
							return app.Li().Body(
								app.If(i == c.page,
									app.A().Class("pagination-link is-current").Href("/").Text(i).OnClick(c.onPage),
								).Else(
									app.A().Class("pagination-link").Href("/").Text(i).OnClick(c.onPage),
								),
							)
						}),
					),
				),
			),
		),
	)
}
