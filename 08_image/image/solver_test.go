package image

import (
	"reflect"
	"testing"
)

func TestCountDigits(t *testing.T) {
	cases := []struct {
		in   Layer
		want map[int]int
	}{
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
		},
		{
			[][]int{[]int{0, 1, 2}, []int{0, 0, 1}},
			map[int]int{0: 3, 1: 2, 2: 1},
		},
	}

	for _, c := range cases {
		got := CountDigits(c.in)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("CountDigits(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
