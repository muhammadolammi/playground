package heap

import "fmt"

const MINHEAP = "MINHEAP"
const MAXHEAP = "MAXHEAP"

type Heap struct {
	Data     []int
	HeapType string
}

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}
