package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) Rows() (res [][]int) {
	res = make([][]int, len(m))
	for i, row := range m {
		res[i] = make([]int, len(row))
		for j, cell := range row {
			res[i][j] = cell
		}
	}
	return res
}

func (m Matrix) Cols() (res [][]int) {
	res = make([][]int, len(m[0]))
	for i := range res {
		res[i] = make([]int, len(m))
	}
	for i, row := range m {
		for j, cell := range row {
			res[j][i] = cell
		}
	}
	return res
}

func (m Matrix) Set(r, c, val int) bool {
	if r >= len(m) || r < 0 || c >= len(m[0]) || c < 0 {
		return false
	}

	m[r][c] = val
	return true
}

func New(in string) (Matrix, error) {
	rows := strings.Split(in, "\n")
	n := len(rows)
	res := make(Matrix, n)
	var m int
	var err error

	for i, row := range rows {
		cells := strings.Split(strings.TrimSpace(row), " ")

		if i == 0 {
			m = len(cells)
		} else if len(cells) != m {
			return nil, errors.New("Row length must be constant.")
		}

		res[i] = make([]int, len(cells))
		for j, cell := range cells {
			if res[i][j], err = strconv.Atoi(cell); err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}
