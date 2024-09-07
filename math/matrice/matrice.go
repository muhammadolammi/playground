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

// Strassen's Matrix Multiplication for 2x2 matrices

// Strassen's Matrix Multiplication for 2x2 matrices
func strassenMultiply(A, B *Matrice) Matrice {
	if A.C != 2 || A.R != 2 || B.C != 2 || B.R != 2 {
		panic("This only works on 2x2 matrices")
	}

	a, b := A.Rows[0][0], A.Rows[0][1]
	c, d := A.Rows[1][0], A.Rows[1][1]
	e, f := B.Rows[0][0], B.Rows[0][1]
	g, h := B.Rows[1][0], B.Rows[1][1]

	p1 := a * (f - h)
	p2 := (a + b) * h
	p3 := (c + d) * e
	p4 := d * (g - e)
	p5 := (a + d) * (e + h)
	p6 := (b - d) * (g + h)
	p7 := (a - c) * (e + f)

	c11 := p5 + p4 - p2 + p6
	c12 := p1 + p2
	c21 := p3 + p4
	c22 := p1 + p5 - p3 - p7

	rows := []Row{
		{c11, c12},
		{c21, c22},
	}

	return Matrice{
		C:    2,
		R:    2,
		Rows: rows,
	}
}

// Function to split a matrix into quadrants
func splitMatrix(A *Matrice) (Matrice, Matrice, Matrice, Matrice) {
	mid := A.R / 2
	A11 := Matrice{R: mid, C: mid, Rows: make([]Row, mid)}
	A12 := Matrice{R: mid, C: mid, Rows: make([]Row, mid)}
	A21 := Matrice{R: mid, C: mid, Rows: make([]Row, mid)}
	A22 := Matrice{R: mid, C: mid, Rows: make([]Row, mid)}

	for i := 0; i < mid; i++ {
		A11.Rows[i] = A.Rows[i][:mid]
		A12.Rows[i] = A.Rows[i][mid:]
		A21.Rows[i] = A.Rows[mid+i][:mid]
		A22.Rows[i] = A.Rows[mid+i][mid:]
	}

	return A11, A12, A21, A22
}

// Strassen's Matrix Multiplication for larger matrices
func (A *Matrice) SMM(B *Matrice) (Matrice, error) {
	if A.R != A.C || B.R != B.C || A.R != B.R {
		return Matrice{}, &CustomError{
			Code:    400,
			Message: "Only works for square matrices of the same size",
		}
	}

	n := A.R
	if n == 2 {
		return strassenMultiply(A, B), nil
	}

	A11, A12, A21, A22 := splitMatrix(A)
	B11, B12, B21, B22 := splitMatrix(B)

	M1, err := A11.Add(&A22)
	if err != nil {
		return Matrice{}, err
	}
	t, _ := B11.Add(&B22)
	M1, err = M1.SMM(&t)
	if err != nil {
		return Matrice{}, err
	}

	M2, err := A21.Add(&A22)
	if err != nil {
		return Matrice{}, err
	}
	M2, err = M2.SMM(&B11)
	if err != nil {
		return Matrice{}, err
	}

	M3, err := B12.Subtract(&B22)
	if err != nil {
		return Matrice{}, err
	}
	M3, err = A11.SMM(&M3)
	if err != nil {
		return Matrice{}, err
	}

	M4, err := B21.Subtract(&B11)
	if err != nil {
		return Matrice{}, err
	}
	M4, err = A22.SMM(&M4)
	if err != nil {
		return Matrice{}, err
	}

	M5, err := A11.Add(&A12)
	if err != nil {
		return Matrice{}, err
	}
	M5, err = M5.SMM(&B22)
	if err != nil {
		return Matrice{}, err
	}

	M6, err := A21.Subtract(&A11)
	if err != nil {
		return Matrice{}, err
	}
	j, _ := B11.Add(&B12)
	M6, err = M6.SMM(&j)
	if err != nil {
		return Matrice{}, err
	}

	M7, err := A12.Subtract(&A22)
	if err != nil {
		return Matrice{}, err
	}
	k, _ := B21.Add(&B22)
	M7, err = M7.SMM(&k)
	if err != nil {
		return Matrice{}, err
	}

	C11, err := M1.Add(&M4)
	if err != nil {
		return Matrice{}, err
	}
	C11, err = C11.Subtract(&M5)
	if err != nil {
		return Matrice{}, err
	}
	C11, err = C11.Add(&M7)
	if err != nil {
		return Matrice{}, err
	}

	C12, err := M3.Add(&M5)
	if err != nil {
		return Matrice{}, err
	}

	C21, err := M2.Add(&M4)
	if err != nil {
		return Matrice{}, err
	}

	C22, err := M1.Add(&M3)
	if err != nil {
		return Matrice{}, err
	}
	C22, err = C22.Subtract(&M2)
	if err != nil {
		return Matrice{}, err
	}
	C22, err = C22.Add(&M6)
	if err != nil {
		return Matrice{}, err
	}

	result := Matrice{R: n, C: n, Rows: make([]Row, n)}
	for i := 0; i < n/2; i++ {
		result.Rows[i] = append(C11.Rows[i], C12.Rows[i]...)
	}
	for i := 0; i < n/2; i++ {
		result.Rows[n/2+i] = append(C21.Rows[i], C22.Rows[i]...)
	}

	return result, nil
}
