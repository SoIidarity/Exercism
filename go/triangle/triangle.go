// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota
	Equ = iota
	Iso = iota
	Sca = iota
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	sideSum := a + b + c
	if math.IsNaN(sideSum) || math.IsInf(sideSum, 1) || math.IsInf(sideSum, -1) {
		return 0
	}
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	if sides[2] > sides[1]+sides[0] || sides[2] == 0 { //triangle inequality and no length in sides is not a triangle
		return 0
	}
	if sides[2] == sides[1] && sides[2] == sides[0] { // all sides are equal
		return 1
	}
	if sides[2] == sides[1] || sides[1] == sides[0] { // two sides are equal
		return 2
	}
	if sides[2] != sides[1] && sides[1] != sides[0] { // no sides are equal
		return 3
	}
	return 0
}
