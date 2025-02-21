package pokemon

import (
	_ "embed"
	"encoding/json"
	"sort"
)

type dataConfig struct {
	Idx  string     `json:"idx"`
	Name nameConfig `json:"name"`
	Slug slugConfig `json:"slug"`
}

type nameConfig struct {
	Eng   string `json:"eng"`
	Chs   string `json:"chs"`
	Jpn   string `json:"jpn"`
	JpnRo string `json:"jpn_ro"`
}

type slugConfig struct {
	Eng   string `json:"eng"`
	Jpn   string `json:"jpn"`
	JpnRo string `json:"jpn_ro"`
}

//go:embed pokemon.json
var pokemonConfig []byte

func LoadSlugs() ([]string, error) {
	var config map[string]dataConfig
	if err := json.Unmarshal(pokemonConfig, &config); err != nil {
		return nil, err
	}

	slugs := make([]string, len(config))

	i := 0
	for _, item := range config {
		slugs[i] = item.Slug.Eng
		i++
	}

	sort.Strings(slugs)

	return slugs, nil
}
