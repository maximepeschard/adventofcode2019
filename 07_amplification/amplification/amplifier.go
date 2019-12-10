package amplification

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/maximepeschard/adventofcode2019/05_diagnostic/diagnostic"
)

// Amplify returns the output signal of an amplifier with the given
// program, input signal and phase.
func Amplify(program []int, input int, phase int) (int, error) {
	programCopy := make([]int, len(program))
	copy(programCopy, program)

	reader := strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, input))
	var writer strings.Builder
	err := diagnostic.Run(programCopy, reader, &writer)
	if err != nil {
		return input, err
	}
	programOutput := strings.Split(strings.TrimSpace(writer.String()), "\n")
	programOutputSize := len(programOutput)
	if programOutputSize < 1 {
		return input, errors.New("no output found when running amplifier program")
	}
	output, err := strconv.Atoi(programOutput[programOutputSize-1])
	if err != nil {
		return input, err
	}

	return output, nil
}
