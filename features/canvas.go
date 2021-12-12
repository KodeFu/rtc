package features

import (
	"fmt"
	"os"

	"sample.com/rtc/utils"
)

type Canvas struct {
	Width, Height int
	Pixel         [][]Color
}

func NewCanvas(width, height int) *Canvas {
	c := Canvas{Width: width, Height: height}
	// # of rows is the height, think y axis
	c.Pixel = make([][]Color, height)
	for row, _ := range c.Pixel {
		// # of columns is the width, think x axis
		c.Pixel[row] = make([]Color, width)
	}

	return &c
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	// since we are height x width, each pixel is at y'th row, x'th column
	c.Pixel[y][x] = color
}

func (c Canvas) PixelAt(x, y int) Color {
	// since we are height x width, each pixel is at y'th row, x'th column
	return c.Pixel[y][x]
}

func (c Canvas) CanvasToPPM() {
	f, _ := os.Create("out.ppm")

	// magic number
	f.WriteString("P3\n")

	// width height
	s := fmt.Sprintf("%d %d\n", c.Width, c.Height)
	f.WriteString(s)
	f.WriteString("255\n")

	// pixel data
	var charCount int
	for _, pixels := range c.Pixel {
		for _, color := range pixels {
			red := utils.ScaleColor(color.red)
			green := utils.ScaleColor(color.green)
			blue := utils.ScaleColor(color.blue)
			//s = fmt.Sprintf("%d %d %d", red, green, blue)
			s = fmt.Sprintf("%d", red)
			// character count + length of new string + space
			if charCount+len(s)+1 < 70 {
				if charCount > 0 {
					f.WriteString(" ")
				}
				f.WriteString(s)
				charCount += len(s) + 1
			} else {
				f.WriteString("\n")
				f.WriteString(s)
				charCount = len(s)
			}

			s = fmt.Sprintf("%d", green)
			// character count + length of new string + space
			if charCount+len(s)+1 < 70 {
				if charCount > 0 {
					f.WriteString(" ")
				}
				f.WriteString(s)
				charCount += len(s) + 1
			} else {
				f.WriteString("\n")
				f.WriteString(s)
				charCount = len(s)
			}

			s = fmt.Sprintf("%d", blue)
			// character count + length of new string + space
			if charCount+len(s)+1 < 70 {
				if charCount > 0 {
					f.WriteString(" ")
				}
				f.WriteString(s)
				charCount += len(s) + 1
			} else {
				f.WriteString("\n")
				f.WriteString(s)
				charCount = len(s)
			}
		}
		charCount = 0
		f.WriteString("\n")
	}

	f.Close()
}
