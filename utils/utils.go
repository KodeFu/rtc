package utils

import "math"

const EPSILON = 0.0001

func Equal(a, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	}

	return false
}
