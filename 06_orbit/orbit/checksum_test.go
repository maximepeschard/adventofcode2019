package orbit

import "testing"

func TestChecksumTotalOrbits(t *testing.T) {
	cases := []struct {
		in   *Map
		want int
	}{
		{testMapPart1, 42},
	}
	for i, c := range cases {
		if got := ChecksumTotalOrbits(c.in); got != c.want {
			t.Errorf("case %d: ChecksumTotalOrbits --> %d, want %d", i, got, c.want)
		}
	}
}
