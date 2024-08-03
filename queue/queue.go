package queue

func (q *Queue) Push(item int) {
	q.Array = append(q.Array, item)
}

func (q *Queue) Pop() int {
	if len(q.Array) == 0 {

		return 0
	}
	popped := q.Array[0]
	// Remove the popped item
	q.Array = q.Array[1:]

	return popped
}

func (q *Queue) Peek() int {
	if len(q.Array) == 0 {

		return 0
	}
	return q.Array[0]
}
func (q *Queue) Size() int {
	return len(q.Array)
}
