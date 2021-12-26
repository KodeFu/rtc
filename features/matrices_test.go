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
	m := NewMatrix2x2([2][2]float64{{-3, 5}, {1, -2}})

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
	m := NewMatrix3x3([3][3]float64{{-3, 5, 0}, {1, -2, -7}, {0, 1, 1}})

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
	m1 := NewMatrix4x4([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	m2 := NewMatrix4x4([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})

	if !m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}

func TestMatricesEqualityIdenticalEpsilon(testing *testing.T) {
	m1 := NewMatrix4x4([4][4]float64{{1.0005, 2, 3, 4}, {5, 6, 7, 8.000000000}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	m2 := NewMatrix4x4([4][4]float64{{1.0006, 2, 3, 4}, {5, 6, 7, 8.000100000}, {9, 8, 7, 6}, {5, 4, 3, 2}})

	if !m1.Equals(m2) {
		testing.Errorf("unexpected result %v %v", m1, m2)
	}
}

func TestMatricesEqualityIdenticalEpsilonDelta(testing *testing.T) {
	var m1 = NewMatrix4x4([4][4]float64{{1.0005, 2, 3, 4}, {5, 6, 7, 8.000000000}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	var m2 = NewMatrix4x4([4][4]float64{{1.0006, 2, 3, 4}, {5, 6, 7, 8.000200000}, {9, 8, 7, 6}, {5, 4, 3, 2}})

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

func TestMatricesMultiply(testing *testing.T) {
	var m1 = NewMatrix4x4([4][4]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	var m2 = NewMatrix4x4([4][4]float64{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}})
	result := m1.Multiply(m2)

	var expected = NewMatrix4x4([4][4]float64{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}})

	if !expected.Equals(result) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesMultiplyByTuple(testing *testing.T) {
	var m = NewMatrix4x4([4][4]float64{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}})
	var b = NewTuple(1, 2, 3, 1)
	result := m.MultiplyByTuple(b)

	var expected = NewTuple(18, 24, 33, 1)

	if !expected.Equals(result) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesMultiplyByIdentity(testing *testing.T) {
	var m = NewMatrix4x4([4][4]float64{{1, 2, 3, 4}, {1, 2, 4, 8}, {2, 4, 8, 16}, {4, 8, 16, 32}})
	result := m.Multiply(Identity4x4())

	if !m.Equals(result) {
		testing.Errorf("unexpected result %v %v", m, result)
	}
}

func TestMatricesIdentityMutiplyByTuple(testing *testing.T) {
	var m = Identity4x4()
	var a = NewTuple(1, 2, 3, 1)
	result := m.MultiplyByTuple(a)

	if !a.Equals(result) {
		testing.Errorf("unexpected result %v %v", a, result)
	}
}

func TestMatricesTranspose4x4(testing *testing.T) {
	var m = NewMatrix4x4([4][4]float64{{0, 9, 3, 0}, {9, 8, 0, 8}, {1, 8, 5, 3}, {0, 0, 5, 8}})
	result := m.Transpose()

	var expected = NewMatrix4x4([4][4]float64{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}})

	if !expected.Equals(result) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesTranspose3x3(testing *testing.T) {
	var m = NewMatrix3x3([3][3]float64{{0, 9, 3}, {9, 8, 0}, {1, 8, 5}})
	result := m.Transpose()

	var expected = NewMatrix3x3([3][3]float64{{0, 9, 1}, {9, 8, 8}, {3, 0, 5}})

	if !expected.Equals(result) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesTransposeIdentity(testing *testing.T) {
	var m = Identity4x4()
	result := m.Transpose()

	var expected = NewMatrix4x4([4][4]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}})

	if !expected.Equals(result) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}
