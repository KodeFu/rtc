package features

import (
	"testing"
)

func TestMatricesCreate4x4(testing *testing.T) {
	var m = NewMatrix(4, 4)

	m.Elements[0][0] = 1
	m.Elements[0][3] = 4
	m.Elements[1][0] = 5.5
	m.Elements[1][2] = 7.5
	m.Elements[2][2] = 11
	m.Elements[3][0] = 13.5
	m.Elements[3][2] = 15.5

	if m.Elements[0][0] != 1 {
		testing.Errorf("unexpected result %v", m.Elements[0][0])
	}
	if m.Elements[0][3] != 4 {
		testing.Errorf("unexpected result %v", m.Elements[0][3])
	}
	if m.Elements[1][0] != 5.5 {
		testing.Errorf("unexpected result %v", m.Elements[1][0])
	}
	if m.Elements[1][2] != 7.5 {
		testing.Errorf("unexpected result %v", m.Elements[1][2])
	}
	if m.Elements[2][2] != 11 {
		testing.Errorf("unexpected result %v", m.Elements[2][2])
	}
	if m.Elements[3][0] != 13.5 {
		testing.Errorf("unexpected result %v", m.Elements[3][0])
	}
	if m.Elements[3][2] != 15.5 {
		testing.Errorf("unexpected result %v", m.Elements[3][2])
	}
}

func TestMatricesCreate2x2(testing *testing.T) {
	var m = NewMatrix(2, 2)

	m.Elements[0][0] = -3
	m.Elements[0][1] = 5
	m.Elements[1][0] = 1
	m.Elements[1][1] = -2

	if m.Elements[0][0] != -3 {
		testing.Errorf("unexpected result %v", m.Elements[0][0])
	}
	if m.Elements[0][1] != 5 {
		testing.Errorf("unexpected result %v", m.Elements[0][1])
	}
	if m.Elements[1][0] != 1 {
		testing.Errorf("unexpected result %v", m.Elements[1][0])
	}
	if m.Elements[1][1] != -2 {
		testing.Errorf("unexpected result %v", m.Elements[1][1])
	}

}

func TestMatricesCreate3x3(testing *testing.T) {
	var m = NewMatrix(3, 3)

	m.Elements[0][0] = -3
	m.Elements[1][1] = -2
	m.Elements[2][2] = 1

	if m.Elements[0][0] != -3 {
		testing.Errorf("unexpected result %v", m.Elements[0][0])
	}
	if m.Elements[1][1] != -2 {
		testing.Errorf("unexpected result %v", m.Elements[1][1])
	}
	if m.Elements[2][2] != 1 {
		testing.Errorf("unexpected result %v", m.Elements[2][2])
	}

}

func TestMatricesEqualityIdentical(testing *testing.T) {
	var m1 = NewMatrix(4, 4)
	var m2 = NewMatrix(4, 4)

	m1.Elements = [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	m2.Elements = [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}

	if !m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}

func TestMatricesEqualityIdenticalEpsilon(testing *testing.T) {
	var m1 = NewMatrix(4, 4)
	var m2 = NewMatrix(4, 4)

	m1.Elements = [][]float64{{1.0005, 2, 3, 4}, {5, 6, 7, 8.000000000}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	m2.Elements = [][]float64{{1.0006, 2, 3, 4}, {5, 6, 7, 8.000100000}, {9, 8, 7, 6}, {5, 4, 3, 2}}

	if !m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}

func TestMatricesEqualityIdenticalEpsilonDelta(testing *testing.T) {
	var m1 = NewMatrix(4, 4)
	var m2 = NewMatrix(4, 4)

	m1.Elements = [][]float64{{1.0005, 2, 3, 4}, {5, 6, 7, 8.000000000}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	m2.Elements = [][]float64{{1.0006, 2, 3, 4}, {5, 6, 7, 8.000200000}, {9, 8, 7, 6}, {5, 4, 3, 2}}

	if m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}

func TestMatricesEqualityDifferent(testing *testing.T) {
	var m1 = NewMatrix(4, 4)
	var m2 = NewMatrix(4, 4)

	m1.Elements = [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	m2.Elements = [][]float64{{2, 3, 4, 5}, {6, 7, 8, 9}, {8, 7, 6, 5}, {4, 3, 2, 1}}

	if m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}
