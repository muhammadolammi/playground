package main

func isValid(s string) bool {
	pairs := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := Stack{}
	for _, r := range s {

		currentstring := string(r)
		if _, ok := pairs[currentstring]; ok {
			stack.Push(currentstring)
		} else if stack.Size() == 0 || currentstring != pairs[stack.Pop()] {
			return false
		}

		// fmt.Println(string(r))
	}
	return stack.Size() == 0

}

type Stack struct {
	Array []string
}

func (s *Stack) Push(item string) {
	s.Array = append(s.Array, item)
}

func (s *Stack) Pop() string {
	if len(s.Array) == 0 {

		return ""
	}
	popped := s.Array[len(s.Array)-1]
	// Remove the popped item
	s.Array = s.Array[:len(s.Array)-1]

	return popped
}

func (s *Stack) Peek() string {
	if len(s.Array) == 0 {

		return "0"
	}
	return s.Array[len(s.Array)-1]
}
func (s *Stack) Size() int {
	return len(s.Array)
}
