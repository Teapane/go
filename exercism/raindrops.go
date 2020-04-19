// Package raindrops converts numbers to raindrop sounds
package raindrops

import (
	"fmt"
)

// Convert use modulo operation to convert ints to strings
func Convert(input int) string {
	empty := ""

	if input%3 == 0 {
		empty += "Pling"
	}
	if input%5 == 0 {
		empty += "Plang"
	}
	if input%7 == 0 {
		empty += "Plong"
	}
	if len(empty) > 0 {
		return empty
	}
	return fmt.Sprintf("%d", input)
}
