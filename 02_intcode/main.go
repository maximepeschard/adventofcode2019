package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maximepeschard/adventofcode2019/02_intcode/intcode"
)

var usage = `Usage: intcode <input-file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	programBytes, err := ioutil.ReadFile(flag.Arg(0))
	check(err)
	program, err := intcode.Parse(strings.TrimSpace(string(programBytes)))
	check(err)

	// Part 1
	programP1 := make([]int, len(program))
	copy(programP1, program)
	programP1[1] = 12 // noun
	programP1[2] = 2  // verb
	err = intcode.Run(programP1)
	check(err)
	fmt.Println("Value at position 0 after the program halts:", programP1[0])

	// Part 2
	noun, verb, err := intcode.Solve(program, 19690720)
	check(err)
	fmt.Println("100 * noun + verb:", 100*noun+verb)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("failed to close file %v: %s\n", f, err)
	}
}
