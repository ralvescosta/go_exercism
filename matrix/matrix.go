package matrix

import (
	"strconv"
	"strings"
)

type Matrix struct {
	in string
}

func (mtx Matrix) Rows() [][]int {
	rows := strings.Split(mtx.in, "\n")
	rowCount := 1
	matrix := [][]int{}

	for _, row := range rows {
		if len(matrix) < rowCount {
			matrix = append(matrix, []int{})
		}
		columns := strings.Split(row, " ")
		for _, column := range columns {
			cv, _ := strconv.Atoi(column)
			matrix[rowCount-1] = append(matrix[rowCount-1], cv)
		}
		rowCount++
	}
	return matrix
}
func (Matrix) Cols() [][]int {
	return make([][]int, 1)
}
func (Matrix) Set(r, c, v int) bool {
	return true
}

func New(input string) (*Matrix, error) {
	return &Matrix{in: input}, nil
}
