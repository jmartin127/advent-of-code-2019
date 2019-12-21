package main

import (
	"fmt"
	"strconv"
)

var data = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 47, 72, 93, 110, 191, 272, 353, 434, 99999, 3, 9, 102, 3, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 1001, 9, 2, 9, 1002, 9, 2, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 101, 5, 9, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99}

// var data = []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
var permutations = [][]int{}

func main() {
	i := []int{0, 1, 2, 3, 4}
	heapPermutation(i, len(i))

	var maxThrusterSignal int
	for _, phaseSetting := range permutations {
		thrusterSignal := calculateThrusterSignal(phaseSetting)
		if thrusterSignal > maxThrusterSignal {
			maxThrusterSignal = thrusterSignal
		}
	}
	fmt.Printf("Max signal: %d\n", maxThrusterSignal)
}

func heapPermutation(a []int, size int) {
	if size == 1 {
		// copy the array, otherwise it will get modified
		newA := make([]int, len(a))
		copy(newA, a)
		permutations = append(permutations, newA)
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func calculateThrusterSignal(phaseSetting []int) int {
	// A
	output := runAmplifier(phaseSetting[0], 0)

	// B
	output = runAmplifier(phaseSetting[1], output)

	// C
	output = runAmplifier(phaseSetting[2], output)

	// D
	output = runAmplifier(phaseSetting[3], output)

	// E
	output = runAmplifier(phaseSetting[4], output)

	return output
}

func runAmplifier(phaseSetting, prevOutput int) int {
	// fmt.Printf("Running for phaseSetting %d and prevOutput %d\n", phaseSetting, prevOutput)

	// copy the array so we do not edit the original
	array := make([]int, len(data))
	copy(array, data)

	output := runComputer(phaseSetting, prevOutput, array)

	return output
}

func runComputer(firstInput, secondInput int, array []int) int {
	numToSkip := 0
	var output int
	inputUsed := false
	for i := 0; i < len(array); i += numToSkip {
		opcodeRaw := array[i]
		modes := parseModes(opcodeRaw)
		opcode := modes.opcode
		if opcode == 99 {
			break
		} else if opcode == 1 {
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			c := a + b
			array[array[i+3]] = c
			numToSkip = 4
		} else if opcode == 2 {
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			c := a * b
			array[array[i+3]] = c
			numToSkip = 4
		} else if opcode == 3 {
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			var input int
			if inputUsed {
				input = secondInput
			} else {
				input = firstInput
				inputUsed = true
			}

			array[array[i+1]] = input
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := array[array[i+1]]
			output = a
			numToSkip = 2
		} else if opcode == 5 {
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			if a != 0 {
				i = b
				numToSkip = 0
			} else {
				numToSkip = 3
			}
		} else if opcode == 6 {
			// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			if a == 0 {
				i = b
				numToSkip = 0
			} else {
				numToSkip = 3
			}
		} else if opcode == 7 {
			// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			if a < b {
				array[array[i+3]] = 1
			} else {
				array[array[i+3]] = 0
			}
			numToSkip = 4
		} else if opcode == 8 {
			// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := getParameterForMode(modes.modeFirst, i+1, array)
			b := getParameterForMode(modes.modeSecond, i+2, array)
			if a == b {
				array[array[i+3]] = 1
			} else {
				array[array[i+3]] = 0
			}
			numToSkip = 4
		}
	}

	// fmt.Printf("Output: %d\n", output)
	return output
}

func getParameterForMode(mode, position int, a []int) int {
	var v int
	if mode == 0 {
		v = a[a[position]]
	} else {
		v = a[position]
	}

	return v
}

type mode struct {
	opcode     int
	modeFirst  int
	modeSecond int
	modeThird  int
}

// ABCDE
//  1002

//  DE - two-digit opcode,      02 == opcode 2
//   C - mode of 1st parameter,  0 == position mode
//   B - mode of 2nd parameter,  1 == immediate mode
//   A - mode of 3rd parameter,  0 == position mode,
// 								   omitted due to being a leading zero
func parseModes(num int) *mode {
	s := strconv.Itoa(num)

	opcode := 0
	if len(s) > 1 {
		opcode, _ = strconv.Atoi(string(s[len(s)-2]) + string(s[len(s)-1]))
	} else {
		opcode, _ = strconv.Atoi(string(s[len(s)-1]))
	}

	modeFirst := 0
	if len(s) > 2 {
		modeFirst, _ = strconv.Atoi(string(s[len(s)-3]))
	}

	modeSecond := 0
	if len(s) > 3 {
		modeSecond, _ = strconv.Atoi(string(s[len(s)-4]))
	}

	modeThird := 0
	if len(s) > 4 {
		modeThird, _ = strconv.Atoi(string(s[len(s)-5]))
	}

	return &mode{
		opcode:     opcode,
		modeFirst:  modeFirst,
		modeSecond: modeSecond,
		modeThird:  modeThird,
	}
}
