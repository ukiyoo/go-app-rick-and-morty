package main

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type menu struct {
	app.Compo

	CurrentCategory category
}

func (m *menu) Render() app.UI {
	cats := cat.Categories()
	fmt.Println(cats)
	return app.Div().Class("section").Body(
		app.Aside().Class("menu").Body(
			app.P().Class("menu-label").Text("Rick&Morty"),
			app.Ul().Class("menu-list").Body(
				app.Range(cats).Slice(func(i int) app.UI {
					c := cats[i]
					return newMenuItem().
							Url("/"+c.Slug).
							Text(c.Name).
							IsActive(c.Slug == m.CurrentCategory.Slug)
				}),
			),
		),
	)
}

type menuItem struct {
	app.Compo

	url string
	text     string
	isActive string
}

func newMenuItem() *menuItem {
	return &menuItem{}
}

func (m *menuItem) Text(v string) *menuItem {
	m.text = v
	return m
}

func (m *menuItem) IsActive(v bool) *menuItem {
	if v {
		m.isActive = "is-active"
	}
	return m
}

func (m *menuItem) Url(v string) *menuItem {
	m.url = v
	return m
}


func (m *menuItem) Render() app.UI {
	return app.Li().Body(
		app.A().
			Class(m.isActive).
			Href(m.url).
			Text(m.text),
	)
}
