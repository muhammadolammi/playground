package astarsearch

func (grid *Grid) inBound(tile Tile) bool {
	if tile.X >= 0 && tile.X < grid.Width && tile.Y >= 0 && tile.Y < grid.Height {
		return true
	}
	return false

}
func (grid *Grid) neigbors(tile Tile) []Tile {
	possibleNeigbors := make([]Tile, 4)
	// saving all possible neigbors with the order
	//    N
	// W     E
	//    S
	// following anti clockwise order
	// north east south west in the list

	// NORTH
	possibleNeigbors[0] = Tile{
		X: tile.X,
		Y: tile.Y - 1,
	}
	// EAST
	possibleNeigbors[1] = Tile{
		X: tile.X + 1,
		Y: tile.Y,
	}
	// SOUTH
	possibleNeigbors[2] = Tile{
		X: tile.X,
		Y: tile.Y + 1,
	}
	// WEST
	possibleNeigbors[3] = Tile{
		X: tile.X - 1,
		Y: tile.Y,
	}
	result := []Tile{}
	for _, t := range possibleNeigbors {
		if grid.inBound(t) {
			result = append(result, t)
		}
	}
	return result
}
