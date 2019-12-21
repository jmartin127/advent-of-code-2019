package main

import "fmt"

var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 1, 19, 5, 23, 2, 13, 23, 27, 1, 10, 27, 31, 2, 6, 31, 35, 1, 9, 35, 39, 2, 10, 39, 43, 1, 43, 9, 47, 1, 47, 9, 51, 2, 10, 51, 55, 1, 55, 9, 59, 1, 59, 5, 63, 1, 63, 6, 67, 2, 6, 67, 71, 2, 10, 71, 75, 1, 75, 5, 79, 1, 9, 79, 83, 2, 83, 10, 87, 1, 87, 6, 91, 1, 13, 91, 95, 2, 10, 95, 99, 1, 99, 6, 103, 2, 13, 103, 107, 1, 107, 2, 111, 1, 111, 9, 0, 99, 2, 14, 0, 0}

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			output := runComputer(i, j, input)
			if output == 19690720 {
				fmt.Printf("Noun: %d, Verb %d\n", i, j)
			}
		}
	}
}

func runComputer(noun, verb int, original []int) int {
	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	// initialize the memory in the computer
	array[1] = noun
	array[2] = verb

	for i := 0; i < len(array); i += 4 {
		opcode := array[i]
		if opcode == 99 {
			break
		} else if opcode == 1 {
			a := array[array[i+1]]
			b := array[array[i+2]]
			c := a + b
			array[array[i+3]] = c
		} else if opcode == 2 {
			a := array[array[i+1]]
			b := array[array[i+2]]
			c := a * b
			array[array[i+3]] = c
		}
	}

	// the output is stored in position 1
	return array[0]
}
