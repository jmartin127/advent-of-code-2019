package main

import (
	"fmt"
	"testing"
)

func TestFuelForModule(t *testing.T) {
	f := fuelForModule(1969)
	fmt.Printf("Fuel: %d\n", f)
}
