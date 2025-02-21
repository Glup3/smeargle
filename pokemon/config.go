package pokemon

import (
	_ "embed"
	"encoding/json"
	"sort"
)

//go:embed pokemon.json
var pokemonJson []byte

type pokemonData struct {
	Idx string `json:"idx"`

	Name struct {
		Eng   string `json:"eng"`
		Chs   string `json:"chs"`
		Jpn   string `json:"jpn"`
		JpnRo string `json:"jpn_ro"`
	} `json:"name"`

	Slug struct {
		Eng   string `json:"eng"`
		Jpn   string `json:"jpn"`
		JpnRo string `json:"jpn_ro"`
	} `json:"slug"`

	Gen8 struct {
		Forms map[string]interface{} `json:"forms"`
	} `json:"gen-8"`
}

type slugEng = string

type PokemonConfig struct {
	pokemons map[slugEng]pokemonData
}

func NewPokemonConfig() (*PokemonConfig, error) {
	var config map[string]pokemonData
	if err := json.Unmarshal(pokemonJson, &config); err != nil {
		return nil, err
	}

	pokemons := make(map[slugEng]pokemonData)
	for _, p := range config {
		pokemons[p.Slug.Eng] = p
	}

	return &PokemonConfig{
		pokemons: pokemons,
	}, nil
}

func (c *PokemonConfig) GetSlugs() []string {
	slugs := make([]string, len(c.pokemons))

	i := 0
	for _, p := range c.pokemons {
		slugs[i] = p.Slug.Eng
		i++
	}
	sort.Strings(slugs)

	return slugs
}

func (c *PokemonConfig) GetForms(slug slugEng) []string {
	var forms []string

	for form := range c.pokemons[slug].Gen8.Forms {
		if form == "$" {
			continue
		}

		forms = append(forms, form)
	}
	sort.Strings(forms)

	return forms
}
