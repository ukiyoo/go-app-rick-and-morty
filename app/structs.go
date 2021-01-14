package main

import "time"

type CharacterList struct {
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

type LocationList struct {
	Info struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"info"`
	Results []struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Type      string    `json:"type"`
		Dimension string    `json:"dimension"`
		Residents []string  `json:"residents"`
		URL       string    `json:"url"`
		Created   time.Time `json:"created"`
	} `json:"results"`
}

type EpisodeList struct {
	Info struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"info"`
	Results []struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		AirDate    string    `json:"air_date"`
		Episode    string    `json:"episode"`
		Characters []string  `json:"characters"`
		URL        string    `json:"url"`
		Created    time.Time `json:"created"`
	} `json:"results"`
}

type CharacterDetail struct {
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
}


type LocationDetail struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	Residents []string  `json:"residents"`
	URL       string    `json:"url"`
	Created   time.Time `json:"created"`
}

type EpisodeDetail struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	AirDate    string    `json:"air_date"`
	Episode    string    `json:"episode"`
	Characters []string  `json:"characters"`
	URL        string    `json:"url"`
	Created    time.Time `json:"created"`
}