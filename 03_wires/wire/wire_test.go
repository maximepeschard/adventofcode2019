package wire

import (
	"log"
	"os"
	"testing"
)

type testCase struct {
	wiresStr  []string
	wires     []Wire
	wantDist  int
	wantSteps int
}

var cases []testCase

func TestClosestIntersectionDistance(t *testing.T) {
	for _, c := range cases {
		if got := ClosestIntersectionDistance(c.wires...); got != c.wantDist {
			t.Errorf("ClosestIntersectionDistance(%v) == %d, want %d", c.wiresStr, got, c.wantDist)
		}
	}
}

func TestFastestIntersectionSteps(t *testing.T) {
	for _, c := range cases {
		if got := FastestIntersectionSteps(c.wires...); got != c.wantSteps {
			t.Errorf("FastestIntersectionSteps(%v) == %d, want %d", c.wiresStr, got, c.wantSteps)
		}
	}
}

func TestMain(m *testing.M) {
	var wiresStr []string
	var w1, w2 Wire
	var err error

	wiresStr = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	w1, err = Parse(wiresStr[0])
	check(err)
	w2, err = Parse(wiresStr[1])
	check(err)
	cases = append(cases, testCase{wiresStr, []Wire{w1, w2}, 6, 30})

	wiresStr = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	w1, err = Parse(wiresStr[0])
	check(err)
	w2, err = Parse(wiresStr[1])
	check(err)
	cases = append(cases, testCase{wiresStr, []Wire{w1, w2}, 159, 610})

	wiresStr = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	w1, err = Parse(wiresStr[0])
	check(err)
	w2, err = Parse(wiresStr[1])
	check(err)
	cases = append(cases, testCase{wiresStr, []Wire{w1, w2}, 135, 410})

	os.Exit(m.Run())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
