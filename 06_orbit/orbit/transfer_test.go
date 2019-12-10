package orbit

import "testing"

func TestMininumTransfers(t *testing.T) {
	cases := []struct {
		inMap  *Map
		inFrom string
		inTo   string
		want   int
	}{
		{testMapPart2, "YOU", "SAN", 4},
	}
	for i, c := range cases {
		got, err := MininumTransfers(c.inMap, c.inFrom, c.inTo)
		if err != nil {
			t.Errorf("case %d: MininumTransfers returns error: %s", i, err)
		} else if got != c.want {
			t.Errorf("case %d: MininumTransfers --> %d, want %d", i, got, c.want)
		}
	}
}
