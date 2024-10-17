package bellmanford

import (
	"fmt"
	"math"
)

func BellmanFord(graph GRAPH, src, dest string) (int, error) {
	distances := map[string]int{}
	nodes := graph.Data

	for node := range nodes {
		distances[node] = math.MaxInt
	}
	distances[src] = 0

	// relax each edges for n-1 time
	n := len(graph.Data)
	for i := 0; i < n-1; i++ {

		for node, neighborsCost := range nodes {
			if distances[node] == math.MaxInt {
				continue // Skip unreachable nodes
			}
			for neighbor := range neighborsCost {
				newDistance := distances[node] + neighborsCost[neighbor]
				if newDistance < distances[neighbor] {
					distances[neighbor] = newDistance
				}
			}
		}

	}
	// check for relaxation again, any relaxing instance means theres a negative cycle, which means the algorithm will keep decreasing forever

	for node, neighborsCost := range nodes {
		if distances[node] == math.MaxInt {
			continue // Skip unreachable nodes
		}
		for neighbor := range neighborsCost {

			newDistance := distances[node] + neighborsCost[neighbor]
			if newDistance < distances[neighbor] {
				return 0, fmt.Errorf("negative Cycle Detected at node %s -> %s -> %s", node, neighbor, detectThirdNode(node, neighbor, nodes))
			}
		}
	}
	return distances[dest], nil
}

func detectThirdNode(node1, node2 string, nodes map[string]map[string]int) string {
	node2Neighbors := nodes[node2]
	for node := range node2Neighbors {
		if _, exist := nodes[node][node1]; exist {
			return node
		}
	}
	return ""
}
