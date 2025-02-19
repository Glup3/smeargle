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
	var sb strings.Builder

	minX, minY, maxX, maxY := findVisibleBounds(p.Image)

	for y := minY; y < maxY; y += 2 {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := p.Image.At(x, y).RGBA()
			r2, g2, b2, _ := p.Image.At(x, y+1).RGBA()

			foreground := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
			background := fmt.Sprintf("\033[48;2;%d;%d;%dm", r2, g2, b2)
			sb.WriteString(fmt.Sprintf("%s%s%c\033[0m", foreground, background, blockRune))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func findVisibleBounds(m image.Image) (minX, minY, maxX, maxY int) {
	bounds := m.Bounds()
	minX, minY = bounds.Max.X, bounds.Max.Y
	maxX, maxY = bounds.Min.X, bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := m.At(x, y).RGBA()
			if a > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Handle cases where the whole image is transparent
	if minX > maxX || minY > maxY {
		return bounds.Min.X, bounds.Min.Y, bounds.Min.X, bounds.Min.Y
	}

	return minX, minY, maxX, maxY
}
