// Package matrice is a package containing functions to work with matrices
package matrice

import (
	"fmt"
)

type Row []int

type Matrice struct {
	R    int
	C    int
	Rows []Row
}

// Add is a function that add a new row to the matrix ,it  return an error on failure

func (m *Matrice) AddRow(row Row) error {
	if len(row) != m.C {

		return &CustomError{
			Code:    400,
			Message: fmt.Sprintf("you intend to add a row with %d columns to a matrice that requires %d columns", len(row), m.C),
		}
	}
	m.Rows = append(m.Rows, row)
	// Update the number of rows in the matrice
	m.R++
	return nil

}

// Add is a function that return the base function plus the parameter function
func (m *Matrice) Add(m2 *Matrice) (Matrice, error) {
	if m.C != m2.C || m.R != m2.R {
		return Matrice{}, &CustomError{
			Code:    400,
			Message: "you cant add matrices of different lenght",
		}
	}
	rows := make([]Row, m.R)
	for i := 0; i < m.R; i++ {
		row := make([]int, m.C)

		for j := 0; j < m.C; j++ {
			row[j] = m.Rows[i][j] + m2.Rows[i][j]
		}
		rows[i] = row

	}
	return Matrice{
		Rows: rows,
		C:    m.C,
		R:    m.R,
	}, nil
}

// Add is a function that return the base function plus the parameter function
func (m *Matrice) Subtract(m2 *Matrice) (Matrice, error) {
	if m.C != m2.C || m.R != m2.R {
		return Matrice{}, &CustomError{
			Code:    400,
			Message: "you cant subtract matrices of different lenght",
		}
	}
	rows := make([]Row, m.R)
	for i := 0; i < m.R; i++ {
		row := make([]int, m.C)

		for j := 0; j < m.C; j++ {
			row[j] = m.Rows[i][j] - m2.Rows[i][j]
		}
		rows[i] = row

	}
	return Matrice{
		Rows: rows,
		C:    m.C,
		R:    m.R,
	}, nil
}
func (m *Matrice) Multiply(m2 *Matrice) (Matrice, error) {
	if m.C != m2.R {
		return Matrice{}, &CustomError{
			Code:    400,
			Message: "to multiply a matrix, m1.C must be equal to m2.R",
		}
	}
	// Outcome of a matrice will have the order m1.R x m2.C

	rows := make([]Row, m.R)
	for i := 0; i < m.R; i++ {
		row := make([]int, m2.C)
		for j := 0; j < m2.C; j++ {
			row[j] = 0
			// Now lets compute the value of the current row column

			// So we need another loop
			for k := 0; k < m.C; k++ {
				// Sum of all Aik * Bkj
				row[j] += m.Rows[i][k] * m2.Rows[k][j]
			}
		}
		rows[i] = row
	}
	// I can use NewMatrix here to make sure the column of the output tallies with what we get in the rows computation
	newMatrice, err := NewMatrice(rows, m.R, m2.C)
	if err != nil {
		return Matrice{}, err
	}
	return newMatrice, nil
	// return Matrice{
	// 	R: m.R,
	// 	C: m2.C,
	// }, nil
}

// NewMatrice is a constructor function that initializes a Matrice and checks the specifications ,it  return an error on failure
func NewMatrice(rows []Row, r int, c int) (Matrice, error) {
	// Check if the number of rows matches R
	if len(rows) != r {
		return Matrice{}, &CustomError{
			Code:    400,
			Message: fmt.Sprintf("your rows should be a total of %d rows", r),
		}
		// panic(fmt.Sprintf("your rows should be a total of %d rows", r))
	}

	// Check if each row has a length equal to C
	for i, row := range rows {
		if len(row) != c {
			return Matrice{}, &CustomError{
				Code:    400,
				Message: fmt.Sprintf("row %d doesn't have a length of %d", i, c),
			}
			// panic(fmt.Sprintf("row %d doesn't have a length of %d", i, c))
		}
	}

	return Matrice{
		R:    r,
		C:    c,
		Rows: rows,
	}, nil
}
