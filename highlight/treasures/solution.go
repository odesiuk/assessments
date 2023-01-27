package treasures

import (
	"errors"
	"strconv"
)

// FindLocation find Gold Treasures
func FindLocation(x, y int) (string, error) {
	if x < 1 || x > 100000 || y < 1 || y > 100000 {
		return "", errors.New("wrong params")
	}

	stepsToMax := x - 1
	maxWidth := y + stepsToMax

	// formula of Triangular number
	triangularNumber := (maxWidth * (maxWidth + 1)) / 2

	return strconv.Itoa(triangularNumber - stepsToMax), nil
}
