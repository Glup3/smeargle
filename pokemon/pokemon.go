package pokemon

import (
	"fmt"
	"image"
	"strings"
)

const blockRune = '▀'

type Pokemon struct {
	Name  string
	Image image.Image
}

func NewPokemon(name string, image image.Image) Pokemon {
	return Pokemon{
		Name:  name,
		Image: image,
	}
}

func (p Pokemon) String() string {
	m := p.Image

	var sb strings.Builder
	bounds := m.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			r2, g2, b2, _ := m.At(x, y+1).RGBA()

			foreground := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
			background := fmt.Sprintf("\033[48;2;%d;%d;%dm", r2, g2, b2)
			sb.WriteString(fmt.Sprintf("%s%s%c\033[0m", foreground, background, blockRune))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}
