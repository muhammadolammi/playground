package astarsearch

import (
	"fmt"
	"goplayground/heap"
	"goplayground/pq"
	"log"
	"math"
	"slices"
	"strings"
)

func heuristic(next, dest Tile) int {
	abs := math.Abs(float64(next.X)-float64(dest.X)) + math.Abs(float64(next.Y)-float64(dest.Y))
	return int(abs)
}

func AstarSearch(grid Grid, src, dest Tile) []string {

	pq := pq.NewPriorityQueue()
	// we will never revisit a node again
	// and we dont need to visit all so no need to add all to visited
	visited := map[string]string{}
	costs := map[string]int{}
	// unlike dijistra you add just the src  the distances
	// we will get the distances from the loop instead of setting it to infinity
	// same way we will get neccesary node to push to pq from then loop
	// initially pushing just the src with priority zero
	predecessor := map[string]string{}
	pq.Insert(heap.Node{
		Value:    src.String(),
		Priority: 0,
	})
	// visited[src.String()] = src.String()
	costs[src.String()] = 0
	// while the priority queue is not empty
	for !pq.IsEmpty() {

		nextNode, err := pq.Pop()
		if err != nil {
			fmt.Println("ran")
			log.Panicf("error getting next node err: %v", err)
		}
		// if next node is destination break the loop
		if nextNode.Value == dest.String() {

			break
		}
		// if next node is already visited continue
		if _, ok := visited[nextNode.Value]; ok {
			fmt.Println("ran")
			continue
		}
		visited[nextNode.Value] = nextNode.Value
		for _, nei := range grid.neigbors(nodeStringToTile(grid, nextNode.Value)) {
			neiTotalCost := costs[nextNode.Value] + nei.Cost
			if _, ok := costs[nei.String()]; !ok || costs[nei.String()] < neiTotalCost {
				costs[nei.String()] = neiTotalCost
				neiPriority := neiTotalCost + heuristic(nodeStringToTile(grid, nextNode.Value), nei)
				pq.Insert(heap.Node{
					Value:    nei.String(),
					Priority: neiPriority,
				})
				predecessor[nei.String()] = nextNode.Value
			}
		}

	}

	path := []string{}
	pred := dest.String()
	for pred != "" {
		path = append(path, pred)
		pred = predecessor[pred]
	}
	slices.Reverse(path)
	return path
}

func nodeStringToTile(grid Grid, nodestring string) Tile {
	nodestring = strings.Trim(nodestring, "()")
	var x, y int
	fmt.Sscanf(nodestring, "%d,%d,%d", &x, &y)

	return grid.Tiles[y][x]
}
