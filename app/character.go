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

type Character struct {
	app.Compo

	loader          bool
	PageCount       int
	PageId          int
	Response        CharacterList
	CurrentCategory Category
}

func (c *Character) getCharacterList(url string) {
	r, err := http.Get(url)
	if err != nil {
		app.Log(err.Error())
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Log(err.Error())
		return
	}

	var data CharacterList

	err = json.Unmarshal(b, &data)

	if err != nil {
		app.Log(err.Error())
		return
	}
	time.AfterFunc(1*time.Second, c.loaderOff)
	// c.loaderOff()
	c.updateResponse(data)
}

func (c *Character) updateResponse(data CharacterList) {
	c.Response = data
	c.PageCount = data.Info.Pages
	c.Update()
}

func (c *Character) OnMount(ctx app.Context) {
	c.getCharacterList(c.CurrentCategory.URL)
}

func (c *Character) onPageClick(ctx app.Context, e app.Event) {
	c.loaderOn()

	pageInt := ctx.JSSrc.Get("text").String()
	c.PageId, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/%v/?page=%v", c.CurrentCategory.Slug, c.PageId)

	app.Dispatch(func() {
		c.getCharacterList(url)
	})
}

func (c *Character) loaderOff() {
	c.loader = true
	c.Update()
}

func (c *Character) loaderOn() {
	c.loader = false
	c.Update()
}

func (c *Character) Render() app.UI {
	cat := api.GetCategory(c.CurrentCategory.Slug)
	pages := make([]int, c.PageCount)
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.H1().Class("title").Text(c.CurrentCategory.Name),
			app.If(!c.loader,
				newLoader(),
			).Else(
				app.Div().Class("columns is-multiline").Body(
					app.Range(c.Response.Results).Slice(func(i int) app.UI {
						character := c.Response.Results[i]
						return app.Div().Class("column is-3").Body(
							app.A().Href("/" + c.CurrentCategory.Slug + "/" + strconv.Itoa(character.ID)).Body(
								newCharacterCard().
									Name(character.Name).
									Image(character.Image).
									Species(character.Species).
									Status(character.Status).
									Location(character.Location.Name),
							),
						)
					}),
				),
			),
			app.Nav().Class("pagination is-centered").Body(
				app.Ul().Class("pagination-list").Body(
					app.Range(pages).Slice(func(i int) app.UI {
						i++
						return app.Li().Body(
							app.If(c.PageId == 0 && i == 1,
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(c.onPageClick),
							).ElseIf(i == c.PageId,
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(c.onPageClick),
							).Else(
								app.A().Class("pagination-link").Href("/" + cat.Slug).Text(i).OnClick(c.onPageClick),
							),
						)
					}),
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

type characterCard struct {
	app.Compo

	character CharacterDetail
}

func newCharacterCard() *characterCard {
	return &characterCard{}
}

func (c *characterCard) Name(v string) *characterCard {
	c.character.Name = v
	return c
}

func (c *characterCard) Species(v string) *characterCard {
	c.character.Species = v
	return c
}

func (c *characterCard) Image(v string) *characterCard {
	c.character.Image = v
	return c
}

func (c *characterCard) Status(v string) *characterCard {
	c.character.Status = v
	return c
}

func (c *characterCard) Location(v string) *characterCard {
	c.character.Location.Name = v
	return c
}

func (c *characterCard) Render() app.UI {
	return app.Div().Class("card").Body(

		app.Div().Class("card-image").Body(
			app.Figure().Class("image").Body(
				app.Img().Src(c.character.Image),
			),
		),

		app.Div().Class("card-content").Body(
			app.Article().Class("media").Body(
				app.Div().Class("content").Body(
					app.P().Class("title is-4").Text(c.character.Name),
					app.P().Class("subtitle is-6").Body(
						app.Text(c.character.Species+" - "),
						newStatusTag().Text(c.character.Status),
					),
					app.P().Class("is-8 has-text-grey-light").Text("Last known location: "),
					app.P().Class("subtitle is-6").Text(c.character.Location.Name),
				),
			),
		),
	)
}
