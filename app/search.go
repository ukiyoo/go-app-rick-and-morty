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

type Search struct {
	app.Compo

	loader          bool
	PageCount       int
	PageId          int
	Response        CharacterList
	CurrentCategory Category
}

func (s *Search) getCharacterList(url string) {
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
	time.AfterFunc(1*time.Second, s.loaderOff)
	// c.loaderOff()
	s.updateResponse(data)
}

func (s *Search) updateResponse(data CharacterList) {
	s.Response = data
	s.PageCount = data.Info.Pages
	s.Update()
}

func (s *Search) onChangeCharacterList(ctx app.Context, e app.Event) {
	s.loaderOn()

	query := ctx.JSSrc.Get("value").String()
	s.PageId, _ = strconv.Atoi(query)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/?name=%v", query)
	fmt.Println(url)
	s.getCharacterList(url)
}

func (s *Search) onPageClick(ctx app.Context, e app.Event) {
	s.loaderOn()

	pageInt := ctx.JSSrc.Get("text").String()
	s.PageId, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/%v/?page=%v", s.CurrentCategory.Slug, s.PageId)

	app.Dispatch(func() {
		s.getCharacterList(url)
	})
}

func (c *Search) loaderOff() {
	c.loader = true
	c.Update()
}

func (c *Search) loaderOn() {
	c.loader = false
	c.Update()
}

func (s *Search) OnMount(ctx app.Context) {
	s.getCharacterList("https://rickandmortyapi.com/api/character/")
}

func (s *Search) Render() app.UI {
	pages := make([]int, s.PageCount)
	cat := api.GetCategory(s.CurrentCategory.Slug)
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.Div().Class("section").Body(
				app.Div().Class("field has-addons").Body(
					app.Div().Class("control is-grouped is-8").Body(
						app.Input().Class("input is-focused is-primary").Type("text").Placeholder("Find a characters").OnChange(s.onChangeCharacterList),
					),
					app.Div().Class("control").Body(
						app.A().Href("#").Class("button").Text("Search").OnSubmit(s.onChangeCharacterList),
					),
				),
			),

			app.If(!s.loader,
				newLoader(),
			).Else(
				app.Div().Class("columns is-multiline").Body(
					app.Range(s.Response.Results).Slice(func(i int) app.UI {
						character := s.Response.Results[i]
						return app.Div().Class("column is-3").Body(
							app.A().Href("/" + s.CurrentCategory.Slug + "/" + strconv.Itoa(character.ID)).Body(
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
							app.If(s.PageId == 0 && i == 1,
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(s.onPageClick),
							).ElseIf(i == s.PageId,
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(s.onPageClick),
							).Else(
								app.A().Class("pagination-link").Href("/" + cat.Slug).Text(i).OnClick(s.onPageClick),
							),
						)
					}),
				),
			),
		),
	)
}
