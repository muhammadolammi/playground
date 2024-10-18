package heap

import "fmt"

const MINHEAP = "MINHEAP"
const MAXHEAP = "MAXHEAP"

type Node struct {
	Priority int
	Value    string
}
type Heap struct {
	Data     []Node
	HeapType string
}

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}
