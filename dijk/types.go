package dijk

type GRAPH struct {
	// this data map a node to a map that map the node neigbors to their weight
	Data map[string]map[string]int
}
