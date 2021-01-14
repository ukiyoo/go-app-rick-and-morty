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

type CharacterList struct {
	app.Compo

	Url string

	loader     Loader
	page       int
	characters Characters
}

type Characters struct {
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

func (c *CharacterList) getAllCharacters(url string) {
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

	var data map[string]interface{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		app.Log(err.Error())
		return
	}

	info := data["info"].(map[string]interface{})
	fmt.Println(info["count"])

	// fmt.Println(dat)
	// fmt.Println(data)
	// c.isLoader()
	time.AfterFunc(1*time.Second, c.isLoader)
	// c.updateAllCharacters(all)
}

func (c *CharacterList) updateAllCharacters(data Characters) {
	app.Dispatch(func() {
		c.characters = data
		c.Update()
	})
}

func (c *CharacterList) isLoader() {
	app.Dispatch(func() {
		c.loader.loader = true
		c.Update()
	})
}

func (c *CharacterList) OnMount(ctx app.Context) {
	app.Dispatch(func() {
		c.getAllCharacters(c.Url)
	})
}

func (c *CharacterList) onNext(ctx app.Context, e app.Event) {
	c.loader.loader = false
	c.Update()

	app.Dispatch(func() {
		c.getAllCharacters(c.characters.Info.Next)
	})
}

func (c *CharacterList) onPrev(ctx app.Context, e app.Event) {
	c.loader.loader = false
	c.Update()

	app.Dispatch(func() {
		c.getAllCharacters(c.characters.Info.Next)
	})
}

func (c *CharacterList) onPage(ctx app.Context, e app.Event) {
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

func (c *CharacterList) Render() app.UI {
	pages := make([]int, c.characters.Info.Pages)
	return app.If(!c.loader.loader,
		&Loader{},
	).Else(
		app.Div().Class("section").Body(
			app.Div().Class("container").Body(

				//app.If(c.page < 0,
				//	newCharacterBox(),
				//).Else(

				app.Div().Class("columns is-multiline").Body(
					app.Range(c.characters.Results).Slice(func(i int) app.UI {
						return app.Div().Class("column is-6").Body(
							app.A().Href("http://localhost:8000/chararcter/").Body(
								newCharacterBox().
									Name(c.characters.Results[i].Name).
									Image(c.characters.Results[i].Image).
									Species(c.characters.Results[i].Species).
									Status(c.characters.Results[i].Status).
									Location(c.characters.Results[i].Location.Name),
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

type statusTag struct {
	app.Compo
	color  string
	status string
}

func newStatusTag() *statusTag {
	return &statusTag{}
}

func (s *statusTag) Text(v string) *statusTag {
	s.status = v
	switch s.status {
	case "Alive":
		s.color = "is-primary"
	case "Dead":
		s.color = "is-danger"
	default:
		s.color = "is-warning"
	}
	return s
}

func (s *statusTag) Render() app.UI {
	return app.Span().
		Class("tag").
		Class(s.color).
		Text(s.status)
}

type characterBox struct {
	app.Compo

	CharName     string
	CharSpecies  string
	CharImage    string
	CharStatus   string
	CharLocation string
}

func newCharacterBox() *characterBox {
	return &characterBox{}
}

func (c *characterBox) Name(v string) *characterBox {
	c.CharName = v
	return c
}

func (c *characterBox) Species(v string) *characterBox {
	c.CharSpecies = v
	return c
}

func (c *characterBox) Image(v string) *characterBox {
	c.CharImage = v
	return c
}

func (c *characterBox) Status(v string) *characterBox {
	c.CharStatus = v
	return c
}

func (c *characterBox) Location(v string) *characterBox {
	c.CharLocation = v
	return c
}

func (c *characterBox) Render() app.UI {
	return app.Div().Class("box").Body(

		app.Article().Class("media").Body(
			app.Div().Class("media-left").Body(
				app.Figure().Class("image is-128x128").Body(
					app.Img().Class("is-rounded").Src(c.CharImage),
				),
			),

			app.Div().Class("media-content").Body(
				app.Div().Class("content").Body(
					app.P().Body(
						app.Strong().Text(c.CharName),
						app.Br(),
						app.Small().Text(c.CharSpecies),
						app.Br(),
						newStatusTag().Text(c.CharStatus),
						app.Br(),
						app.Br(),
						app.Small().Class("has-text-grey-light").Text("Last known location: "),
						app.Br(),
						app.Text(c.CharLocation),
					),
				),
			),
		),
	)
}
