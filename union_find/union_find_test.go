package unionfind_test

import (
	unionfind "goplayground/union_find"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := unionfind.NewUnionFind(7)
	//[0, 1, 2, 3, 4]

	uf.Union(2, 4)
	//[0, 1, 4, 3, 4]

	uf.Union(2, 3)
	//[0, 1, 4, 2, 4]

	uf.Union(1, 2)

	//[0, 4, 4, 4, 4]

	uf.Union(1, 6)

}
