package main

// type UnionFind struct {
// 	Array []int
// 	Size  []int
// }

// func NewUnionFind(n int) UnionFind {
// 	// SPACE AND INITIALIZATION
// 	// O(n)
// 	arr := make([]int, n)
// 	size := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		arr[i] = i
// 		size[i] = 1
// 	}
// 	return UnionFind{
// 		Array: arr,
// 		Size:  size,
// 	}
// }

// func Root(array []int, i int) int {
// 	for i != array[i] {
// 		// path compression
// 		array[i] = array[array[i]]
// 		i = array[i]

// 	}
// 	return i

// }

// // func (u UnionFind) connected(p, q int) bool {
// // 	// O(n+n) = O(n)

// // 	return Root(u.Array, p) == Root(u.Array, q)
// // }

// func (u UnionFind) Max() int {
// 	max := -1
// 	// highest := 0
// 	for i := 0; i < len(u.Size); i++ {
// 		if max < u.Size[i] {
// 			max = u.Size[i]

// 		}
// 	}
// 	return max
// }

// func (u UnionFind) Union(p, q int) {
// 	// here just one value change
// 	// this can take O(N) + the two roots computation
// 	rootP := Root(u.Array, p)
// 	rootQ := Root(u.Array, q)
// 	if rootP == rootQ {
// 		return // They are already in the same set
// 	}
// 	// get the size of each root the append the lower to the higher
// 	if u.Size[rootP] > u.Size[rootQ] {
// 		u.Array[rootQ] = rootP
// 		u.Size[rootP] += u.Size[rootQ]
// 	} else {
// 		u.Array[rootP] = rootQ
// 		u.Size[rootQ] += u.Size[rootP]

// 	}

// }

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uf := NewUnionFind(len(nums))
	valueToIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		valueToIndex[num] = i
	}
	for _, num := range nums {
		if _, exists := valueToIndex[num-1]; exists {
			uf.Union(valueToIndex[num], valueToIndex[num-1])
		} else if _, exists := valueToIndex[num+1]; exists {
			uf.Union(valueToIndex[num], valueToIndex[num+1])
		}
	}

	return uf.Max()
}
