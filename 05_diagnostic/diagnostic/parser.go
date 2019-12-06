package diagnostic

import (
	"strconv"
	"strings"
)

// ParseInts returns a list of integers parsed from comma separated strings.
func ParseInts(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	ints := make([]int, len(parts))
	for i, part := range parts {
		partInt, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ints[i] = partInt
	}

	return ints, nil
}
