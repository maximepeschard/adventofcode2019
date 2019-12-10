package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maximepeschard/adventofcode2019/05_diagnostic/diagnostic"
	"github.com/maximepeschard/adventofcode2019/07_amplification/amplification"
	"github.com/maximepeschard/adventofcode2019/07_amplification/permutation"
)

var usage = `Usage: amplification <input-file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	programBytes, err := ioutil.ReadFile(flag.Arg(0))
	check(err)
	program, err := diagnostic.ParseInts(strings.TrimSpace(string(programBytes)))
	check(err)

	phasesCombinations := permutation.Ints([]int{0, 1, 2, 3, 4})
	maxSignal := 0
	for _, phases := range phasesCombinations {
		signal, err := amplification.AmplifyCircuit(program, phases)
		check(err)
		if signal > maxSignal {
			maxSignal = signal
		}
	}

	fmt.Println("Highest signal that can be sent to the thrusters:", maxSignal)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
