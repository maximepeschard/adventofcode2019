package orbit

import "testing"

func TestParse(t *testing.T) {
	validCases := []struct {
		in   string
		want Orbit
	}{
		{"A)B", Orbit{Center: "A", Satellite: "B"}},
	}
	for _, c := range validCases {
		got, err := ParseOrbit(c.in)
		if err != nil {
			t.Errorf("Parse(%q) returns error: %s", c.in, err)
		} else if *got != c.want {
			t.Errorf("Parse(%q) == %v, want %v", c.in, *got, c.want)
		}
	}

	invalidCases := []string{"", "A", "AB", "A)", ")B"}
	for _, c := range invalidCases {
		_, err := ParseOrbit(c)
		if err == nil {
			t.Errorf("Parse(%q) returns nil error", c)
		}
	}

}
