package features

import (
	"testing"

	"sample.com/rtc/utils"
)

func TestColors(testing *testing.T) {
	var c = NewColor(-0.5, 0.4, 1.7)

	if c.red != -0.5 || c.green != 0.4 || c.blue != 1.7 {
		testing.Errorf("unexpected result %v", c)
	}
}

func TestColorsAdd(testing *testing.T) {
	var a = NewColor(0.9, 0.6, 0.75)
	var b = NewColor(0.7, 0.1, 0.25)

	r := a.Add(b)

	if r.red != 1.6 || r.green != 0.7 || r.blue != 1.0 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestColorsSub(testing *testing.T) {
	var a = NewColor(0.9, 0.6, 0.75)
	var b = NewColor(0.7, 0.1, 0.25)

	r := a.Sub(b)

	if !utils.Equal(r.red, 0.2) || r.green != 0.5 || r.blue != 0.5 {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestColorsMultScalar(testing *testing.T) {
	var a = NewColor(0.2, 0.3, 0.4)
	var b = 2.0

	r := a.Mult(b)

	if r.red != 0.4 || r.green != 0.6 || !utils.Equal(r.blue, 0.8) {
		testing.Errorf("unexpected result %v", r)
	}
}

func TestColorsMult(testing *testing.T) {
	var a = NewColor(1.0, 0.2, 0.4)
	var b = NewColor(0.9, 1.0, 0.1)

	r := a.HadamardProduct(b)

	if r.red != 0.9 || r.green != 0.2 || !utils.Equal(r.blue, 0.04) {
		testing.Errorf("unexpected result %v", r)
	}
}
