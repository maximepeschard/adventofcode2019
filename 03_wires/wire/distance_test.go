package wire

import "testing"

func TestManhattanDistance(t *testing.T) {
	cases := []struct {
		in1, in2 Point
		want     int
	}{
		{Point{0, 0}, Point{0, 4}, 4},
		{Point{0, 0}, Point{0, -4}, 4},
		{Point{0, 0}, Point{4, 0}, 4},
		{Point{0, 0}, Point{-4, 0}, 4},
		{Point{0, 0}, Point{4, 5}, 9},
		{Point{0, 0}, Point{-4, 5}, 9},
		{Point{3, -1}, Point{-4, 5}, 13},
		{Point{2, 7}, Point{2, 7}, 0},
	}

	for _, c := range cases {
		if got := ManhattanDistance(c.in1, c.in2); got != c.want {
			t.Errorf("ManhattanDistance(%v, %v) == %d, want %d", c.in1, c.in2, got, c.want)
		}
	}
}
