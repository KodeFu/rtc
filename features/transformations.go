package features

func Translation(x, y, z float64) Matrix {
	m := NewMatrix4x4([4][4]float64{{1, 0, 0, x}, {0, 1, 0, y}, {0, 0, 1, z}, {0, 0, 0, 1}})

	return m
}
