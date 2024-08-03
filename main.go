package main

import (
	"fmt"
	"goplayground/linkedlist"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Helper function to recursively walk the tree
func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v2, ok := <-ch2; !ok || v1 != v2 {
			return false
		}

	}
	_, ok := <-ch2
	return !ok
	// return true
}

func PrintTreeAddresses(t *tree.Tree) {
	if t.Left != nil {
		PrintTreeAddresses(t.Left)
	}
	fmt.Println(&t)
	if t.Right != nil {
		PrintTreeAddresses(t.Right)
	}
}

func main() {
	ch1 := make(chan int, 10)
	go Walk(tree.New(1), ch1)
	for v := range ch1 {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	PrintTreeAddresses(tree.New(1))

	fmt.Println(func1())

	f2 := func2()
	for i := 0; i < 5; i++ {
		fmt.Println(f2())
	}
	f3 := func3("Muhammad")
	fmt.Println(f3())

	f4 := func4("Muhammad")
	fmt.Println(f4("Ogun"))
	// l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	// in := binary_search.BinarySearch(l, 10, 0, len(l)-1)
	// fmt.Println(in)

	// bt := binary_tree.BinaryTree{
	// 	Root: &binary_tree.Node{
	// 		Value: 5,
	// 	},
	// }
	// bt.Insert(10)
	// bt.Insert(9)
	// bt.Insert(2)
	// bt.Insert(3)
	// bt.Insert(20)
	// bt.Insert(30)
	// a := bt.GetArrayRepresentation()

	// fmt.Println(a)

	ll := linkedlist.LinkedList{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)
	ll.Insert(20)
	ll.Tranverse()
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())

	// fmt.Println(st.Size())
}
