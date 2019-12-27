package image

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		width      int
		height     int
		data       string
		wantLayers []Layer
	}{
		{3, 2, "123456789012", []Layer{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			[][]int{[]int{7, 8, 9}, []int{0, 1, 2}},
		}},
	}

	for _, c := range cases {
		got, err := Parse(c.width, c.height, c.data)
		if err != nil {
			t.Errorf("Parse(%d, %d, %q) returns error: %s", c.width, c.height, c.data, err)
		} else if got.Width != c.width {
			t.Errorf("Parse(%d, %d, %q).Width == %d, want %d", c.width, c.height, c.data, got.Width, c.width)
		} else if got.Height != c.height {
			t.Errorf("Parse(%d, %d, %q).Height == %d, want %d", c.width, c.height, c.data, got.Height, c.height)
		} else if !reflect.DeepEqual(c.wantLayers, got.Layers) {
			t.Errorf("Parse(%d, %d, %q).Layers == %v, want %v", c.width, c.height, c.data, got.Layers, c.wantLayers)
		}
	}
}
