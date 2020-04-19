// Package hamming provides a way to calculate the hamming distance of two strands of DNA.
package hamming

import (
	"errors"
)

// Distance function calculates hamming distance between two strands of DNA.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("strands are not equal length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
