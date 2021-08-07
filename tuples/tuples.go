package tuples

import "log"

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
