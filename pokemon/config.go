package pokemon

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"sort"
	"strconv"

	"github.com/glup3/smeargle/images"
)

const defaultForm = "$"

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
		Forms map[string]struct {
			IsAliasOf string `json:"is_alias_of"`
		} `json:"forms"`
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
		if form == defaultForm {
			continue
		}

		forms = append(forms, form)
	}
	sort.Strings(forms)

	return forms
}

func (c *PokemonConfig) FindImage(slug, form string, shiny bool) (image.Image, error) {
	fileName := slug
	if form != "" {
		alias := c.pokemons[slug].Gen8.Forms[form].IsAliasOf

		if alias == "" {
			fileName += fmt.Sprintf("-%s", form)
		} else if alias != defaultForm {
			fileName += fmt.Sprintf("-%s", alias)
		}
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

func (c *PokemonConfig) RandomPokemon(shinyOdds float32, generations []int) (Pokemon, error) {
	var slugs []string
	generationIds := map[int][2]int{
		1: {1, 151},
		2: {152, 251},
		3: {252, 386},
		4: {387, 493},
		5: {494, 649},
		6: {650, 721},
		7: {722, 809},
		8: {810, 905},
	}

	for slug, pokemon := range c.pokemons {
		if len(generations) == 0 {
			slugs = append(slugs, slug)
			continue
		}

		idx, err := strconv.Atoi(pokemon.Idx)
		if err != nil {
			return Pokemon{}, err
		}

		for _, genId := range generations {
			if idx >= generationIds[genId][0] && idx <= generationIds[genId][1] {
				slugs = append(slugs, slug)
				continue
			}
		}
	}

	if len(slugs) == 0 {
		return Pokemon{}, errors.New("given arguments results in empty list to choose from")
	}

	x := rand.Intn(len(slugs))
	slug := slugs[x]

	forms := c.GetForms(slug)
	forms = append(forms, "")
	x = rand.Intn(len(forms))
	form := forms[x]

	shiny := false
	if rand.Float32() <= shinyOdds {
		shiny = true
	}

	im, err := c.FindImage(slug, form, shiny)
	if err != nil {
		return Pokemon{}, err
	}

	return NewPokemon(slug, im), nil
}
