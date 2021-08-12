package features

import (
	"math"
)

type Tuple struct {
	X, Y, Z, W float64
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func IsPoint(a Tuple) bool {
	if a.W == 1 {
		return true
	}

	return false
}

func IsVector(a Tuple) bool {
	if a.W == 0 {
		return true
	}

	return false
}

func (a Tuple) Add(b Tuple) Tuple {
	var r Tuple
	r.X = a.X + b.X
	r.Y = a.Y + b.Y
	r.Z = a.Z + b.Z
	r.W = a.W + b.W

	return r
}

func (a Tuple) Sub(b Tuple) Tuple {
	var r Tuple
	r.X = a.X - b.X
	r.Y = a.Y - b.Y
	r.Z = a.Z - b.Z
	r.W = a.W - b.W

	return r
}

func (a Tuple) Negate() Tuple {
	var r Tuple
	r.X = -a.X
	r.Y = -a.Y
	r.Z = -a.Z
	r.W = -a.W

	return r
}

func (a Tuple) Mult(b float64) Tuple {
	var r Tuple
	r.X = b * a.X
	r.Y = b * a.Y
	r.Z = b * a.Z
	r.W = b * a.W

	return r
}

func (a Tuple) Div(b float64) Tuple {
	var r Tuple
	r.X = a.X / b
	r.Y = a.Y / b
	r.Z = a.Z / b
	r.W = a.W / b

	return r
}

func (a Tuple) Magnitude() float64 {
	var r float64
	r = math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.W, 2))

	return r
}

func (a Tuple) Normalize() Tuple {
	var r Tuple
	var mag = a.Magnitude()
	r = Tuple{a.X / mag, a.Y / mag, a.Z / mag, a.W / mag}

	return r
}

func (a Tuple) Dot(b Tuple) float64 {
	var r float64
	r = a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W

	return r
}

func (a Tuple) Cross(b Tuple) Tuple {
	var r Tuple

	r.X = a.Y*b.Z - a.Z*b.Y
	r.Y = a.Z*b.X - a.X*b.Z
	r.Z = a.X*b.Y - a.Y*b.X
	r.W = 0

	return r
}
