package matrix

import (
	"strconv"
	"strings"
)

type Matrix struct {
	in string
	m  [][]int
}

func (Matrix) Rows() [][]int {
	return make([][]int, 1)
}
func (Matrix) Cols() [][]int {
	return make([][]int, 1)
}
func (Matrix) Set(r, c, v int) bool {
	return true
}

func New(input string) (*Matrix, error) {
	rows := strings.Split(input, "\n")
	rowCount := 0
	matrix := [][]int{}

	for _, row := range rows {
		columns := strings.Split(row, " ")

		for columnIndex, column := range columns {
			cv, _ := strconv.Atoi(column)
			matrix[rowCount][columnIndex] = cv
		}
		rowCount++
	}

	return &Matrix{in: input, m: matrix}, nil
}
