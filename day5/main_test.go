package main

import "testing"

func TestParseModes(t *testing.T) {
	m := parseModes(1002)
	t.Fatalf("P: %+v\n", m)
}
