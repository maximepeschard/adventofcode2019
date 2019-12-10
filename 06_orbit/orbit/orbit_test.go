package orbit

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const testDataPart1 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

const testDataPart2 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

var (
	testMapPart1 *Map
	testMapPart2 *Map
)

func TestMain(m *testing.M) {
	var orbits []*Orbit

	orbits = nil
	for _, s := range strings.Split(testDataPart1, "\n") {
		orbit, err := ParseOrbit(s)
		if err != nil {
			fmt.Printf("failed to initialize test resource: %s\n", err)
			os.Exit(1)
		}
		orbits = append(orbits, orbit)
	}
	testMapPart1 = ParseMap(orbits)

	orbits = nil
	for _, s := range strings.Split(testDataPart2, "\n") {
		orbit, err := ParseOrbit(s)
		if err != nil {
			fmt.Printf("failed to initialize test resource: %s\n", err)
			os.Exit(1)
		}
		orbits = append(orbits, orbit)
	}
	testMapPart2 = ParseMap(orbits)

	os.Exit(m.Run())
}
