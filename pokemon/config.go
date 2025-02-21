package pokemon

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"sort"

	"github.com/glup3/smeargle/images"
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

func (c *PokemonConfig) GetImage(slug, form string, shiny bool) (image.Image, error) {
	fileName := slug
	if form != "" {
		fileName += fmt.Sprintf("-%s", form)
	}

	folder := "regular"
	if shiny {
		folder = "shiny"
	}

	f, err := images.PokemonImages.Open(fmt.Sprintf("%s/%s.png", folder, fileName))
	if err != nil {
		return nil, fmt.Errorf("%s %s: name and form combination does not exist", slug, form)
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return im, nil
}
