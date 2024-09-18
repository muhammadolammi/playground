package unionfind

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

// This first implementation is the quick-find but the Union is very expensive.
// func (u UnionFind) Find(p, q int) bool {
// 	// O(1)
// 	return u.Array[p] == u.Array[q]
// }

// func (u UnionFind) Union(p, q int) {
// 	//for this implementation. an idex with a value in that index means its in that set
// 	// array[index] is the set/component the index belongs to.

// 	// to perferom union, we set p and it member to set qid
// 	// qid == array[q]
// 	// O(n)
// 	pid := u.Array[p]
// 	qid := u.Array[q]
// 	for i := 0; i < len(u.Array); i++ {
// 		if u.Array[i] == pid {
// 			u.Array[i] = qid
// 		}
// 	}
// }

// This implementaion is for quick union

// We use tree for this implementation
// u.Array[i] is the parent of i
// the root is the index where u.Array[i]=i

func root(array []int, i int) int {
	for i != array[i] {
		// path compression
		array[i] = array[array[i]]
		i = array[i]

	}
	return i

}

// for find , we check if root(q) == root(p)
func (u UnionFind) connected(p, q int) bool {
	// O(n+n) = O(n)

	return root(u.Array, p) == root(u.Array, q)
}

// since we use tree in the quick-union
// we set the root of p root to the root of q
// i.e we join the tree of p to tree of q
// the root of q becomes the new root of the two elements
// func (u UnionFind) Union(p, q int) {
// 	// here just one value change
// 	// this can take O(N) + the two roots computation
// 	rootP := root(u.Array, p)
// 	rootQ := root(u.Array, q)
// 	u.Array[rootP] = rootQ
// }

// we still have an issue with our quick union, they get worst compare to quick find
// Tree can also get worse , where they reach N
// So we need to work on making our union better, i.e keep our trees flat.
// we do this by always adding the smaller tree to the bigger.
// we we try joining two trees, we check for the taller/bigger tree and we append or join the smaller tree to it.

// For this, our join will be between two UNIONFIND struct to make one.
// we need to create an array of size with same len as the normal array
// root and find stay the same.

// lets rewrite union.

func (u UnionFind) Union(p, q int) {
	// here just one value change
	// this can take O(N) + the two roots computation
	rootP := root(u.Array, p)
	rootQ := root(u.Array, q)
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
