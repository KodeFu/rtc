package features

import "sample.com/rtc/utils"

type Matrix struct {
	Elements [][]float64
}

func NewMatrix(rows, columns int) Matrix {
	var m Matrix

	// # of rows
	m.Elements = make([][]float64, rows)
	for i := range m.Elements {
		// # of columns
		m.Elements[i] = make([]float64, columns)
	}

	return m
}

func NewMatrix4x4(elements [4][4]float64) Matrix {
	m := NewMatrix(4, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.Elements[i][j] = elements[i][j]
		}
	}

	return m
}

func NewMatrix3x3(elements [3][3]float64) Matrix {
	m := NewMatrix(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m.Elements[i][j] = elements[i][j]
		}
	}

	return m
}

func NewMatrix2x2(elements [2][2]float64) Matrix {
	m := NewMatrix(2, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			m.Elements[i][j] = elements[i][j]
		}
	}

	return m
}

func Identity4x4() Matrix {
	m := NewMatrix4x4([4][4]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}})
	return m
}

func (a Matrix) Equals(b Matrix) bool {
	if a.NumRows() != b.NumRows() || a.NumCols() != b.NumCols() {
		return false
	}

	for i := 0; i < a.NumRows(); i++ {
		for j := 0; j < a.NumCols(); j++ {
			if !utils.Equal(a.Elements[i][j], b.Elements[i][j]) {
				return false
			}
		}
	}

	return true
}

func (a Matrix) NumRows() int {
	return len(a.Elements)
}

func (a Matrix) NumCols() int {
	return len(a.Elements[0])
}

func (a Matrix) Multiply(b Matrix) Matrix {
	var total = NewMatrix(a.NumRows(), a.NumCols())
	for i, row := range a.Elements {
		for j := range row {
			for k := 0; k < a.NumCols(); k++ {
				total.Elements[i][j] += a.Elements[i][k] * b.Elements[k][j]
			}
		}
	}

	return total
}

func (a Matrix) MultiplyByTuple(b Tuple) Tuple {
	var total Tuple
	for i, row := range a.Elements {
		for j := range row {
			total.Elements[i] += a.Elements[i][j] * b.Elements[j]
		}
	}

	return total
}

func (a Matrix) Transpose() Matrix {
	var t = NewMatrix(a.NumRows(), a.NumCols())
	for i, row := range a.Elements {
		for j := range row {
			t.Elements[j][i] = a.Elements[i][j]
		}
	}

	return t
}

func (a Matrix) Determinant() float64 {
	det := 0.0

	if a.NumRows() == 2 && a.NumCols() == 2 {
		det = a.Elements[0][0]*a.Elements[1][1] - a.Elements[0][1]*a.Elements[1][0]
	} else {
		for column := 0; column < a.NumCols(); column++ {
			det += a.Elements[0][column] * a.Cofactor(0, column)
		}
	}

	return det
}

func (a Matrix) Submatrix(row, col int) Matrix {
	// remove row phase
	var rowPhase = NewMatrix(a.NumRows()-1, a.NumCols())
	rowPhaseCount := 0
	for i, rowElement := range a.Elements {
		for j := range rowElement {
			if i != row {
				rowPhase.Elements[rowPhaseCount][j] = a.Elements[i][j]
			}
		}
		if i != row {
			rowPhaseCount++
		}
	}

	// remove column phase
	var colPhase = NewMatrix(a.NumRows()-1, a.NumCols()-1)
	for i, rowElement := range rowPhase.Elements {
		colPhaseCount := 0
		for j := range rowElement {
			if j != col {
				colPhase.Elements[i][colPhaseCount] = rowPhase.Elements[i][j]
				colPhaseCount++
			}
		}
	}

	return colPhase
}

func (a Matrix) Minor(row, col int) float64 {
	return a.Submatrix(row, col).Determinant()
}

func (a Matrix) Cofactor(row, col int) float64 {
	minor := a.Minor(row, col)

	if (row+col)%2 != 0 {
		return -minor
	}

	return minor
}
