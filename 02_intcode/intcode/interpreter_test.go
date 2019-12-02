package intcode

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct{ initial, final []int }{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, c := range cases {
		program := c.initial
		Run(program)
		if !reflect.DeepEqual(program, c.final) {
			t.Errorf("Run(%v) -> %v, want %v", c.initial, program, c.final)
		}
	}
}
