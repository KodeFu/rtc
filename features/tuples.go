package features

import (
	"math"

	"sample.com/rtc/utils"
)

type Tuple struct {
	Elements [4]float64
}

func NewTuple(x, y, z, w float64) Tuple {
	var t Tuple
	t.Elements[0] = x
	t.Elements[1] = y
	t.Elements[2] = z
	t.Elements[3] = w
	return t
}

func NewPoint(x, y, z float64) Tuple {
	var t Tuple
	t.Elements[0] = x
	t.Elements[1] = y
	t.Elements[2] = z
	t.Elements[3] = 1
	return t
}

func NewVector(x, y, z float64) Tuple {
	var t Tuple
	t.Elements[0] = x
	t.Elements[1] = y
	t.Elements[2] = z
	t.Elements[3] = 0
	return t
}

func IsPoint(a Tuple) bool {
	if a.Elements[3] == 1 {
		return true
	}

	return false
}

func IsVector(a Tuple) bool {
	if a.Elements[3] == 0 {
		return true
	}

	return false
}

func (a Tuple) Add(b Tuple) Tuple {
	var r Tuple
	r.Elements[0] = a.Elements[0] + b.Elements[0]
	r.Elements[1] = a.Elements[1] + b.Elements[1]
	r.Elements[2] = a.Elements[2] + b.Elements[2]
	r.Elements[3] = a.Elements[3] + b.Elements[3]
	return r
}

func (a Tuple) Sub(b Tuple) Tuple {
	var r Tuple
	r.Elements[0] = a.Elements[0] - b.Elements[0]
	r.Elements[1] = a.Elements[1] - b.Elements[1]
	r.Elements[2] = a.Elements[2] - b.Elements[2]
	r.Elements[3] = a.Elements[3] - b.Elements[3]
	return r
}

func (a Tuple) Negate() Tuple {
	var r Tuple
	r.Elements[0] = -a.Elements[0]
	r.Elements[1] = -a.Elements[1]
	r.Elements[2] = -a.Elements[2]
	r.Elements[3] = -a.Elements[3]
	return r
}

func (a Tuple) Mult(b float64) Tuple {
	var r Tuple
	r.Elements[0] = b * a.Elements[0]
	r.Elements[1] = b * a.Elements[1]
	r.Elements[2] = b * a.Elements[2]
	r.Elements[3] = b * a.Elements[3]
	return r
}

func (a Tuple) Div(b float64) Tuple {
	var r Tuple
	r.Elements[0] = a.Elements[0] / b
	r.Elements[1] = a.Elements[1] / b
	r.Elements[2] = a.Elements[2] / b
	r.Elements[3] = a.Elements[3] / b
	return r
}

func (a Tuple) Magnitude() float64 {
	var r float64
	r = math.Sqrt(math.Pow(a.Elements[0], 2) + math.Pow(a.Elements[1], 2) + math.Pow(a.Elements[2], 2) + math.Pow(a.Elements[3], 2))

	return r
}

func (a Tuple) Normalize() Tuple {
	var r Tuple
	var mag = a.Magnitude()
	r.Elements[0] = a.Elements[0] / mag
	r.Elements[1] = a.Elements[1] / mag
	r.Elements[2] = a.Elements[2] / mag
	r.Elements[3] = a.Elements[3] / mag

	return r
}

func (a Tuple) Dot(b Tuple) float64 {
	var r float64
	r = a.Elements[0]*b.Elements[0] + a.Elements[1]*b.Elements[1] + a.Elements[2]*b.Elements[2] + a.Elements[3]*b.Elements[3]

	return r
}

func (a Tuple) Cross(b Tuple) Tuple {
	var r Tuple

	r.Elements[0] = a.Elements[1]*b.Elements[2] - a.Elements[2]*b.Elements[1]
	r.Elements[1] = a.Elements[2]*b.Elements[0] - a.Elements[0]*b.Elements[2]
	r.Elements[2] = a.Elements[0]*b.Elements[1] - a.Elements[1]*b.Elements[0]
	r.Elements[3] = 0

	return r
}

func (a Tuple) Equals(b Tuple) bool {
	for i := 0; i < 4; i++ {
		if !utils.Equal(a.Elements[i], b.Elements[i]) {
			return false
		}
	}

	return true
}
