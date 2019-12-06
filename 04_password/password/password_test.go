package password

import "testing"

func TestValidPart1(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}

	for _, c := range cases {
		if got := ValidPart1(c.in); got != c.want {
			t.Errorf("Valid(%s) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestValidPart2(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}

	for _, c := range cases {
		if got := ValidPart2(c.in); got != c.want {
			t.Errorf("Valid(%s) == %t, want %t", c.in, got, c.want)
		}
	}
}
