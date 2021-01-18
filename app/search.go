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
	query string
	error           bool
}

func (s *Search) getCharacterList(url string) {
	r, err := http.Get(url)
	if err != nil {
		app.Log(err.Error())
		return
	}
	if r.StatusCode != 200 {
		s.error = true
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

	s.query = ctx.JSSrc.Get("value").String()
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character/?name=%v", s.query)
	s.getCharacterList(url)
}

func (s *Search) onPageClick(ctx app.Context, e app.Event) {
	s.loaderOn()

	pageInt := ctx.JSSrc.Get("text").String()
	s.PageId, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/%v/?page=%v&name=%v", CHARACTER, s.PageId, s.query)

	app.Dispatch(func() {
		s.getCharacterList(url)
	})
}

func (s *Search) loaderOff() {
	s.loader = true
	s.Update()
}

func (s *Search) loaderOn() {
	s.loader = false
	s.Update()
}

func (s *Search) OnMount(ctx app.Context) {
	s.getCharacterList("https://rickandmortyapi.com/api/character/")
}

func (s *Search) Render() app.UI {
	pages := make([]int, s.PageCount)
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.Div().Class("columns").Body(
				app.Div().Class("column is-6").Body(
					app.Div().Class("field is-grouped").Body(
						app.Div().Class("control is-expanded").Body(
							app.Input().Class("input").Type("text").Placeholder("Find a characters").OnChange(s.onChangeCharacterList),
						),
						app.Div().Class("control").Body(
							app.A().Href("#").Class("button").Text("Find").OnSubmit(s.onChangeCharacterList),
						),
					),
				),
			),
			app.If(!s.loader,
				newLoader(),
			).ElseIf(s.error,
				app.Div().Class("has-text-centered").Body(
					app.P().Class("is-size-2").Text("There is nothing here"),
				),
			).Else(
				app.Div().Class("columns is-multiline").Body(
					app.Range(s.Response.Results).Slice(func(i int) app.UI {
						character := s.Response.Results[i]
						return app.Div().Class("column is-3").Body(
							app.A().Href("#").Body(
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
								app.A().Class("pagination-link is-current").Href("#").Text(i).OnClick(s.onPageClick),
							).ElseIf(i == s.PageId,
								app.A().Class("pagination-link is-current").Href("#").Text(i).OnClick(s.onPageClick),
							).Else(
								app.A().Class("pagination-link").Href("#").Text(i).OnClick(s.onPageClick),
							),
						)
					}),
				),
			),
		),
	)
}
