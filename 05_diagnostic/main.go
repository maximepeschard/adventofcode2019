package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maximepeschard/adventofcode2019/05_diagnostic/diagnostic"
)

var usage = `Usage: diagnostic <input-file>`

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

	err = diagnostic.Run(program, os.Stdin, os.Stdout)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
