package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/maximepeschard/adventofcode2019/08_image/image"
)

var usage = `Usage: image <width> <height> <input-file>`

func main() {
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()
	if flag.NArg() != 3 {
		flag.Usage()
		os.Exit(1)
	}

	imageWidth, err := strconv.Atoi(flag.Arg(0))
	check(err)
	imageHeight, err := strconv.Atoi(flag.Arg(1))
	check(err)
	imageData, err := ioutil.ReadFile(flag.Arg(2))
	check(err)

	img, err := image.Parse(imageWidth, imageHeight, string(imageData))
	check(err)

	layerIndex := image.FindLayerWithFewest(img, 0)
	counts := image.CountDigits(img.Layers[layerIndex])
	fmt.Println("Number of 1 digits multiplied by the number of 2 digits:", counts[1]*counts[2])

	fmt.Println("Decoded message:")
	decoded := image.Decode(img)
	image.PrintDecoded(decoded)
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
