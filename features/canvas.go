package features

import (
	"fmt"
	"os"
)

type Canvas struct {
	Width, Height int
	Pixels        [][]Color
}

func NewCanvas(width, height int) *Canvas {
	c := Canvas{Width: width, Height: height}

	// we make() len set to width so we can address via
	// Pixels[x][y], otherwise we'd address Pixels[y][x]
	c.Pixels = make([][]Color, width)
	for i := range c.Pixels {
		c.Pixels[i] = make([]Color, height)
		fmt.Printf("Row %d: %v\n", i, c.Pixels[i])
	}

	return &c
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	c.Pixels[x][y] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	return c.Pixels[x][y]
}

func (c *Canvas) CanvasToPPM() {
	f, _ := os.Create("out.ppm")
	f.WriteString("P3\n")
	f.WriteString("255\n")
	s := fmt.Sprintf("%d %d\n", c.Width, c.Height)
	f.WriteString(s)

	for index, color := range c.Pixels {
		fmt.Println("index", index, "color", color)
	}

	f.Close()
}
