package astarsearch

type Tile struct {
	X    int
	Y    int
	Cost int
}
type Grid struct {
	Tiles  [][]Tile
	Width  int
	Height int
}
