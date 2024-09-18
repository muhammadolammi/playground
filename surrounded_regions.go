package main

type UnionFind struct {
	Array []int
	Size  []int
}

func NewUnionFind(n int) UnionFind {
	// SPACE AND INITIALIZATION
	// O(n)
	arr := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		size[i] = 1
	}
	return UnionFind{
		Array: arr,
		Size:  size,
	}
}

func Root(array []int, i int) int {
	for i != array[i] {
		// path compression
		array[i] = array[array[i]]
		i = array[i]

	}
	return i

}

func (u UnionFind) Connected(p, q int) bool {
	// O(n+n) = O(n)

	return Root(u.Array, p) == Root(u.Array, q)
}

func (u UnionFind) Max() int {
	max := -1
	// highest := 0
	for i := 0; i < len(u.Size); i++ {
		if max < u.Size[i] {
			max = u.Size[i]

		}
	}
	return max
}

func (u UnionFind) Union(p, q int) {
	// here just one value change
	// this can take O(N) + the two roots computation
	rootP := Root(u.Array, p)
	rootQ := Root(u.Array, q)
	if rootP == rootQ {
		return // They are already in the same set
	}
	// get the size of each root the append the lower to the higher
	if u.Size[rootP] > u.Size[rootQ] {
		u.Array[rootQ] = rootP
		u.Size[rootP] += u.Size[rootQ]
	} else {
		u.Array[rootP] = rootQ
		u.Size[rootQ] += u.Size[rootP]

	}

}

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

func index(i, j, n int) int {
	return i*n + j
}
