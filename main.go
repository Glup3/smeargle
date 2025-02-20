package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log/slog"
	"os"

	"github.com/glup3/farbeagle/pokemon"
)

func main() {
	name := "houndoom"

	p1 := loadPokemon(name, false)
	p2 := loadPokemon(name, true)

	fmt.Println(p1.String())
	fmt.Println(p2.String())
}

func loadPokemon(name string, shiny bool) pokemon.Pokemon {
	var folder string
	if shiny {
		folder = "shiny"
	} else {
		folder = "regular"
	}

	reader, err := os.Open(fmt.Sprintf("images/%s/%s.png", folder, name))
	if err != nil {
		slog.Error("unable to open image", slog.Any("error", err))
		os.Exit(1)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		slog.Error("decoding image failed", slog.Any("error", err))
		os.Exit(1)
	}

	return pokemon.NewPokemon("pikachu", m)
}
