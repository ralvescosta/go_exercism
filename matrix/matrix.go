package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	in string
}

func (mtx Matrix) Rows() [][]int {
	rowsString := strings.Split(mtx.in, "\n")
	rows := [][]int{}

	for rowIndex, row := range rowsString {
		if len(rows) < rowIndex+1 {
			rows = append(rows, []int{})
		}
		columns := strings.Split(row, " ")
		for _, column := range columns {
			cv, _ := strconv.Atoi(column)
			rows[rowIndex] = append(rows[rowIndex], cv)
		}
	}

	return rows
}
func (mtx Matrix) Cols() [][]int {
	rowsString := strings.Split(mtx.in, "\n")
	cols := [][]int{}

	for _, row := range rowsString {
		columns := strings.Split(row, " ")
		for columnIndex, column := range columns {
			if len(cols) < columnIndex+1 {
				cols = append(cols, []int{})
			}
			cv, _ := strconv.Atoi(column)
			cols[columnIndex] = append(cols[columnIndex], cv)
		}
	}

	return cols
}

func (Matrix) toString(mtx [][]int) string {
	mtxString := ""
	for _, row := range mtx {
		for _, colum := range row {
			mtxString += fmt.Sprintf(" %d", colum)
		}
		mtxString = mtxString[1:]
	}

	return mtxString
}

func (mtx *Matrix) Set(row, col, value int) bool {
	rows := mtx.Rows()

	if len(rows) < row-1 {
		return false
	}
	if len(rows) < row {
		rows = append(rows, []int{})
	}

	rows[row][col] = value
	mtx.in = mtx.toString(rows)
	return true
}

func New(input string) (*Matrix, error) {
	return &Matrix{in: input}, nil
}
