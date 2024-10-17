package heap

func (h *Heap) Insert(value int) {
	h.Data = append(h.Data, value)
	h.sortUp(len(h.Data) - 1)

}
func (h *Heap) Pop() int {
	n := len(h.Data)
	if n == 0 {
		// TODO represent with something better to indicate empty heap, maybe an error
		return -999
	}

	element := h.Data[0]
	if n == 1 {
		h.Data = []int{}
		return element
	}

	lastElement := h.Data[n-1] // get last element
	//move the last element to the top of the heap
	h.Data[0] = lastElement
	// delete the last element node in the heap
	h.Data = h.Data[:n-1]
	h.sortDown(0)

	return element
}

func (h *Heap) sortDown(parentIndex int) {
	n := len(h.Data)
	parentValue := h.Data[parentIndex]

	leftIndex := (2 * parentIndex) + 1

	rightIndex := (2 * parentIndex) + 2

	if h.HeapType == MINHEAP {
		// lets check if left is a valid index and less than the parent value.
		if leftIndex < n && h.Data[leftIndex] < parentValue {
			h.Data[parentIndex] = h.Data[leftIndex]
			h.Data[leftIndex] = parentValue
			h.sortDown(leftIndex)
		}
		// lets check if right is a valid index and less than the parent value.
		if rightIndex < n && h.Data[rightIndex] < parentValue {
			h.Data[parentIndex] = h.Data[rightIndex]
			h.Data[rightIndex] = parentValue
			h.sortDown(rightIndex)
		}

	}
	if h.HeapType == MAXHEAP {
		if leftIndex < n && h.Data[leftIndex] > parentValue {
			largestChild := leftIndex
			if rightIndex < n && h.Data[rightIndex] > h.Data[leftIndex] {
				largestChild = rightIndex
			}
			if h.Data[largestChild] > parentValue {
				h.Data[parentIndex] = h.Data[largestChild]
				h.Data[largestChild] = parentValue
				h.sortDown(largestChild)
			}
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
		if indexValue < parentValue {
			h.Data[index] = parentValue
			h.Data[parentIndex] = indexValue
			h.sortUp(parentIndex)
		}

	}
	if h.HeapType == MAXHEAP {
		// sort up based on the higher value
		if indexValue > parentValue {
			h.Data[index] = parentValue
			h.Data[parentIndex] = indexValue
			h.sortUp(parentIndex)
		}
	}

}

// will return the min/max element in the heap based on the heap type
func (h *Heap) Peek() int {
	return h.Data[0]
}
