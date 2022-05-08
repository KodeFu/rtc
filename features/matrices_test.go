package features

import (
	"testing"

	"sample.com/rtc/utils"
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

func TestMatricesDeterminant(testing *testing.T) {
	m := NewMatrix2x2([2][2]float64{{1, 5}, {-3, 2}})
	result := m.Determinant()

	var expected = 17.0

	if !utils.Equal(result, expected) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesSubmatrix3x3(testing *testing.T) {
	m := NewMatrix3x3([3][3]float64{{1, 5, 0}, {-3, 2, 7}, {0, 6, -3}})
	result := m.Submatrix(0, 2)

	var expected = NewMatrix2x2([2][2]float64{{-3, 2}, {0, 6}})

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesSubmatrix4x4(testing *testing.T) {
	m := NewMatrix4x4([4][4]float64{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}})
	result := m.Submatrix(2, 1)

	var expected = NewMatrix3x3([3][3]float64{{-6, 1, 6}, {-8, 8, 6}, {-7, -1, 1}})

	if !result.Equals(expected) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesMinor(testing *testing.T) {
	m := NewMatrix3x3([3][3]float64{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}})
	result := m.Minor(1, 0)

	var expected = 25.0

	if !utils.Equal(result, expected) {
		testing.Errorf("unexpected result %v %v", expected, result)
	}
}

func TestMatricesCofactor1(testing *testing.T) {
	m := NewMatrix3x3([3][3]float64{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}})
	resultMinor := m.Minor(0, 0)
	resultCofactor := m.Cofactor(0, 0)

	var expectedMinor = -12.0
	var expectedCofactor = -12.0

	if !utils.Equal(resultMinor, expectedMinor) {
		testing.Errorf("unexpected result %v %v", expectedMinor, resultMinor)
	}

	if !utils.Equal(resultCofactor, expectedCofactor) {
		testing.Errorf("unexpected result %v %v", expectedCofactor, resultCofactor)
	}
}

func TestMatricesCofactor2(testing *testing.T) {
	m := NewMatrix3x3([3][3]float64{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}})
	resultMinor := m.Minor(1, 0)
	resultCofactor := m.Cofactor(1, 0)

	var expectedMinor = 25.0
	var expectedCofactor = -25.0

	if !utils.Equal(resultMinor, expectedMinor) {
		testing.Errorf("unexpected result %v %v", expectedMinor, resultMinor)
	}

	if !utils.Equal(resultCofactor, expectedCofactor) {
		testing.Errorf("unexpected result %v %v", expectedCofactor, resultCofactor)
	}
}

func TestMatricesDeterminant3x3(testing *testing.T) {
	m := NewMatrix3x3([3][3]float64{{1, 2, 6}, {-5, 8, -4}, {2, 6, 4}})
	resultCofactor00 := m.Cofactor(0, 0)
	resultCofactor01 := m.Cofactor(0, 1)
	resultCofactor02 := m.Cofactor(0, 2)
	resultDeterminant := m.Determinant()

	var expectedCofactor00 = 56.0
	var expectedCofactor01 = 12.0
	var expectedCofactor02 = -46.0
	var expectedDeterminant = -196.0

	if !utils.Equal(resultCofactor00, expectedCofactor00) {
		testing.Errorf("unexpected result %v %v", resultCofactor00, expectedCofactor00)
	}

	if !utils.Equal(resultCofactor01, expectedCofactor01) {
		testing.Errorf("unexpected result %v %v", resultCofactor01, expectedCofactor01)
	}

	if !utils.Equal(resultCofactor02, expectedCofactor02) {
		testing.Errorf("unexpected result %v %v", resultCofactor02, expectedCofactor02)
	}

	if !utils.Equal(resultDeterminant, expectedDeterminant) {
		testing.Errorf("unexpected result %v %v", resultDeterminant, expectedDeterminant)
	}
}

func TestMatricesDeterminant4x4(testing *testing.T) {
	m := NewMatrix4x4([4][4]float64{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}})
	resultCofactor00 := m.Cofactor(0, 0)
	resultCofactor01 := m.Cofactor(0, 1)
	resultCofactor02 := m.Cofactor(0, 2)
	resultCofactor03 := m.Cofactor(0, 3)
	resultDeterminant := m.Determinant()

	var expectedCofactor00 = 690.0
	var expectedCofactor01 = 447.0
	var expectedCofactor02 = 210.0
	var expectedCofactor03 = 51.0
	var expectedDeterminant = -4071.0

	if !utils.Equal(resultCofactor00, expectedCofactor00) {
		testing.Errorf("unexpected result %v %v", resultCofactor00, expectedCofactor00)
	}

	if !utils.Equal(resultCofactor01, expectedCofactor01) {
		testing.Errorf("unexpected result %v %v", resultCofactor01, expectedCofactor01)
	}

	if !utils.Equal(resultCofactor02, expectedCofactor02) {
		testing.Errorf("unexpected result %v %v", resultCofactor02, expectedCofactor02)
	}

	if !utils.Equal(resultCofactor03, expectedCofactor03) {
		testing.Errorf("unexpected result %v %v", resultCofactor03, expectedCofactor03)
	}

	if !utils.Equal(resultDeterminant, expectedDeterminant) {
		testing.Errorf("unexpected result %v %v", resultDeterminant, expectedDeterminant)
	}
}
