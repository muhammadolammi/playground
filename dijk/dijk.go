package dijk

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

// this will return a map of string of nodes to their shortest distance from the soruces
func Dijkstra(graph GRAPH, src string) (map[string]int, error) {
	_, ok := graph.Data[src]
	if !ok {
		return nil, errors.New("src not  in graph")
	}

	distances := map[string]int{}
	// we need unvisited set , to keep track of what we havent visited.
	// we can use map to mirrow set

	unvisited := map[string]string{}
	nodes := graph.Data

	// set all distances to infinity for each node
	// then add all nodes to unvisited
	for node := range nodes {
		distances[node] = math.MaxInt
		unvisited[node] = node
	}
	distances[src] = 0

	// while theres still a node not visited

	for len(unvisited) != 0 {
		// find the univisited node with shortest distances
		nodeWithShortestPath := unvistedNodeWithShortestPath(unvisited, distances)

		// mark it as visited by deleting it from the map
		delete(unvisited, nodeWithShortestPath)

		// relax the nodeWithShortestPath children nodes
		for neigborNode, dist := range nodes[nodeWithShortestPath] {
			// check if the node wasnt visited before

			if _, ok := unvisited[neigborNode]; ok {
				if dist+distances[nodeWithShortestPath] < distances[neigborNode] {
					distances[neigborNode] = dist + distances[nodeWithShortestPath]
				}
			}

		}
	}

	return distances, nil

}

func DijkstraSrcToDest(graph GRAPH, src, dest string) ([]string, error) {
	_, ok := graph.Data[src]
	if !ok {
		return nil, errors.New("src not  in graph")
	}
	_, ok = graph.Data[dest]
	if !ok {
		return nil, errors.New("dest not  in graph")
	}
	distances := map[string]int{}
	// we need unvisited set , to keep track of what we havent visited.
	// we can use map to mirrow set

	unvisited := map[string]string{}
	// predecessors map, this map a node to it predecessor
	// For example, a: b means that b leads to a
	predecessors := map[string]string{}
	nodes := graph.Data

	// set all distances to infinity for each node
	// then add all nodes to unvisited
	for node := range nodes {
		distances[node] = math.MaxInt
		unvisited[node] = node
	}
	distances[src] = 0

	// while theres still a node not visited

	for len(unvisited) != 0 {
		// find the univisited node with shortest distances
		nodeWithShortestPath := unvistedNodeWithShortestPath(unvisited, distances)
		fmt.Println(nodeWithShortestPath)

		// mark it as visited by deleting it from the map
		delete(unvisited, nodeWithShortestPath)
		if nodeWithShortestPath == dest {
			return getPath(dest, predecessors), nil
		}

		// relax the nodeWithShortestPath children nodes
		for neigborNode, dist := range nodes[nodeWithShortestPath] {
			if _, ok := unvisited[neigborNode]; ok {
				if dist+distances[nodeWithShortestPath] < distances[neigborNode] {
					distances[neigborNode] = dist + distances[nodeWithShortestPath]
					// since we found a short path , we need to update the predecessor of the node
					predecessors[neigborNode] = nodeWithShortestPath

				}
			}

		}
	}

	return nil, nil

}

func unvistedNodeWithShortestPath(unvisited map[string]string, distances map[string]int) string {
	nodeWithShortestPath := ""
	shortestPath := math.MaxInt
	for node := range unvisited {
		if distances[node] < shortestPath {
			shortestPath = distances[node]
			nodeWithShortestPath = node
		}

	}
	return nodeWithShortestPath

}

func getPath(dest string, predecessors map[string]string) []string {
	ls := []string{}
	next := dest
	for {
		ls = append(ls, next)
		if _, ok := predecessors[next]; ok {
			next = predecessors[next]
		} else {
			break
		}

	}
	slices.Reverse(ls)
	return ls
}
