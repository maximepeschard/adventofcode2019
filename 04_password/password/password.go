package password

import "strconv"

const validPasswordLength = 6

// ValidPart1 reports whether the password meets the criteria of Part 1.
func ValidPart1(pwd string) bool {
	pwdLength := len(pwd)

	if pwdLength != validPasswordLength {
		return false
	}

	adjacent := false
	prevDigit := pwd[0]
	for index := 1; index < pwdLength; index++ {
		digit := pwd[index]

		if digit < prevDigit {
			return false
		} else if digit == prevDigit {
			adjacent = true
		}

		prevDigit = digit
	}

	return adjacent
}

// ValidPart2 reports whether the password meets the criteria of Part 2.
func ValidPart2(pwd string) bool {
	pwdLength := len(pwd)

	if pwdLength != validPasswordLength {
		return false
	}

	currentAdjacent := 1
	adjacent := false
	prevDigit := pwd[0]
	for index := 1; index < pwdLength; index++ {
		digit := pwd[index]

		if digit < prevDigit {
			return false
		} else if digit == prevDigit {
			currentAdjacent++
		} else {
			adjacent = adjacent || currentAdjacent == 2
			currentAdjacent = 1
		}

		prevDigit = digit
	}
	adjacent = adjacent || currentAdjacent == 2

	return adjacent
}

// CountValidInRange returns the number of passwords in the range satisfying a validator function.
func CountValidInRange(start int, stop int, validator func(s string) bool) int {
	count := 0

	for candidate := start; candidate <= stop; candidate++ {
		if validator(strconv.Itoa(candidate)) {
			count++
		}
	}

	return count
}
