package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/maximepeschard/adventofcode2019/06_orbit/orbit"
)

var usage = `Usage: orbit <input-file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	check(err)
	defer closeFile(f)

	scanner := bufio.NewScanner(f)
	var orbits []*orbit.Orbit
	for scanner.Scan() {
		orbit, err := orbit.ParseOrbit(scanner.Text())
		check(err)
		orbits = append(orbits, orbit)
	}
	check(scanner.Err())

	orbitMap := orbit.ParseMap(orbits)
	totalOrbits := orbit.ChecksumTotalOrbits(orbitMap)
	fmt.Println("Total number of direct and indirect orbits:", totalOrbits)
	minTransfers, err := orbit.MininumTransfers(orbitMap, "YOU", "SAN")
	check(err)
	fmt.Println("Minimum number of orbital transfers required:", minTransfers)
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("failed to close file %v: %s\n", f, err)
	}
}
