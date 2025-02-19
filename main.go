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
	reader, err := os.Open("pikachu.png")
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

	pokemon := pokemon.NewPokemon("pikachu", m)

	fmt.Println(pokemon.String())
}
