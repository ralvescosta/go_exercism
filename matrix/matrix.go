package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	var matrix Matrix
	for _, r := range strings.Split(s, "\n") {
		var row []int
		r = strings.TrimSpace(r)
		for _, c := range strings.Split(r, " ") {
			val, err := strconv.Atoi(c)
			if err != nil {
				return nil, fmt.Errorf("convert string to int: %v", err)
			}
			row = append(row, val)
		}
		matrix = append(matrix, row)
	}

	if err := matrix.validate(); err != nil {
		return nil, fmt.Errorf("matrix is not valid: %v", err)
	}

	return matrix, nil
}

func (m Matrix) validate() error {
	if len(m) == 0 {
		return fmt.Errorf("empty input")
	}
	sLen := len(m[0])
	for _, row := range m {
		if len(row) != sLen {
			return fmt.Errorf("rows of different length are not allowed")
		}
	}
	return nil
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, row := range m {
		rows[i] = make([]int, len(row))
		copy(rows[i], row)
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := range cols {
		cols[i] = make([]int, len(m))
	}
	for i, row := range m {
		for j, cell := range row {
			cols[j][i] = cell
		}
	}
	return cols
}

func (m Matrix) Set(row, column, value int) bool {
	if row >= len(m) || row < 0 {
		return false
	}
	if column >= len(m[0]) || column < 0 {
		return false
	}
	m[row][column] = value
	return true
}
