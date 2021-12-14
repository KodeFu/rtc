package utils

import (
	"math"
)

const EPSILON float64 = 0.0001

func Equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func ScaleColor(a float64) int {
	color := a * 255
	color = math.Round(color)
	if color <= 0 {
		return 0
	}
	if color >= 255 {
		return 255
	}

	return int(color)
}
