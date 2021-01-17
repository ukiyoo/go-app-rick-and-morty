package main

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type All struct {
	characterList CharacterList
	locationList  LocationList
}

type Content struct {
	app.Compo

	loader          bool
	PageCount       int
	PageId          int
	Response        All
	CurrentCategory Category
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

	switch c.CurrentCategory.Slug {
	case CHARACTER:
		err = json.Unmarshal(b, &data.characterList)
		c.PageCount = data.characterList.Info.Pages
	case LOCATION:
		err = json.Unmarshal(b, &data.locationList)
		c.PageCount = data.locationList.Info.Pages
	}

	if err != nil {
		app.Log(err.Error())
		return
	}
	time.AfterFunc(1*time.Second, c.loaderOff)
	//c.loaderOff()
	c.updateResponse(data)
}

func (c *Content) updateResponse(data All) {
	c.Response = data
	c.Update()
}

func (c *Content) OnMount(ctx app.Context) {
	c.getApi(c.CurrentCategory.URL)
}

func (c *Content) onPage(ctx app.Context, e app.Event) {
	c.loaderOn()

	pageInt := ctx.JSSrc.Get("text").String()
	c.PageId, _ = strconv.Atoi(pageInt)
	url := fmt.Sprintf("https://rickandmortyapi.com/api/%v/?page=%v", c.CurrentCategory.Slug, c.PageId)

	app.Dispatch(func() {
		c.getApi(url)
	})
}

func (c *Content) loaderOff() {
	c.loader = true
	c.Update()
}

func (c *Content) loaderOn() {
	c.loader = false
	c.Update()
}

func (c *Content) Render() app.UI {
	cat := api.GetCategory(c.CurrentCategory.Slug)
	pages := make([]int, c.PageCount)
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.H1().Class("title").Text(c.CurrentCategory.Name),
			app.If(!c.loader,
				newLoader(),
			).Else(
				app.Div().Class("columns is-multiline").Body(
					app.If(c.CurrentCategory.Slug == CHARACTER,
						app.Range(c.Response.characterList.Results).Slice(func(i int) app.UI {
							character := c.Response.characterList.Results[i]
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
					).ElseIf(c.CurrentCategory.Slug == LOCATION,
						app.Range(c.Response.locationList.Results).Slice(func(i int) app.UI {
							location := c.Response.locationList.Results[i]
							return app.Div().Class("column is-6").Body(
								app.A().Href("/").Body(
									newLocationBox().
										Name(location.Name).
										Dimension(location.Dimension),
								),
							)
						}),
					).ElseIf(c.CurrentCategory.Slug == EPISODE,
						&Episode{},
					),
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
						app.Text(c.character.Location.Name),
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

//type episodeBox struct {
//	app.Compo
//
//	episode EpisodeDetail
//}
//
//func newEpisodeBox() *episodeBox {
//	return &episodeBox{}
//}
//
//func (e *episodeBox) Name(v string) *episodeBox {
//	e.episode.Name = v
//	return e
//}
//
//func (e *episodeBox) AirDate(v string) *episodeBox {
//	e.episode.AirDate = v
//	return e
//}
//
//func (e *episodeBox) Episode(v string) *episodeBox {
//	e.episode.Episode = v
//	return e
//}
//
//func (e *episodeBox) Render() app.UI {
//	return app.Div().Class("box").Body(
//
//		app.Article().Class("media").Body(
//			app.Div().Class("media-content").Body(
//				app.Div().Class("content").Body(
//					app.P().Body(
//						app.Strong().Text(e.episode.Name),
//						app.Br(),
//						app.Strong().Text(e.episode.AirDate),
//						app.Br(),
//						app.Small().Class("has-text-grey-light").Text("Episode: "),
//						app.Br(),
//						app.Text(e.episode.Episode),
//					),
//				),
//			),
//		),
//	)
//}

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

type episodeTabs struct {
	app.Compo

	episode EpisodeDetail
}

func newEpisodeTabs() *episodeTabs {
	return &episodeTabs{}
}

func (e *episodeTabs) Name(v string) *episodeTabs {
	e.episode.Name = v
	return e
}

func (e *episodeTabs) AirDate(v string) *episodeTabs {
	e.episode.AirDate = v
	return e
}

func (e *episodeTabs) Episode(v string) *episodeTabs {
	e.episode.Episode = v
	return e
}

func (e *episodeTabs) Render() app.UI {
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
