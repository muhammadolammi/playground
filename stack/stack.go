package stack

func (s *Stack) Push(item int) {
	s.Array = append(s.Array, item)
}

func (s *Stack) Pop() int {
	if len(s.Array) == 0 {

		return 0
	}
	popped := s.Array[len(s.Array)-1]
	// Remove the popped item
	s.Array = s.Array[:len(s.Array)-1]

	return popped
}

func (s *Stack) Peek() int {
	if len(s.Array) == 0 {

		return 0
	}
	return s.Array[len(s.Array)-1]
}
func (s *Stack) Size() int {
	return len(s.Array)
}
