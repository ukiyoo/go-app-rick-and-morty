package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Episode struct {
	app.Compo

	//loader          bool
	//PageCount       int
	//PageId          int
	activeTab       string
	Response        EpisodeList
	CurrentCategory Category
}

func (e *Episode) getEpisodesFromSeasons(url string) {
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

	var data EpisodeList

	err = json.Unmarshal(b, &data)

	if err != nil {
		app.Log(err.Error())
		return
	}
	//time.AfterFunc(1*time.Second, c.loaderOff)
	//e.loaderOff()
	e.updateResponse(data)
}

func (e *Episode) updateResponse(data EpisodeList) {
	e.Response = data
	e.Update()
}

func (e *Episode) onShowEpisodes(ctx app.Context, event app.Event) {
	//c.loaderOn()
	seasonID := ctx.JSSrc.Get("text").String()
	e.activeTab = seasonID
	e.getEpisodesFromSeasons("https://rickandmortyapi.com/api/episode/?episode=" + seasonID)

}

func (e *Episode) OnMount(ctx app.Context) {
	e.getEpisodesFromSeasons("https://rickandmortyapi.com/api/episode/?episode=S01")
	e.activeTab = "S01"

}

func (e *Episode) Render() app.UI {
	seasons := api.Seasons()
	return app.Div().Class("column").Body(
		app.Div().Class("tabs is-toggle").Body(
			app.Ul().Body(
				app.Range(seasons).Slice(func(i int) app.UI {
					return app.If(seasons[i] == e.activeTab,
						app.Li().Class("is-active").Style("z-index", "0").Body(
							app.A().Href("#").OnClick(e.onShowEpisodes).Body(
								app.Span().Text(seasons[i]),
							),
						),
					).Else(
						app.Li().Class("").Body(
							app.A().Href("#").OnClick(e.onShowEpisodes).Body(
								app.Span().Text(seasons[i]),
							),
						),
					)
				}),
			),
		),
		app.Div().Class("columns is-multiline").Body(
			app.Range(e.Response.Results).Slice(func(i int) app.UI {
				episode := e.Response.Results[i]
				return app.Div().Class("column is-6").Body(
					app.A().Href("#").Body(
						app.Div().Class("box").Body(
							app.Article().Class("media").Body(
								app.Div().Class("media-content").Body(
									app.Div().Class("content").Body(
										app.P().Body(
											app.Strong().Text(episode.Name),
											app.Br(),
											app.Strong().Text(episode.AirDate),
											app.Br(),
											app.Small().Class("has-text-grey-light").Text("Episode: "),
											app.Br(),
											app.Text(episode.Episode),
										),
									),
								),
							),
						),
					),
				)
			}),
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
