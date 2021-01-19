package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type EpisodeSingle struct {
	app.Compo

	loader          bool
	PageCount       int
	PageId          int
	Response        EpisodeDetail
	Characters      CharacterList
	CurrentCategory Category
}

func (c *EpisodeSingle) getEpisodeDetail(url string) {
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

	var data EpisodeDetail

	err = json.Unmarshal(b, &data)

	if err != nil {
		app.Log(err.Error())
		return
	}
	// time.AfterFunc(1*time.Second, c.loaderOff)
	// c.loaderOff()
	c.updateResponse(data)
}

func (c *EpisodeSingle) getCharacterList(url string) {
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
	// time.AfterFunc(1*time.Second, c.loaderOff)
	// c.loaderOff()
	c.updateCharacters(data)
}

func (e *EpisodeSingle) updateCharacters(data CharacterList) {
	e.Characters = data
	e.Update()
}

func (e *EpisodeSingle) updateResponse(data EpisodeDetail) {
	e.Response = data
	e.Update()
}

func (e *EpisodeSingle) getCharacterFromOneEpisode() {
	re := regexp.MustCompile("[0-9]+")
	characters := []string{}
	for _, u := range e.Response.Characters {
		fmt.Println(u)
		id := re.FindStringSubmatch(u)
		fmt.Println(id)
		characters = append(characters, id[0])
	}
	fmt.Println(characters)
	args := strings.Join(characters, ",")
	e.getCharacterList("https://rickandmortyapi.com/api/character/" + args)
}

func (c *EpisodeSingle) Render() app.UI {
	return app.Div().Class("section").Body(
		app.Div().Class("container").Body(
			app.H1().Class("title").Text(c.CurrentCategory.Name),
			app.If(!c.loader,
				newLoader(),
			).Else(
				app.Div().Body(
					app.P().Text(c.Response.Name),
				),

				app.Div().Class("columns is-multiline").Body(
					app.Range(c.Characters.Results).Slice(func(i int) app.UI {
						character := c.Characters.Results[i]
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
		),
	)
}
