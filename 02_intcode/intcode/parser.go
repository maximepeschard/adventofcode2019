package intcode

import (
	"errors"
	"strconv"
	"strings"
)

// Parse returns a valid Intcode program parsed from a string.
func Parse(s string) ([]int, error) {
	codes := strings.Split(s, ",")
	program := make([]int, len(codes))
	for i, code := range codes {
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			return nil, errors.New("invalid program")
		}
		program[i] = codeInt
	}

	if len(program) < instructionSize && program[0] != opcodeHalt {
		return nil, errors.New("invalid program")
	}

	return program, nil
}
