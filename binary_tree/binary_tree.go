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
