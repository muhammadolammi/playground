package linkedlist

type Node struct {
	Child  *Node
	Parent *Node
	Value  int
}

type LinkedList struct {
	Root *Node
}
