package intcode

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	validCases := []struct {
		in   string
		want []int
	}{
		{"99", []int{99}},
		{"99,1,2", []int{99, 1, 2}},
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}},
	}
	for _, c := range validCases {
		parsed, err := Parse(c.in)
		if err != nil {
			t.Errorf("Parse(%q) returns error: %s", c.in, err)
		} else if !reflect.DeepEqual(parsed, c.want) {
			t.Errorf("Parse(%q) == %v, want %v", c.in, parsed, c.want)
		}
	}

	invalidCases := []string{
		"",
		"1",
		"1,2,3",
		"1,2,3,x,5",
		"1,2,3,4,",
	}
	for _, c := range invalidCases {
		_, err := Parse(c)
		if err == nil {
			t.Errorf("Parse(%q) returns nil error", c)
		}
	}
}
