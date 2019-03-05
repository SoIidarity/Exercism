package hamming

import (
	"errors"
)

//Distance returns the hamming distance between two strings.
//Function counts how many chars within the string differs
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Input lengths do not match")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
