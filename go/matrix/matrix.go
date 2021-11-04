package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	// create a matrix instance
	m := make(Matrix, 0)

	// split the string matrix into rows
	mString := strings.Split(s, "\n")

	for i := range mString {
		// clean up the row by trimming leading/trailing spaces
		row := strings.TrimSpace(mString[i])

		// create the columns by splitting the spaces
		columns := strings.Split(row, " ")

		// append a new array for every row
		m = append(m, make([]int, 0))

		// if the column length is not equal to the next then the matrix is not a matrix
		if i != len(mString)-1 {
			if len(columns) != len(strings.Split(strings.TrimSpace(mString[i+1]), " ")) {
				return &Matrix{}, errors.New("inconsistent column length")
			}
		}

		for j := range columns {
			// convert the column string to an integer
			colInt, err := strconv.Atoi(columns[j])

			// if the integer cannot be converted return an error
			if err != nil {
				return &Matrix{}, errors.New("failed to convert integer")
			}

			// if everything is OK append the column value to the matrix.
			m[i] = append(m[i], colInt)
		}
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	mCopy := make([][]int, len((*m)[0]))

	for _, r := range *m {
		for j := range r {
			mCopy[j] = append(mCopy[j], r[j])
		}
	}

	return mCopy
}

func (m *Matrix) Rows() [][]int {
	mCopy := make([][]int, 0)

	for i, r := range *m {
		mCopy = append(mCopy, make([]int, 0))
		mCopy[i] = append(mCopy[i], r...)
	}

	return mCopy
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < len(*m) && row >= 0 && col < len((*m)[row]) && col >= 0 && len(*m) > 0 {
		(*m)[row][col] = val
		return true
	}

	return false
}
