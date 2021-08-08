package tuples

import (
	"log"
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

func IsPoint(t Tuple) bool {
	if t.W == 1 {
		return true
	}

	return false
}

func IsVector(t Tuple) bool {
	if t.W == 0 {
		return true
	}

	return false
}

func (t Tuple) Add(a Tuple) Tuple {
	if IsPoint(t) && IsPoint(a) {
		log.Panic("adding two points")
	}

	var r Tuple
	r.X = t.X + a.X
	r.Y = t.Y + a.Y
	r.Z = t.Z + a.Z
	r.W = t.W + a.W

	return r
}

func (t Tuple) Sub(a Tuple) Tuple {
	var r Tuple
	r.X = t.X - a.X
	r.Y = t.Y - a.Y
	r.Z = t.Z - a.Z
	r.W = t.W - a.W

	return r
}

func (t Tuple) Negate() Tuple {
	var r Tuple
	r.X = -t.X
	r.Y = -t.Y
	r.Z = -t.Z
	r.W = -t.W

	return r
}

func (t Tuple) Mult(a float64) Tuple {
	var r Tuple
	r.X = a * t.X
	r.Y = a * t.Y
	r.Z = a * t.Z
	r.W = a * t.W

	return r
}

func (t Tuple) Div(a float64) Tuple {
	var r Tuple
	r.X = t.X / a
	r.Y = t.Y / a
	r.Z = t.Z / a
	r.W = t.W / a

	return r
}

func (t Tuple) Magnitude() float64 {
	var r float64
	r = math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))

	return r
}

func (t Tuple) Normalize() Tuple {
	var r Tuple
	var mag = t.Magnitude()
	r = Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}

	return r
}
