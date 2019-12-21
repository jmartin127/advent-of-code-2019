package main

import "testing"

func TestTwoAdjacentSame(t *testing.T) {
	same := twoAdjacentSame(111122)
	t.Fatalf("Answer: %t", same)
}
