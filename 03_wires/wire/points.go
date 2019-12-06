package wire

// Point represents a position in the grid.
type Point struct {
	x int
	y int
}

// PointSet represents a set of points with associated values.
type PointSet struct {
	points map[Point]int
}

// NewPointSet creates an empty point set.
func NewPointSet() PointSet {
	return PointSet{points: make(map[Point]int)}
}

// Elements returns all the points and associated values of a point set.
func (ps PointSet) Elements() map[Point]int {
	return ps.points
}

// Update adds a point and its value to the set if absent or if the value is
// smaller than the current value for this point. The return value indicates
// if the point has effectively been added.
func (ps *PointSet) Update(p Point, value int) bool {
	if v, prs := ps.points[p]; !prs || value < v {
		ps.points[p] = value
		return true
	}

	return false
}

// Remove deletes a point and its value from the set if present.
// The return value indicates if the point has effectively been removed.
func (ps *PointSet) Remove(p Point) bool {
	if _, prs := ps.points[p]; prs {
		delete(ps.points, p)
		return true
	}

	return false
}

// Merge adds all points from another set to the set.
// The return value indicates the number of points effectively added.
func (ps *PointSet) Merge(other PointSet) int {
	merged := 0
	for point, value := range other.points {
		if updated := ps.Update(point, value); updated {
			merged++
		}
	}

	return merged
}

// Intersection returns the intersection of the set with another set, with
// new values computed from the values of both sets using the given function.
func (ps PointSet) Intersection(other PointSet, op func(v1, v2 int) int) PointSet {
	intersection := NewPointSet()

	smallest := ps.points
	largest := other.points
	if len(other.points) < len(ps.points) {
		smallest = other.points
		largest = ps.points
	}

	for point, value := range smallest {
		if valueOther, prs := largest[point]; prs {
			intersection.Update(point, op(value, valueOther))
		}
	}

	return intersection
}
