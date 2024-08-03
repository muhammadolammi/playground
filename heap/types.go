package heap

import "fmt"

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type Heap struct {
	Root     *Node
	HeapType string
}

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}
