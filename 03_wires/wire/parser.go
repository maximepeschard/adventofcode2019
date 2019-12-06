package wire

import (
	"fmt"
	"strconv"
	"strings"
)

// Direction represents a direction in the grid.
type Direction string

// Valid directions.
const (
	Up    Direction = "U"
	Left  Direction = "L"
	Down  Direction = "D"
	Right Direction = "R"
)

// A Move is defined by a direction and a distance.
type Move struct {
	Direction Direction
	Distance  int
}

// A Wire is represented as a sequence of moves.
type Wire []Move

// Parse returns a valid wire parsed from a string.
func Parse(s string) (Wire, error) {
	movesStr := strings.Split(s, ",")
	moves := make([]Move, len(movesStr))
	for i, move := range movesStr {
		if len(move) < 2 {
			return nil, fmt.Errorf("invalid move %s", move)
		}

		dir := Direction(move[0])
		if dir != Up && dir != Left && dir != Down && dir != Right {
			return nil, fmt.Errorf("invalid direction '%s' in move %s", dir, move)
		}

		dist, err := strconv.Atoi(move[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance '%s' in move %s", move[1:], move)
		}

		moves[i] = Move{dir, dist}
	}

	return moves, nil
}
