package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel int
	for scanner.Scan() {
		l := scanner.Text()
		moduleMass, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		moduleFuel := fuelForModule(moduleMass)
		totalFuel += moduleFuel
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total fuel: %d\n", totalFuel)
}

func fuelForModule(mass int) int {
	var totalFuel int
	for lastFuel := mass; lastFuel > 0; {
		fuel := fuel(lastFuel)
		if fuel > 0 {
			totalFuel += fuel
		}
		lastFuel = fuel
	}
	return totalFuel
}

func fuel(num int) int {
	f := (num / 3) - 2
	fmt.Printf("F: %d\n", f)
	return f
}
