package tuples

type Tuple struct {
	X, Y, Z, W float64
}

func Point(x, y, z float64) Tuple {
	// point when w is 1
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	// vector when w is 0
	return Tuple{x, y, z, 0}
}

func Add(a, b int) (c int) {
	return a + b
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
