package main

import "testing"

func TestRunComputer(t *testing.T) {
	// i := []int{1, 0, 0, 0, 99}
	// runComputer(i)

	// i = []int{2, 3, 0, 3, 99}
	// runComputer(i)

	// i = []int{2, 4, 4, 5, 99, 0}
	// runComputer(i)

	i := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	runComputer(i)
}
