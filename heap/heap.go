package heap

func (h *Heap) Insert(node Node) {
	h.Data = append(h.Data, node)
	h.sortUp(h.n() - 1)

}
func (h *Heap) Pop() (Node, error) {
	n := h.n()
	if n == 0 {
		// TODO represent with something better to indicate empty heap, maybe an error
		return Node{}, &CustomError{
			Code:    404,
			Message: "Empty Graph",
		}
	}

	element := h.Data[0]
	if n == 1 {
		h.Data = []Node{}
		return element, nil
	}

	lastElement := h.Data[n-1] // get last element
	//move the last element to the top of the heap
	h.Data[0] = lastElement
	// delete the last element node in the heap
	h.Data = h.Data[:n-1]
	h.heapifyDown(0)

	return element, nil
}

func (h *Heap) heapifyDown(parentIndex int) {
	n := h.n()
	parentNode := h.Data[parentIndex]

	leftIndex := (2 * parentIndex) + 1

	rightIndex := (2 * parentIndex) + 2

	if h.HeapType == MINHEAP {
		// lets check if left is a valid index and less than the parent value.
		if leftIndex < n && h.Data[leftIndex].Priority < parentNode.Priority {
			h.Data[parentIndex] = h.Data[leftIndex]
			h.Data[leftIndex] = parentNode
			h.heapifyDown(leftIndex)
		}
		// lets check if right is a valid index and less than the parent value.
		if rightIndex < n && h.Data[rightIndex].Priority < parentNode.Priority {
			h.Data[parentIndex] = h.Data[rightIndex]
			h.Data[rightIndex] = parentNode
			h.heapifyDown(rightIndex)
		}

	}
	if h.HeapType == MAXHEAP {
		largestChild := 0
		// lets find largest child between right and left
		if leftIndex < n && h.Data[leftIndex].Priority > parentNode.Priority {
			largestChild = leftIndex

		}
		if rightIndex < n && h.Data[rightIndex].Priority > h.Data[leftIndex].Priority {
			largestChild = rightIndex
		}
		if h.Data[largestChild].Priority > parentNode.Priority {
			h.Data[parentIndex] = h.Data[largestChild]
			h.Data[largestChild] = parentNode
			h.heapifyDown(largestChild)
		}
	}

}

func (h *Heap) sortUp(index int) {
	// this will traverse the tree/list up to find if the just inserted value is greater than or equals to
	if index == 0 {
		return
	}
	// get the parent of just inserted value/node
	parentIndex := 0
	parentIndex = (index - 1) / 2
	indexValue := h.Data[index]
	parentValue := h.Data[parentIndex]

	if h.HeapType == MINHEAP {
		// sort up based on the lesser value
		if indexValue.Priority < parentValue.Priority {
			h.Data[index] = parentValue
			h.Data[parentIndex] = indexValue
			h.sortUp(parentIndex)
		}

	}
	if h.HeapType == MAXHEAP {
		// sort up based on the higher value
		if indexValue.Priority > parentValue.Priority {
			h.Data[index] = parentValue
			h.Data[parentIndex] = indexValue
			h.sortUp(parentIndex)
		}
	}

}

// will return the min/max element in the heap based on the heap type
func (h *Heap) Peek() (Node, error) {
	if h.n() == 0 {
		return Node{}, &CustomError{
			Code:    404,
			Message: "Empty Graph",
		}
	}
	return h.Data[0], nil
}

func (h *Heap) n() int {
	return len(h.Data)
}
