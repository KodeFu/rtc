package features

import (
	"testing"

	"sample.com/rtc/utils"
)

func TestColors(testing *testing.T) {
	var c = Color{red: -0.5, green: 0.4, blue: 1.7}

	if c.red != -0.5 || c.green != 0.4 || c.blue != 1.7 {
		testing.Errorf("unexpected result %v", c)
	}
}

func TestColorsAdd(testing *testing.T) {
	var a = Color{red: 0.9, green: 0.6, blue: 0.75}
	var b = Color{red: 0.7, green: 0.1, blue: 0.25}

	r := a.Add(b)

	if r.red != 1.6 || r.green != 0.7 || r.blue != 1.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestColorsSub(testing *testing.T) {
	var a = Color{red: 0.9, green: 0.6, blue: 0.75}
	var b = Color{red: 0.7, green: 0.1, blue: 0.25}

	r := a.Sub(b)

	if !utils.Equal(r.red, 0.2) || r.green != 0.5 || r.blue != 0.5 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestColorsMult(testing *testing.T) {
	var a = Color{red: 1.0, green: 0.2, blue: 0.4}
	var b = Color{red: 0.9, green: 1.0, blue: 0.1}

	r := a.Mult(b)

	if r.red != 0.9 || r.green != 0.2 || !utils.Equal(r.blue, 0.04) {
		testing.Errorf("unexpected result %v", r)
	}
}
