package main

import (
	"fmt"

	"github.com/maximepeschard/adventofcode2019/04_password/password"
)

func main() {
	rangeStart := 372304
	rangeStop := 847060

	countPart1 := password.CountValidInRange(rangeStart, rangeStop, password.ValidPart1)
	fmt.Println("Number of valid passwords in the input range (part 1):", countPart1)

	countPart2 := password.CountValidInRange(rangeStart, rangeStop, password.ValidPart2)
	fmt.Println("Number of valid passwords in the input range (part 2):", countPart2)
}
