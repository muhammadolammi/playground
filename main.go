package main

import (
	"fmt"
	"goplayground/astarsearch"

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
	// ch1 := make(chan int, 10)
	// go Walk(tree.New(1), ch1)
	// for v := range ch1 {
	// 	fmt.Println(v)
	// }

	// fmt.Println(Same(tree.New(1), tree.New(1)))
	// fmt.Println(Same(tree.New(1), tree.New(2)))
	// PrintTreeAddresses(tree.New(1))

	// fmt.Println(func1())

	// f2 := func2()
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(f2())
	// }
	// f3 := func3("Muhammad")
	// fmt.Println(f3())

	// f4 := func4("Muhammad")
	// fmt.Println(f4("Ogun"))
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
	// bt.Insert(500)
	// // should print 2
	// fmt.Println(bt.GetMin())
	// // should print 500
	// fmt.Println(bt.GetMax())
	// p := bt.Search(1000)
	// fmt.Println(p)

	// a := bt.GetArrayRepresentation()

	// fmt.Println(a)

	// ll := linkedlist.LinkedList{}
	// ll.Insert(5)
	// ll.Insert(10)
	// ll.Insert(15)
	// ll.Insert(20)
	// ll.Tranverse()
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())
	// fmt.Println(st.Pop())

	// fmt.Println(st.Size())
	// l1 := []int{1, 2, 3}
	// l2 := []int{2, 5, 7, 8, 1, 3, 3, 2, 54, 60, 50, 40, 33, 24, 10}
	// l3 := mergesort.Merge(l1, l2)
	// fmt.Println(l3)
	// l4 := mergesort.MergeM(l1, l2, l3, []int{100, 500, 10000})
	// l5 := mergesort.MergeSort(l2)

	// fmt.Println(l4)
	// fmt.Println(l5)

	// // Example matrices for testing
	// A := &matrice.Matrice{
	// 	R: 2,
	// 	C: 2,
	// 	Rows: []matrice.Row{
	// 		{1, 2},
	// 		{3, 4},
	// 	},
	// }

	// B := &matrice.Matrice{
	// 	R: 2,
	// 	C: 2,
	// 	Rows: []matrice.Row{
	// 		{5, 6},
	// 		{7, 8},
	// 	},
	// }

	// result, err := A.SMM(B)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Result of Strassen's multiplication for 4x4 matrices:")
	// for _, row := range result.Rows {
	// 	fmt.Println(row)

	// }
	// E := &matrice.Matrice{
	// 	R: 4,
	// 	C: 4,
	// 	Rows: []matrice.Row{
	// 		{1, 2, 3, 4},
	// 		{5, 6, 7, 8},
	// 		{9, 10, 11, 12},
	// 		{13, 14, 15, 16},
	// 	},
	// }

	// F := &matrice.Matrice{
	// 	R: 4,
	// 	C: 4,
	// 	Rows: []matrice.Row{
	// 		{16, 15, 14, 13},
	// 		{12, 11, 10, 9},
	// 		{8, 7, 6, 5},
	// 		{4, 3, 2, 1},
	// 	},
	// }

	// result, err = E.SMM(F)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Result of Strassen's multiplication for 2x2 matrices:")
	// for _, row := range result.Rows {
	// 	fmt.Println(row)
	// }

	// l := twoSum([]int{2, 7, 11, 15}, 9)
	// fmt.Println(l)

	// [5,6,4]
	// [2,4,3]

	// l1 := ListNode{Val: 5}
	// l2 := ListNode{Val: 6}
	// l3 := ListNode{Val: 4}

	// l4 := ListNode{Val: 2}
	// l5 := ListNode{Val: 4}
	// l6 := ListNode{Val: 3}

	// l1.Next = &l2
	// l2.Next = &l3

	// l4.Next = &l5
	// l5.Next = &l6

	// l := addTwoNumbers(&l1, &l4)
	// fmt.Println(l)

	// var p *int
	// i := 5
	// p = &i
	// fmt.Println(p)

	// sum := hexaDecimalToDecimal(p)
	// fmt.Println(sum)

	// valid := isValid("({[]})")

	// notValid := isValid("[")

	// fmt.Println(valid)
	// fmt.Println(notValid)

	// l1 := ListNode{Val: 1}
	// l2 := ListNode{Val: 2}
	// l3 := ListNode{Val: 4}

	// l1.Next = &l2
	// l2.Next = &l3

	// l4 := ListNode{Val: 1}
	// l5 := ListNode{Val: 3}
	// l6 := ListNode{Val: 4}

	// l4.Next = &l5
	// l5.Next = &l6

	// re := mergeTwoLists(nil, &l4)

	// tranverlinkedlist(re)

	// nums := []int{2, 3, 0, 1, 4}

	// n := jump(nums)
	// fmt.Println(n)

	// n := dp.PaidStairCasesPath(8, []int{0, 3, 2, 4, 6, 1, 1, 5, 3})

	// s := kata.Solution(600)
	// fmt.Println(s)
	// s := uncommonFromSentences("apple apple", "banana")
	// fmt.Println(s)
	// board := [][]byte{
	// 	{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
	// 	{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
	// 	{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
	// 	{'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'},
	// 	{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
	// 	{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'},
	// 	{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'},
	// 	{'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'},
	// 	{'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'},
	// 	{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'},
	// }

	// solve(board)

	// for _, row := range board {
	// 	fmt.Println(string(row))
	// }
	// grid := [][]byte{
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '0', '1', '0', '1'},
	// 	{'1', '1', '1', '0', '1'},
	// }
	// n := numIslands(grid)
	// fmt.Println(n)

	// uf := unionfind.NewUnionFind(10)
	// uf.Union(1, 2)
	// uf.Union(1, 6)
	// uf.Union(1, 9)
	// uf.Union(3, 5)
	// fmt.Println(uf.Find(2))

	// l1 := &ListNode{
	// 	Val: 2, Next: &ListNode{
	// 		Val: 4, Next: &ListNode{
	// 			Val: 9, Next: nil},
	// 	},
	// }

	// // Create l2 = [5,6,4]
	// l2 := &ListNode{
	// 	Val: 5, Next: &ListNode{
	// 		Val: 6, Next: &ListNode{
	// 			Val: 4, Next: &ListNode{
	// 				Val: 9, Next: nil,
	// 			},
	// 		},
	// 	},
	// }
	// l := addTwoNumbers(l1, l2)
	// tranverlinkedlist(l)

	// nums1 := []int{1, 3}
	// nums2 := []int{2}
	//CORRECT ANSWER 16
	// nums1 := []int{1, 12, 15, 26, 38}
	// nums2 := []int{2, 13, 17, 30, 45}
	// nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// nums2 := []int{1, 2, 3, 4, 5}

	// median := findMedianSortedArrays(nums1, nums2)
	// fmt.Printf("Median: %v\n", median)

	// Create a sample graph
	// graph := dijk.GRAPH{
	// 	Data: map[string]map[string]int{
	// 		"A": {"B": 3, "D": 5},
	// 		"B": {},
	// 		"C": {"B": -10},
	// 		"D": {"C": 2},
	// 	},
	// }

	// // Choose the source node
	// source := "A"

	// // Call Dijkstra’s algorithm
	// distances, err := dijk.Dijkstra(graph, source)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Print the shortest distances from the source node to each other node
	// fmt.Printf("Shortest distances from node %s:\n", source)
	// for node, dist := range distances {
	// 	fmt.Printf("%s -> %d\n", node, dist)
	// }
	// fmt.Println(distances)
	// path, err := dijk.DijkstraSrcToDest(graph, "A", "B")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(path)
	// graph := bellmanford.GRAPH{
	// 	Data: map[string]map[string]int{
	// 		"A": {"B": -4, "D": 3},
	// 		"B": {"C": -5},
	// 		"C": {"A": 2, "D": 2},
	// 		"D": {},
	// 		"E": {"A": 2, "D": 1},
	// 	},
	// }
	// shortestPath, err := bellmanford.BellmanFord(graph, "A", "C")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(shortestPath)
	// heap := heap.Heap{
	// 	HeapType: heap.MAXHEAP,
	// }
	// heap.Insert(21)
	// heap.Insert(17)
	// heap.Insert(14)
	// heap.Insert(15)
	// heap.Insert(1)
	// heap.Insert(12)
	// heap.Insert(13)
	// heap.Insert(5)
	// heap.Insert(9)
	// heap.Insert(8)
	// heap.Insert(50000)
	// heap.Pop()
	// fmt.Println(heap.Data)
	// fmt.Println(heap.Peek())

	// pq := pq.NewPriorityQueue()

	// pq.Insert(heap.Node{
	// 	Priority: 21,
	// 	Value:    "A",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 17,
	// 	Value:    "B",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 14,
	// 	Value:    "C",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 15,
	// 	Value:    "D",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 1,
	// 	Value:    "E",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 12,
	// 	Value:    "F",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 13,
	// 	Value:    "G",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 5,
	// 	Value:    "H",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 9,
	// 	Value:    "I",
	// })
	// pq.Insert(heap.Node{
	// 	Priority: 5000,
	// 	Value:    "J",
	// })
	// fmt.Println(pq.Data())
	grid := astarsearch.Grid{
		Tiles: [][]astarsearch.Tile{
			{{0, 0, 1}, {1, 0, 1}, {2, 0, 5}, {3, 0, 1}, {4, 0, 1}},
			{{0, 1, 1}, {1, 1, 3}, {2, 1, 1}, {3, 1, 1}, {4, 1, 5}},
			{{0, 2, 5}, {1, 2, 1}, {2, 2, 1}, {3, 2, 3}, {4, 2, 1}},
			{{0, 3, 1}, {1, 3, 5}, {2, 3, 1}, {3, 3, 1}, {4, 3, 1}},
			{{0, 4, 1}, {1, 4, 1}, {2, 4, 5}, {3, 4, 1}, {4, 4, 1}},
		},
		Width:  5,
		Height: 5,
	}
	src := astarsearch.Tile{
		X: 0,
		Y: 0,
	}
	dest := astarsearch.Tile{
		X: 4,
		Y: 4,
	}

	path := astarsearch.AstarSearch(grid, src, dest)
	fmt.Println(path)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	l := &ListNode{}
// 	helper(l, l1, l2)
// 	return l

// }

// func helper(l, l1, l2 *ListNode) {
// 	if l1.Next == nil {

// 		l.Val = l1.Val + l2.Val
// 		return
// 	}
// 	if l2.Next == nil {
// 		l2.Next = &ListNode{Val: 0}
// 	}
// 	helper(l, l1.Next, l2.Next)
// }

// func twoSum(nums []int, target int) []int {
// 	seen := map[int]int{}
// 	for i := range nums {
// 		diff := target - nums[i]
// 		n, ok := seen[diff]
// 		if ok {
// 			return []int{n, i}
// 		}
// 		//    We are saving the diffs here as a key and the index as the value
// 		seen[nums[i]] = i
// 	}
// 	fmt.Println(seen)
// 	return []int{}
// }
