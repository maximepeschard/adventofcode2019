package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maximepeschard/adventofcode2019/03_wires/wire"
)

var usage = `Usage: wires <input-file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	text, err := ioutil.ReadFile(flag.Arg(0))
	check(err)
	lines := strings.Split(strings.TrimSpace(string(text)), "\n")
	wires := make([]wire.Wire, len(lines))
	for i, line := range lines {
		w, err := wire.Parse(line)
		check(err)
		wires[i] = w
	}

	minDistance := wire.ClosestIntersectionDistance(wires...)
	fmt.Println("Manhattan distance from the central port to the closest intersection:", minDistance)

	minSteps := wire.FastestIntersectionSteps(wires...)
	fmt.Println("Fewest combined steps the wires must take to reach an intersection:", minSteps)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
