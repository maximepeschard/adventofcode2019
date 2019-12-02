package intcode

import "fmt"

// Solve returns the input noun and verb that cause the program to produce the output.
func Solve(program []int, output int) (int, int, error) {
	testProgram := make([]int, len(program))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(testProgram, program)
			testProgram[1] = noun
			testProgram[2] = verb
			err := Run(testProgram)
			if err != nil {
				return 0, 0, err
			}

			if testProgram[0] == output {
				return noun, verb, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("no noun and verb found for output %d", output)
}
