package wire

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	validCases := []struct {
		in   string
		want Wire
	}{
		{"R8,U5,L5,D3", []Move{Move{Right, 8}, Move{Up, 5}, Move{Left, 5}, Move{Down, 3}}},
		{"U7,R6,D4,L4", []Move{Move{Up, 7}, Move{Right, 6}, Move{Down, 4}, Move{Left, 4}}},
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
		"U",
		"Z1",
		"Uz",
		"R8,U,L5,D3",
		"R8,U5,Lz,D3",
	}
	for _, c := range invalidCases {
		_, err := Parse(c)
		if err == nil {
			t.Errorf("Parse(%q) returns nil error", c)
		}
	}
}
