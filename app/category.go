package main

var (
	cat = newAPIStore()
)

type APIStore struct {
	categories []category
}

type category struct {
	Name string
	Slug string
	URL  string
}

func newAPIStore() *APIStore {
	c := []category{
		{
			Name: "CHARACTERS",
			Slug: "characters",
			URL:  "https://rickandmortyapi.com/api/character",
		},
		{
			Name: "LOCATION",
			Slug: "locations",
			URL:  "https://rickandmortyapi.com/api/location",
		},
		{
			Name: "EPISODES",
			Slug: "episodes",
			URL:  "https://rickandmortyapi.com/api/episode",
		},
	}
	return &APIStore{categories: c}
}

func (s *APIStore) Get(slug string) category {
	var c category
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

func (s *APIStore) Categories() []category {
	c := make([]category, len(s.categories))
	copy(c, s.categories)
	return c
}
