package wire

// ManhattanDistance returns the Manhattan distance between two points.
func ManhattanDistance(p1, p2 Point) int {
	dx := p1.x - p2.x
	if dx < 0 {
		dx = -dx
	}
	dy := p1.y - p2.y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}
