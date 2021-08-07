package utils

import (
	"testing"
)

func TestEqual(testing *testing.T) {
	a := 10.00005
	b := 10.00010
	r := Equal(a, b)
	if r != true {
		testing.Errorf("equal, but reported not equal %v", r)
	}

	a = 10.00030
	b = 10.00010
	r = Equal(a, b)
	if r == true {
		testing.Errorf("not equal, but reported equal %v", r)
	}
}
