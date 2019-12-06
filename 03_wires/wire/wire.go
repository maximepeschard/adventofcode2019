package wire

func movePoints(start Point, startValue int, m Move) (PointSet, Point, int) {
	var final Point
	points := NewPointSet()

	switch m.Direction {
	case Up:
		for i := 0; i < m.Distance; i++ {
			final = Point{x: start.x, y: start.y + i + 1}
			points.Update(final, startValue+i+1)
		}
	case Down:
		for i := 0; i < m.Distance; i++ {
			final = Point{x: start.x, y: start.y - (i + 1)}
			points.Update(final, startValue+i+1)
		}
	case Right:
		for i := 0; i < m.Distance; i++ {
			final = Point{x: start.x + i + 1, y: start.y}
			points.Update(final, startValue+i+1)
		}
	case Left:
		for i := 0; i < m.Distance; i++ {
			final = Point{x: start.x - (i + 1), y: start.y}
			points.Update(final, startValue+i+1)
		}
	}

	return points, final, startValue + m.Distance
}

func wirePoints(start Point, w Wire) PointSet {
	points := NewPointSet()
	current := start
	currentValue := 0
	points.Update(current, currentValue)

	for _, move := range w {
		newPoints, final, finalValue := movePoints(current, currentValue, move)
		points.Merge(newPoints)
		current = final
		currentValue = finalValue
	}

	return points
}

func intersection(start Point, wires ...Wire) PointSet {
	intersection := NewPointSet()

	for i, wire := range wires {
		points := wirePoints(start, wire)

		if i == 0 {
			intersection = points
		} else {
			intersection = intersection.Intersection(points, func(v1, v2 int) int { return v1 + v2 })
		}
	}

	return intersection
}

// ClosestIntersectionDistance returns the Manhattan distance from the
// central port to the closest intersection of all the given wires.
func ClosestIntersectionDistance(wires ...Wire) int {
	start := Point{0, 0}
	intersection := intersection(start, wires...)
	intersection.Remove(start)

	var minDistance int
	i := 0
	for point := range intersection.Elements() {
		distance := ManhattanDistance(point, start)
		if i == 0 || distance < minDistance {
			minDistance = distance
		}
		i++
	}

	return minDistance
}

// FastestIntersectionSteps returns the fewest combined steps
// all the given wires must take to reach an intersection.
func FastestIntersectionSteps(wires ...Wire) int {
	start := Point{0, 0}
	intersection := intersection(start, wires...)
	intersection.Remove(start)

	var minSteps int
	i := 0
	for _, steps := range intersection.Elements() {
		if i == 0 || steps < minSteps {
			minSteps = steps
		}
		i++
	}

	return minSteps
}
