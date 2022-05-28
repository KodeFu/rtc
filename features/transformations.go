package features

import "math"

func Translation(x, y, z float64) Matrix {
	m := NewMatrix4x4([4][4]float64{{1, 0, 0, x}, {0, 1, 0, y}, {0, 0, 1, z}, {0, 0, 0, 1}})

	return m
}

func Scaling(x, y, z float64) Matrix {
	m := NewMatrix4x4([4][4]float64{{x, 0, 0, 0}, {0, y, 0, 0}, {0, 0, z, 0}, {0, 0, 0, 1}})

	return m
}

func Rotation_X(r float64) Matrix {
	m := NewMatrix4x4([4][4]float64{{1, 0, 0, 0}, {0, math.Cos(r), -math.Sin(r), 0}, {0, math.Sin(r), math.Cos(r), 0}, {0, 0, 0, 1}})

	return m
}

func Rotation_Y(r float64) Matrix {
	m := NewMatrix4x4([4][4]float64{
		{math.Cos(r), 0, math.Sin(r), 0},
		{0, 1, 0, 0},
		{-math.Sin(r), 0, math.Cos(r), 0},
		{0, 0, 0, 1}})

	return m
}

func Rotation_Z(r float64) Matrix {
	m := NewMatrix4x4([4][4]float64{
		{math.Cos(r), -math.Sin(r), 0, 0},
		{math.Sin(r), math.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1}})

	return m
}
