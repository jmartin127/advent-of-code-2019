package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type computer struct {
	data             []int
	numToSkip        int
	relativeModeBase int
}

func NewComputer(filepath string) *computer {
	data := readData(filepath)

	return &computer{
		data: data,
	}
}

// The computer's available memory should be much larger than the initial program. Memory beyond the initial program
// starts with the value 0 and can be read or written like any other memory. (It is invalid to try to access memory at a
// negative address, though.)
func readData(filepath string) []int {
	fmt.Println("Making very large array")
	data := make([]int, 100000000)
	fmt.Println("Done making array")

	file, err := os.Open(filepath)
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

	return data
}

func (comp *computer) runComputer(input int) (ended bool, output int) {

	for i := 0; i < len(comp.data); i += comp.numToSkip {
		opcodeRaw := comp.data[i]
		modes := parseModes(opcodeRaw)
		opcode := modes.opcode
		//fmt.Printf("Opcode: %d, Mode: %+v\n", opcode, modes)
		if opcode == 99 {
			break
		} else if opcode == 1 {
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			c := a + b
			comp.setParameterForMode(modes.modeThird, i+3, c)
			comp.numToSkip = 4
		} else if opcode == 2 {
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			c := a * b
			comp.setParameterForMode(modes.modeThird, i+3, c)
			comp.numToSkip = 4
		} else if opcode == 3 {
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			comp.setParameterForMode(modes.modeFirst, i+1, input)
			comp.numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			return false, a
			comp.numToSkip = 2
		} else if opcode == 5 {
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			if a != 0 {
				i = b
				comp.numToSkip = 0
			} else {
				comp.numToSkip = 3
			}
		} else if opcode == 6 {
			// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			if a == 0 {
				i = b
				comp.numToSkip = 0
			} else {
				comp.numToSkip = 3
			}
		} else if opcode == 7 {
			// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			if a < b {
				comp.setParameterForMode(modes.modeThird, i+3, 1)
			} else {
				comp.setParameterForMode(modes.modeThird, i+3, 0)
			}
			comp.numToSkip = 4
		} else if opcode == 8 {
			// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			b := comp.getParameterForMode(modes.modeSecond, i+2)
			if a == b {
				comp.setParameterForMode(modes.modeThird, i+3, 1)
			} else {
				comp.setParameterForMode(modes.modeThird, i+3, 0)
			}
			comp.numToSkip = 4
		} else if opcode == 9 {
			// Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases (or decreases, if the value is negative) by the value of the parameter.
			a := comp.getParameterForMode(modes.modeFirst, i+1)
			comp.relativeModeBase = comp.relativeModeBase + a
			comp.numToSkip = 2
		}
	}

	return true, 0
}

func (c *computer) getParameterForMode(mode, position int) int {
	var v int
	if mode == 0 { // position mode
		v = c.data[c.data[position]]
	} else if mode == 1 { // immediate mode
		v = c.data[position]
	} else if mode == 2 { // relative mode
		v = c.data[c.data[position]+c.relativeModeBase]
	} else {
		panic("This mode is not supported for getting parameters!")
	}

	return v
}

func (c *computer) setParameterForMode(mode, position, value int) {
	if mode == 0 { // position mode
		c.data[c.data[position]] = value
	} else if mode == 2 { // relative mode
		c.data[c.data[position]+c.relativeModeBase] = value
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
