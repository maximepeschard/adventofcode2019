package diagnostic

import (
	"reflect"
	"testing"
)

func TestFillParamModes(t *testing.T) {
	cases := []struct {
		inParamModes []int
		inNbParams   int
		want         []int
	}{
		{[]int{}, 3, []int{parameterModePosition, parameterModePosition, parameterModePosition}},
		{[]int{parameterModeImmediate}, 2, []int{parameterModeImmediate, parameterModePosition}},
		{[]int{parameterModeImmediate}, 1, []int{parameterModeImmediate}},
	}

	for _, c := range cases {
		if got := fillParamModes(c.inParamModes, c.inNbParams); !reflect.DeepEqual(got, c.want) {
			t.Errorf("fillParamModes(%v, %d) == %v, want %v", c.inParamModes, c.inNbParams, got, c.want)
		}
	}
}

func TestValidOpcode(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{opcodeAdd, true},
		{opcodeMultiply, true},
		{opcodeInput, true},
		{opcodeOutput, true},
		{opcodeJumpIfTrue, true},
		{opcodeJumpIfFalse, true},
		{opcodeLessThan, true},
		{opcodeEquals, true},
		{opcodeHalt, true},
		{999, false},
	}

	for _, c := range cases {
		if got := validOpcode(c.in); got != c.want {
			t.Errorf("validOpcode(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestInstruction(t *testing.T) {
	cases := []struct {
		in             int
		wantOpcode     int
		wantParamModes []int
	}{
		{3, 3, nil},
		{99, 99, nil},
		{1002, 2, []int{parameterModePosition, parameterModeImmediate}},
		{1101, 1, []int{parameterModeImmediate, parameterModeImmediate}},
	}

	for _, c := range cases {
		gotOpcode, gotParamModes, err := instruction(c.in)
		if err != nil {
			t.Error(err)
		} else if gotOpcode != c.wantOpcode {
			t.Errorf("instruction(%d) returns opcode %d, want %d", c.in, gotOpcode, c.wantOpcode)
		} else if !reflect.DeepEqual(gotParamModes, c.wantParamModes) {
			t.Errorf("instruction(%d) returns param modes %v, want %v", c.in, gotParamModes, c.wantParamModes)
		}
	}
}

func TestParamValue(t *testing.T) {
	cases := []struct {
		inProgram []int
		inMode    int
		inValue   int
		want      int
	}{
		{[]int{0, 10, 20, 30, 40}, parameterModeImmediate, 8, 8},
		{[]int{0, 10, 20, 30, 40}, parameterModeImmediate, 2, 2},
		{[]int{0, 10, 20, 30, 40}, parameterModePosition, 2, 20},
	}

	for _, c := range cases {
		if got := paramValue(c.inProgram, c.inMode, c.inValue); got != c.want {
			t.Errorf("paramValue(%v, %d, %d) == %d, want %d", c.inProgram, c.inMode, c.inValue, got, c.want)
		}
	}
}
