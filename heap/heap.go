package heap

func sort(value int, node *Node) {
	if node.Left == nil && node.Right == nil {
		node.Value = value
	}

}

func (h *Heap) Insert(value int) {
	if h.Root == nil {
		h.Root = &Node{
			Value: value,
		}
	}

}
