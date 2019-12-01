package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/maximepeschard/adventofcode2019/01_rocket/rocket"
)

var usage = `Usage: rocket <input-file>`

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
	var masses []int
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		masses = append(masses, mass)
	}
	check(err)
	check(scanner.Err())

	fuelRequirements := 0
	totalFuelRequirements := 0
	for _, mass := range masses {
		fuelRequirements += rocket.ModuleFuel(mass)
		totalFuelRequirements += rocket.TotalModuleFuel(mass)
	}

	fmt.Println("Fuel requirements :", fuelRequirements)
	fmt.Println("Total fuel requirements :", totalFuelRequirements)
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
