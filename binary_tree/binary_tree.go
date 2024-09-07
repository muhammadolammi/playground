package binary_tree

func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}
	t.Root.insert(value)
}

func (t *Node) insert(value int) {

	if value < t.Value {
		if t.Left == nil {
			t.Left = &Node{Value: value}
			return
		}
		t.Left.insert(value)
	}
	if value > t.Value {
		if t.Right == nil {
			t.Right = &Node{Value: value}
			return
		}
		t.Right.insert(value)
	}

}
func getArrayrecursively(a *[]int, n *Node) {
	if n == nil {
		return
	}
	if n.Left != nil {
		getArrayrecursively(a, n.Left)
	}
	*a = append(*a, n.Value)
	if n.Right != nil {
		getArrayrecursively(a, n.Right)
	}
}
func (t *BinaryTree) GetArrayRepresentation() []int {
	a := []int{}
	getArrayrecursively(&a, t.Root)
	return a
}

func (t *BinaryTree) Search(value int) *Node {
	if t.Root.Left == nil && t.Root.Right == nil {
		if value == t.Root.Value {
			return t.Root
		}
		return nil
	}
	if value < t.Root.Value {

		return searchHelper(value, t.Root.Left)
	}

	return searchHelper(value, t.Root.Right)

}
func searchHelper(value int, n *Node) *Node {
	if value == n.Value {
		return n
	}
	// If its the last node and value not found that means it doest exit
	if n.Right == nil && n.Left == nil {
		return nil
	}

	if value < n.Value {

		return searchHelper(value, n.Left)
	}

	return searchHelper(value, n.Right)

}
func (t *BinaryTree) GetMin() int {

	if t.Root.Left == nil {
		return t.Root.Value
	}
	return getMinHelper(t.Root)

}

func getMinHelper(n *Node) int {
	if n.Left == nil {
		return n.Value
	}
	return getMinHelper(n.Left)
}

func (t *BinaryTree) GetMax() int {

	if t.Root.Right == nil {
		return t.Root.Value
	}
	return getMaxHelper(t.Root)

}

func getMaxHelper(n *Node) int {
	if n.Right == nil {
		return n.Value
	}
	return getMaxHelper(n.Right)
}
