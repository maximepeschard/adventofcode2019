package amplification

// AmplifyCircuit returns the output signal of an amplification
// circuit with the given program, input signal and phases.
func AmplifyCircuit(program []int, phases []int) (int, error) {
	var err error
	signal := 0

	for _, phase := range phases {
		signal, err = Amplify(program, signal, phase)
		if err != nil {
			return signal, err
		}
	}

	return signal, nil
}
