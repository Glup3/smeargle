package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log/slog"
	"os"
	"strings"
)

const blockRune = '▀'

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

		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}
