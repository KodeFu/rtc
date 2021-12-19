package features

import (
	"fmt"
	"os"

	"sample.com/rtc/utils"
)

type Canvas struct {
	Pixels [][]Color
}

func NewCanvas(width, height int) *Canvas {
	var c Canvas

	// # of rows is the height, think y axis
	c.Pixels = make([][]Color, height)
	for rowIndex := range c.Pixels {
		// # of columns is the width, think x axis
		c.Pixels[rowIndex] = make([]Color, width)
	}

	return &c
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	// since we are height x width, each pixel is at y'th row, x'th column
	c.Pixels[y][x] = color
}

func (c Canvas) PixelAt(x, y int) Color {
	// since we are height x width, each pixel is at y'th row, x'th column
	return c.Pixels[y][x]
}

func (c Canvas) CanvasToPPM() {
	f, _ := os.Create("out.ppm")

	// magic number
	f.WriteString("P3\n")

	// width height
	s := fmt.Sprintf("%d %d\n", c.Width(), c.Height())
	f.WriteString(s)

	// color depth
	f.WriteString("255\n")

	// pixel data
	var charCount int
	for _, row := range c.Pixels {
		for _, color := range row {
			// get the colors and create a list
			red := utils.ScaleColor(color.red)
			green := utils.ScaleColor(color.green)
			blue := utils.ScaleColor(color.blue)
			rgb := []int{red, green, blue}

			// write color components
			for _, component := range rgb {
				s = fmt.Sprintf("%d", component)
				if charCount+len(s)+1 < 70 {
					if charCount > 0 {
						// add space if not first component in a line
						f.WriteString(" ")
					}

					// write component
					f.WriteString(s)

					// length of string plus space
					charCount += len(s) + 1
				} else {
					// write a newline
					f.WriteString("\n")

					// write the component
					f.WriteString(s)

					// length of string
					charCount = len(s)
				}
			}
		}
		// reset character count
		charCount = 0

		// newline for a new row
		f.WriteString("\n")
	}

	f.Close()
}

func (a Canvas) Height() int {
	// rows
	return len(a.Pixels)
}

func (a Canvas) Width() int {
	// columns
	return len(a.Pixels[0])
}
