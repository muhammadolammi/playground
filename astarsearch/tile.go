package astarsearch

import (
	"fmt"
)

//	func (t *Tile) cost() int {
//		// generate a random cost for the tile
//		// i.e the cost of traveling the tile/node
//		return rand.Intn(20)
//	}
func (t *Tile) equal(other Tile) bool {
	if other.X == t.X && other.Y == t.Y {
		return true
	}
	return false
}
func (t *Tile) String() string {
	return fmt.Sprintf("(%d,%d)", t.X, t.Y)
}
