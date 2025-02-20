package pokemon

import (
	"fmt"
	"image"
	"strings"
)

const (
	upperBlock = '▀'
	lowerBlock = '▄'
	emptyBlock = ' '
)

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

	for y := minY; y <= maxY; y += 2 {
		for x := minX; x <= maxX; x++ {
			r, g, b, a := p.Image.At(x, y).RGBA()
			r2, g2, b2, a2 := p.Image.At(x, y+1).RGBA()

			_, _ = a, a2

			if a == 0 && a2 == 0 {
				sb.WriteRune(emptyBlock)
				continue
			}

			if a == 0 {
				sb.WriteString(fmt.Sprintf("%s%c", fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), lowerBlock))
				continue
			}

			if a2 == 0 {
				sb.WriteString(fmt.Sprintf("%s%c", fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), upperBlock))
				continue
			}

			foreground := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
			background := fmt.Sprintf("\033[48;2;%d;%d;%dm", r2, g2, b2)
			sb.WriteString(fmt.Sprintf("%s%s%c\033[0m", foreground, background, upperBlock))
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
			r, g, b, a := m.At(x, y).RGBA()

			// Include non-transparent and black pixels in bounds
			if a > 0 || (r == 0 && g == 0 && b == 0 && a == 65535) {
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
