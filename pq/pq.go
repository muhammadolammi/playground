package pq

import "goplayground/heap"

func NewPriorityQueue() PriorityQueue {

	return PriorityQueue{
		heap: heap.Heap{
			HeapType: heap.MINHEAP,
		},
	}
}

// find a way to pop and insert based on the priority of the node
func (pq *PriorityQueue) Pop() (heap.Node, error) {
	return pq.heap.Pop()

}
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Data()) == 0
}
func (pq *PriorityQueue) Insert(node heap.Node) {
	pq.heap.Insert(node)

}

func (pq *PriorityQueue) Data() []heap.Node {
	return pq.heap.Data
}
