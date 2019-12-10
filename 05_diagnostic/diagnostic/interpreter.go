package diagnostic

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	opcodeAdd         = 1
	opcodeMultiply    = 2
	opcodeInput       = 3
	opcodeOutput      = 4
	opcodeJumpIfTrue  = 5
	opcodeJumpIfFalse = 6
	opcodeLessThan    = 7
	opcodeEquals      = 8
	opcodeHalt        = 99
)

const (
	parameterModePosition  = 0
	parameterModeImmediate = 1
)

// Run executes an Intcode program.
func Run(program []int, reader io.Reader, writer io.Writer) error {
	iptr := 0

	for {
		opcode, paramModes, err := instruction(program[iptr])
		if err != nil {
			return err
		} else if opcode == opcodeHalt {
			break
		}

		switch opcode {
		case opcodeAdd, opcodeMultiply:
			nbParams := 3
			paramModes = fillParamModes(paramModes, nbParams)
			target := program[iptr+3]
			operand1 := paramValue(program, paramModes[0], program[iptr+1])
			operand2 := paramValue(program, paramModes[1], program[iptr+2])
			if opcode == opcodeAdd {
				// fmt.Printf("(debug) writing %d + %d to position %d\n", operand1, operand2, target)
				program[target] = operand1 + operand2
			} else {
				// fmt.Printf("(debug) writing %d * %d to position %d\n", operand1, operand2, target)
				program[target] = operand1 * operand2
			}
			iptr += nbParams + 1
		case opcodeInput:
			nbParams := 1
			target := program[iptr+1]
			value, err := readInput(reader)
			if err != nil {
				return err
			}
			program[target] = value
			iptr += nbParams + 1
		case opcodeOutput:
			nbParams := 1
			paramModes = fillParamModes(paramModes, nbParams)
			value := paramValue(program, paramModes[0], program[iptr+1])
			fmt.Fprintf(writer, "%d\n", value)
			iptr += nbParams + 1
		case opcodeJumpIfTrue, opcodeJumpIfFalse:
			nbParams := 2
			paramModes = fillParamModes(paramModes, nbParams)
			testValue := paramValue(program, paramModes[0], program[iptr+1])
			newIptr := paramValue(program, paramModes[1], program[iptr+2])
			var test bool
			if opcode == opcodeJumpIfTrue {
				test = testValue != 0
			} else {
				test = testValue == 0
			}
			if test {
				iptr = newIptr
			} else {
				iptr += nbParams + 1
			}
		case opcodeLessThan, opcodeEquals:
			nbParams := 3
			paramModes = fillParamModes(paramModes, nbParams)
			target := program[iptr+3]
			operand1 := paramValue(program, paramModes[0], program[iptr+1])
			operand2 := paramValue(program, paramModes[1], program[iptr+2])
			var test bool
			if opcode == opcodeLessThan {
				test = operand1 < operand2
			} else {
				test = operand1 == operand2
			}
			if test {
				program[target] = 1
			} else {
				program[target] = 0
			}
			iptr += nbParams + 1
		}
	}

	return nil
}

func instruction(value int) (int, []int, error) {
	opcode := value % 100
	if !validOpcode(opcode) {
		return 0, nil, fmt.Errorf("invalid opcode: %d", opcode)
	}

	var paramModes []int
	paramModeValues := value / 100
	for paramModeValues > 0 {
		mode := paramModeValues % 10
		if mode != parameterModePosition && mode != parameterModeImmediate {
			return 0, nil, fmt.Errorf("invalid parameter mode: %d", mode)
		}
		paramModes = append(paramModes, mode)
		paramModeValues = paramModeValues / 10
	}

	return opcode, paramModes, nil
}

func validOpcode(code int) bool {
	validOpcodes := []int{
		opcodeAdd,
		opcodeMultiply,
		opcodeInput,
		opcodeOutput,
		opcodeJumpIfTrue,
		opcodeJumpIfFalse,
		opcodeLessThan,
		opcodeEquals,
		opcodeHalt,
	}
	for _, c := range validOpcodes {
		if code == c {
			return true
		}
	}

	return false
}

func paramValue(program []int, mode int, value int) int {
	if mode == parameterModeImmediate {
		return value
	}

	return program[value]
}

func readInput(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	if ok := scanner.Scan(); !ok {
		return 0, scanner.Err()
	}

	scanned := scanner.Text()
	value, err := strconv.Atoi(scanned)
	if err != nil {
		return 0, nil
	}

	return value, nil
}

func fillParamModes(paramModes []int, nbParams int) []int {
	length := len(paramModes)
	for index := 0; index < nbParams-length; index++ {
		paramModes = append(paramModes, parameterModePosition)
	}

	return paramModes
}
