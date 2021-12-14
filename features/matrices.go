package features

import "sample.com/rtc/utils"

type Matrix struct {
	Rows, Columns int
	Elements      [][]float64
}

func NewMatrix(rows, columns int) Matrix {
	m := Matrix{Rows: rows, Columns: columns}
	// # of rows
	m.Elements = make([][]float64, rows)
	for i := range m.Elements {
		// # of columns
		m.Elements[i] = make([]float64, columns)
	}

	return m
}

func (a Matrix) Equals(b Matrix) bool {
	if a.Rows != b.Rows || a.Columns != b.Columns {
		return false
	}

	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Columns; j++ {
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

func (a Matrix) NumColumns() int {
	return len(a.Elements[0])
}
