package rocket

import "testing"

func TestModuleFuel(t *testing.T) {
	cases := []struct{ in, want int }{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		got := ModuleFuel(c.in)
		if got != c.want {
			t.Errorf("ModuleFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestTotalModuleFuel(t *testing.T) {
	cases := []struct{ in, want int }{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		got := TotalModuleFuel(c.in)
		if got != c.want {
			t.Errorf("TotalModuleFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
