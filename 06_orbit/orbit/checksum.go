package orbit

// ChecksumTotalOrbits returns the total number of direct and indirect orbits in a map.
func ChecksumTotalOrbits(m *Map) int {
	return countTotalOrbits(m, 0)
}

func countTotalOrbits(m *Map, acc int) int {
	if m.Satellites == nil {
		return acc
	}

	out := acc
	for _, sat := range m.Satellites {
		out += countTotalOrbits(sat, acc+1)
	}

	return out
}
