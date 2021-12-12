package features

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(testing *testing.T) {
	var c = NewCanvas(10, 20)

	if c.Width != 10 {
		testing.Errorf("unexpected width %v", c.Width)
	}

	if c.Height != 20 {
		testing.Errorf("unexpected height %v", c.Height)
	}

	for i := 0; i < c.Width; i++ {
		for j := 0; j < c.Height; j++ {
			if (c.PixelAt(i, j) != Color{0, 0, 0}) {
				testing.Errorf("unexpected result %v", c)
			}
		}
	}
}

func TestWritePixel(testing *testing.T) {
	var c = NewCanvas(10, 20)

	red := Color{1, 0, 0}
	c.WritePixel(2, 3, red)

	if c.PixelAt(2, 3) != red {
		testing.Errorf("unexpected pixel %v", c.PixelAt(2, 3))
	}
}

func TestContructingPPMHeader(testing *testing.T) {
	var c = NewCanvas(5, 3)

	c.CanvasToPPM()

	f, _ := os.Open("out.ppm")
	var r *bufio.Reader = bufio.NewReader(f)

	s1, _ := r.ReadString('\n')
	s2, _ := r.ReadString('\n')
	s3, _ := r.ReadString('\n')

	if s1 != "P3\n" {
		testing.Errorf("unexpected string %v", s1)
	}

	if s2 != "5 3\n" {
		testing.Errorf("unexpected string %v", s3)
	}

	if s3 != "255\n" {
		testing.Errorf("unexpected string %v", s2)
	}

}

func TestCanvasToPPM(testing *testing.T) {
	var c = NewCanvas(5, 3)

	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	c.CanvasToPPM()

	if c.PixelAt(0, 0) != c1 {
		testing.Errorf("unexpected pixel %v", c.PixelAt(0, 0))
	}

	if c.PixelAt(2, 1) != c2 {
		testing.Errorf("unexpected pixel %v", c.PixelAt(2, 1))
	}

	if c.PixelAt(4, 2) != c3 {
		testing.Errorf("unexpected pixel %v", c.PixelAt(4, 2))
	}
}

func TestSplittingLongLines(testing *testing.T) {
	var c = NewCanvas(10, 2)

	color := Color{1.0, 0.8, 0.6}

	for i := 0; i < c.Width; i++ {
		for j := 0; j < c.Height; j++ {
			c.WritePixel(i, j, color)
		}
	}

	c.CanvasToPPM()
}

func TestNewLine(testing *testing.T) {
	var c = NewCanvas(5, 3)

	color := Color{0.0, 0.0, 0.0}

	for i := 0; i < c.Width; i++ {
		for j := 0; j < c.Height; j++ {
			c.WritePixel(i, j, color)
		}
	}

	c.CanvasToPPM()

	// check for newline at the end
	bytes, _ := ioutil.ReadFile("out.ppm")
	if bytes[len(bytes)-1] != '\n' {
		testing.Errorf("no newline at end %d", bytes[len(bytes)-1])
	}
}
