package main

var (
	api       = newAPIStore()
	CHARACTER = "character"
	LOCATION  = "location"
	EPISODE   = "episode"
)

type APIStore struct {
	categories []Category
	seasons    []Season
}

type Season struct {
	SeasonName string
	SeasonURL  string
}

type Category struct {
	Name string
	Slug string
	URL  string
}

func newAPIStore() *APIStore {
	c := []Category{
		{
			Name: "CHARACTERS",
			Slug: CHARACTER,
			URL:  "https://rickandmortyapi.com/api/character",
		},
		{
			Name: "LOCATIONS",
			Slug: LOCATION,
			URL:  "https://rickandmortyapi.com/api/location",
		},
		{
			Name: "EPISODES",
			Slug: EPISODE,
			URL:  "https://rickandmortyapi.com/api/episode",
		},
	}
	s := []Season{
		{
			SeasonName: "S01",
			SeasonURL:  "https://rickandmortyapi.com/api/episode/?episode=S01",
		},
		{
			SeasonName: "S02",
			SeasonURL:  "https://rickandmortyapi.com/api/episode/?episode=S02",
		},
		{
			SeasonName: "S03",
			SeasonURL:  "https://rickandmortyapi.com/api/episode/?episode=S03",
		},
		{
			SeasonName: "S04",
			SeasonURL:  "https://rickandmortyapi.com/api/episode/?episode=S04",
		},
	}

	return &APIStore{categories: c, seasons: s}
}

func (s *APIStore) Seasons() []string {
	seasons := make([]string, len(s.seasons))
	for i, c := range s.seasons {
		seasons[i] = c.SeasonName
	}
	return seasons
}

func (s *APIStore) GetCategory(slug string) Category {
	var c Category
	for _, c := range s.categories {
		if c.Slug == slug {
			return c
		}
	}
	return c
}

func (s *APIStore) Slugs() []string {
	slugs := make([]string, len(s.categories))

	for i, c := range s.categories {
		slugs[i] = c.Slug
	}

	return slugs
}

func (s *APIStore) Categories() []Category {
	c := make([]Category, len(s.categories))
	copy(c, s.categories)
	return c
}
