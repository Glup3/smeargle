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
	resetCode  = "\033[0m"
)

type Pokemon struct {
	name  string
	image image.Image
}

type RGBAOverride struct {
	From [4]uint32
	To   [4]uint32
}

func NewPokemon(name string, image image.Image) Pokemon {
	return Pokemon{
		name:  name,
		image: image,
	}
}

func (p Pokemon) String(colorOverrides []RGBAOverride) string {
	var sb strings.Builder

	minX, minY, maxX, maxY := findVisibleBounds(p.image)

	for y := minY; y <= maxY; y += 2 {
		for x := minX; x <= maxX; x++ {
			r, g, b, a := p.image.At(x, y).RGBA()
			r2, g2, b2, a2 := p.image.At(x, y+1).RGBA()

			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			r2, g2, b2, a2 = r2>>8, g2>>8, b2>>8, a2>>8

			for _, o := range colorOverrides {
				if r == o.From[0] && g == o.From[1] && b == o.From[2] && a == o.From[3] {
					r, g, b, a = o.To[0], o.To[1], o.To[2], o.To[3]
				}

				if r2 == o.From[0] && g2 == o.From[1] && b2 == o.From[2] && a2 == o.From[3] {
					r2, g2, b2, a2 = o.To[0], o.To[1], o.To[2], o.To[3]
				}
			}

			foreground := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)

			if a == 0 && a2 == 0 {
				sb.WriteRune(emptyBlock)
				continue
			}

			if a == 0 {
				sb.WriteString(fmt.Sprintf("%s%c%s", foreground, lowerBlock, resetCode))
				continue
			}

			if a2 == 0 {
				sb.WriteString(fmt.Sprintf("%s%c%s", foreground, upperBlock, resetCode))
				continue
			}

			background := fmt.Sprintf("\033[48;2;%d;%d;%dm", r2, g2, b2)
			sb.WriteString(fmt.Sprintf("%s%s%c%s", foreground, background, upperBlock, resetCode))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (p Pokemon) ColorHistogram() map[string]int {
	minX, minY, maxX, maxY := findVisibleBounds(p.image)

	histogram := make(map[string]int)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			r, g, b, a := p.image.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8

			rgbaKey := fmt.Sprintf("%d %d %d %d", r, g, b, a)
			histogram[rgbaKey]++
		}
	}

	return histogram
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

func ParseRGBAOverride(override string) (RGBAOverride, error) {
	var from, to [4]uint32
	_, err := fmt.Sscanf(override, "%d %d %d %d=%d %d %d %d", &from[0], &from[1], &from[2], &from[3], &to[0], &to[1], &to[2], &to[3])
	if err != nil {
		return RGBAOverride{}, fmt.Errorf("invalid override format: %v", err)
	}
	return RGBAOverride{From: from, To: to}, nil
}
