package main

var (
	api = newAPIStore()
	CHARACTER = "character"
	LOCATION = "location"
	EPISODE = "episode"
)

type APIStore struct {
	categories []Category
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
	return &APIStore{categories: c}
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
