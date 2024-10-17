package main

import "fmt"

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
	// for i != array[i] {
	// 	// path compression
	// 	array[i] = array[array[i]]
	// 	i = array[i]

	// }
	// return i

	// this approach will help with this function
	root := i
	// Find the root of the element
	for root != array[root] {
		root = array[root]
	}
	// Path compression
	for i != root {
		next := array[i]
		array[i] = root
		i = next
	}
	return root

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

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	uf := NewUnionFind(n * m)

	// Iterate over the grid and apply Union-Find
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// Check neighboring cells
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(index(i, j, n), index(i-1, j, n))
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union(index(i, j, n), index(i+1, j, n))
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(index(i, j, n), index(i, j-1, n))
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(index(i, j, n), index(i, j+1, n))
				}
			}
		}
	}

	// Count distinct roots corresponding to '1's
	connected := map[int]bool{}
	fmt.Println(uf.Array)
	for idx, _ := range uf.Array {
		i, j := reindex(idx, n)
		if grid[i][j] == '1' {
			root := Root(uf.Array, idx) // Get the root for compression.
			connected[root] = true      // Count unique roots (islands).
		}
	}
	return len(connected)
}

func index(i, j, n int) int {
	return i*n + j
}
func reindex(idx, n int) (int, int) {
	i := idx / n // This gives the row (i)
	j := idx % n // This gives the column (j)
	return i, j
}
