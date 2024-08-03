package linkedlist

import "fmt"

func insertHelper(value int, n *Node) {
	if n.Child == nil {
		n.Child = &Node{Value: value, Parent: n}
		return
	}
	insertHelper(value, n.Child)
}
func (ll *LinkedList) Insert(value int) {
	if ll.Root == nil {
		ll.Root = &Node{Value: value}
		return
	}
	insertHelper(value, ll.Root)
}

func (ll *LinkedList) Tranverse() {
	current := ll.Root
	for current != nil {
		fmt.Printf("%d \n", current.Value)
		current = current.Child
	}
}
