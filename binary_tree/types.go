package binary_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree represents the binary tree itself
type BinaryTree struct {
	Root *Node
}
