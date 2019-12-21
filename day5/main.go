package main

import (
	"fmt"
	"strconv"
)

var data = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 65, 73, 225, 1101, 37, 7, 225, 1101, 42, 58, 225, 1102, 62, 44, 224, 101, -2728, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1, 69, 126, 224, 101, -92, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1102, 41, 84, 225, 1001, 22, 92, 224, 101, -150, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 80, 65, 225, 1101, 32, 13, 224, 101, -45, 224, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1101, 21, 18, 225, 1102, 5, 51, 225, 2, 17, 14, 224, 1001, 224, -2701, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 101, 68, 95, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 12, 22, 225, 102, 58, 173, 224, 1001, 224, -696, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 1002, 121, 62, 224, 1001, 224, -1302, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 344, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 374, 1001, 223, 1, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 101, 1, 223, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 101, 1, 223, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 524, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 554, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 569, 101, 1, 223, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 629, 101, 1, 223, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 659, 1001, 223, 1, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

// var data = []int{3, 0, 4, 0, 99} // outputs whatever it is given
// var data = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9} // outputs zero if input is zero, otherwise 1 (using position mode)
// var data = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1} // outputs zero if input is zero, otherwise 1 (using immediate mode)

func main() {
	input := 5 // The TEST diagnostic program will start by requesting from the user the ID of the system to test by running an input instruction - provide it 1, the ID for the ship's air conditioner unit.
	runComputer(input, data)
}

func runComputer(input int, original []int) {
	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	numToSkip := 0
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
			array[array[i+1]] = input
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := array[array[i+1]]
			input = a
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

	fmt.Printf("Output: %d\n", input)
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
