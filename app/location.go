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

type Location struct {
	app.Compo

	loader          bool
	PageCount       int
	PageId          int
	Response        LocationList
	CurrentCategory Category
}

func (c *Location) getLocationList(url string) {
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

	var data LocationList

	err = json.Unmarshal(b, &data)

	if err != nil {
		app.Log(err.Error())
		return
	}
	time.AfterFunc(1*time.Second, c.loaderOff)
	//c.loaderOff()
	c.updateResponse(data)
}

func (c *Location) updateResponse(data LocationList) {
	c.Response = data
	c.PageCount = data.Info.Pages
	c.Update()
}

func (c *Location) OnMount(ctx app.Context) {
	c.getLocationList(c.CurrentCategory.URL)
}

func (c *Location) onPage(ctx app.Context, e app.Event) {
	c.loaderOn()

	pageInt := ctx.JSSrc.Get("text").String()
	c.PageId, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/%v/?page=%v", c.CurrentCategory.Slug, c.PageId)

	app.Dispatch(func() {
		c.getLocationList(url)
	})
}

func (c *Location) loaderOff() {
	c.loader = true
	c.Update()
}

func (c *Location) loaderOn() {
	c.loader = false
	c.Update()
}

func (c *Location) Render() app.UI {
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
						location := c.Response.Results[i]
						return app.Div().Class("column is-6").Body(
							app.A().Href("/").Body(
								newLocationBox().
									Name(location.Name).
									Dimension(location.Dimension),
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
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(c.onPage),
							).ElseIf(i == c.PageId,
								app.A().Class("pagination-link is-current").Href("/"+cat.Slug).Text(i).OnClick(c.onPage),
							).Else(
								app.A().Class("pagination-link").Href("/" + cat.Slug).Text(i).OnClick(c.onPage),
							),
						)
					}),
				),
			),
		),
	)
}

type loader struct {
	app.Compo
}

func newLoader() *loader {
	return &loader{}
}

func (l *loader) Render() app.UI {
	return app.Section().Class("hero is-fullheight has-text-centered").Style("min-height", "80vh").Body(
		app.Div().Class("hero-body").Body(
			app.Div().Class("container").Body(
				app.Div().Class("lds-dual-ring"),
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
