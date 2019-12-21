package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var data = []int{}

func main() {
	input := 2 // The BOOST program will ask for a single input; run it in test mode by providing it the value 1.
	// It will perform a series of checks on each opcode, output any opcodes (and the associated parameter modes)
	// that seem to be functioning incorrectly, and finally output a BOOST keycode.

	readData()
	runComputer(input, data)
}

// The computer's available memory should be much larger than the initial program. Memory beyond the initial program
// starts with the value 0 and can be read or written like any other memory. (It is invalid to try to access memory at a
// negative address, though.)
func readData() {
	fmt.Println("Making very large array")
	data = make([]int, 100000000)
	fmt.Println("Done making array")

	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day9/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stringValues := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		stringValues = strings.Split(l, ",")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i, v := range stringValues {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		data[i] = intValue
	}
}

func runComputer(input int, original []int) {
	//fmt.Printf("Running for input %d\n", input)

	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	numToSkip := 0
	relativeModeBase := 0
	for i := 0; i < len(array); i += numToSkip {
		opcodeRaw := array[i]
		modes := parseModes(opcodeRaw)
		opcode := modes.opcode
		fmt.Printf("Opcode: %d, Mode: %+v\n", opcode, modes)
		if opcode == 99 {
			break
		} else if opcode == 1 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a + b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			fmt.Printf("Added %d + %d and set to correct place\n", a, b)
			numToSkip = 4
		} else if opcode == 2 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a * b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			numToSkip = 4
		} else if opcode == 3 {
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			// OLD: array[array[i+1]] = input
			setParameterForMode(modes.modeFirst, i+1, relativeModeBase, input, array)
			// TODO should this handle the relative base?
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			//fmt.Printf("Opcode 4 relative: %d\n", relativeModeBase)
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			fmt.Printf("OUTPUT: %d\n", a)
			input = a
			numToSkip = 2
		} else if opcode == 5 {
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			if a != 0 {
				i = b
				numToSkip = 0
			} else {
				numToSkip = 3
			}
		} else if opcode == 6 {
			// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			if a == 0 {
				i = b
				fmt.Printf("Setting i to %d\n", i)
				numToSkip = 0
			} else {
				numToSkip = 3
			}
		} else if opcode == 7 {
			// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			if a < b {
				setParameterForMode(modes.modeThird, i+3, relativeModeBase, 1, array)
			} else {
				setParameterForMode(modes.modeThird, i+3, relativeModeBase, 0, array)
			}
			numToSkip = 4
		} else if opcode == 8 {
			// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			if a == b {
				setParameterForMode(modes.modeThird, i+3, relativeModeBase, 1, array)
			} else {
				setParameterForMode(modes.modeThird, i+3, relativeModeBase, 0, array)
			}
			numToSkip = 4
		} else if opcode == 9 {
			// Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases (or decreases, if the value is negative) by the value of the parameter.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			relativeModeBase = relativeModeBase + a
			fmt.Printf("New relative base: %d\n", relativeModeBase)
			numToSkip = 2
		}
	}

	//fmt.Printf("Output: %d\n", input)
}

func getParameterForMode(mode, position, relativeModeBase int, a []int) int {
	var v int
	if mode == 0 { // position mode
		v = a[a[position]]
	} else if mode == 1 { // immediate mode
		v = a[position]
	} else if mode == 2 { // relative mode
		v = a[a[position]+relativeModeBase]
	} else {
		panic("This mode is not supported for getting parameters!")
	}

	return v
}

func setParameterForMode(mode, position, relativeModeBase, value int, a []int) {
	if mode == 0 { // position mode
		a[a[position]] = value
		fmt.Printf("Setting position %d to %d\n", a[position], value)
	} else if mode == 2 { // relative mode
		a[a[position]+relativeModeBase] = value
		fmt.Printf("Setting position %d to %d\n", a[position]+relativeModeBase, value)
	} else {
		panic("This mode is not supported for setting parameters!")
	}
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
