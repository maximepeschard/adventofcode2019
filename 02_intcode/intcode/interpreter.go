package intcode

import (
	"fmt"
)

const (
	opcodeAdd      = 1
	opcodeMultiply = 2
	opcodeHalt     = 99
)

const instructionSize = 4

// Run executes a valid Intcode program.
func Run(program []int) error {
	iptr := 0
	for program[iptr] != opcodeHalt {
		switch program[iptr] {
		case opcodeAdd:
			program[program[iptr+3]] = program[program[iptr+1]] + program[program[iptr+2]]
		case opcodeMultiply:
			program[program[iptr+3]] = program[program[iptr+1]] * program[program[iptr+2]]
		default:
			return fmt.Errorf("invalid opcode: %d", program[iptr])
		}

		iptr += instructionSize
	}

	return nil
}
