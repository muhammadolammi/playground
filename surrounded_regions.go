package main

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	uf := NewUnionFind((m * n) + 1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				// lets check if its first row or last row or first collumn or last column
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					// we join this to the last index in the uf
					uf.Union(index(i, j, n), m*n)
					continue
				}
				// Now lets work on 4 connection possibilities
				// if we reach here , its definately not a border O, so we check if they are connected to another O

				if board[i-1][j] == 'O' {
					// i-1 means the upper row, then j means the current column
					// we join them together if they are connected
					uf.Union(index(i, j, n), index(i-1, j, n))
				}
				if board[i+1][j] == 'O' {
					// i-1 means the lower row, then j means the current column
					// we join them together if they are connected
					uf.Union(index(i, j, n), index(i+1, j, n))
				}
				if board[i][j+1] == 'O' {
					// i means the current row, then j+1 means the next column
					// we join them together if they are connected
					uf.Union(index(i, j, n), index(i, j+1, n))
				}
				if board[i][j-1] == 'O' {
					// i means the current row, then j-1 means the previous column
					// we join them together if they are connected
					uf.Union(index(i, j, n), index(i, j-1, n))
				}

			}

		}
	}

	// let's make another m*n loop
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				// lets check if its not connected to the virtual node
				if !uf.Connected(index(i, j, n), m*n) {
					board[i][j] = 'X'
				}
			}
		}
	}

}
