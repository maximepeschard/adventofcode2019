package orbit

import (
	"fmt"
	"strings"
)

// An Orbit represents the link between a satellite and a center object.
type Orbit struct {
	Center    string
	Satellite string
}

func (o Orbit) String() string {
	return fmt.Sprintf("%s)%s", o.Center, o.Satellite)
}

// An Map represents a set of orbits.
type Map struct {
	Parent     *Map
	Object     string
	Satellites []*Map
}

// ParseOrbit returns a valid orbit parsed from a string.
func ParseOrbit(s string) (*Orbit, error) {
	parts := strings.Split(strings.TrimSpace(s), ")")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return nil, fmt.Errorf("invalid orbit: %s", s)
	}

	return &Orbit{Center: parts[0], Satellite: parts[1]}, nil
}

func parseMapRecursive(links map[string][]string, root string, parent *Map) *Map {
	children, prs := links[root]
	if !prs {
		return &Map{Parent: parent, Object: root}
	}

	satellites := make([]*Map, len(children))
	m := Map{Parent: parent, Object: root, Satellites: satellites}
	for i, child := range children {
		satellites[i] = parseMapRecursive(links, child, &m)
	}
	return &m
}

func parseOrbitLinks(orbits []*Orbit) map[string][]string {
	links := make(map[string][]string)
	for _, orbit := range orbits {
		if _, prs := links[orbit.Center]; !prs {
			links[orbit.Center] = []string{orbit.Satellite}
		} else {
			links[orbit.Center] = append(links[orbit.Center], orbit.Satellite)
		}
	}

	return links
}

// ParseMap returns an orbit map built from a slice of orbits.
func ParseMap(orbits []*Orbit) *Map {
	links := parseOrbitLinks(orbits)
	m := parseMapRecursive(links, "COM", nil)
	return m
}
