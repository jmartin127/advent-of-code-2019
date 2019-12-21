package main

import (
	"fmt"
	"strconv"
	"sync"
)

// REAL
var data = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 47, 72, 93, 110, 191, 272, 353, 434, 99999, 3, 9, 102, 3, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 1001, 9, 2, 9, 1002, 9, 2, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 101, 5, 9, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99}

// TEST
// var data = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
// var data = []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

var permutations = [][]int{}
var finalAnswer = 0

func main() {
	i := []int{5, 6, 7, 8, 9}
	heapPermutation(i, len(i))

	var maxThrusterSignal int
	for _, phaseSetting := range permutations {
		fmt.Printf("Phase Setting %v\n", phaseSetting)
		thrusterSignal := calculateThrusterSignal(phaseSetting)
		if thrusterSignal > maxThrusterSignal {
			maxThrusterSignal = thrusterSignal
		}
	}
	fmt.Printf("Max signal: %d\n", maxThrusterSignal)

	// TESTING a single phase setting input
	// setting := []int{9, 8, 7, 6, 5}
	// answer := calculateThrusterSignal(setting)
	// fmt.Printf("ANSWER: %d\n", answer)
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
	ac := make(chan int)
	bc := make(chan int)
	cc := make(chan int)
	dc := make(chan int)
	ec := make(chan int)

	fmt.Println("Starting amplifiers")
	wg := &sync.WaitGroup{}
	wg.Add(5)

	go runComputer("E", ec, ac, data, wg)
	ec <- phaseSetting[4]

	go runComputer("D", dc, ec, data, wg)
	dc <- phaseSetting[3]

	go runComputer("C", cc, dc, data, wg)
	cc <- phaseSetting[2]

	go runComputer("B", bc, cc, data, wg)
	bc <- phaseSetting[1]

	go runComputer("A", ac, bc, data, wg)
	ac <- phaseSetting[0]
	ac <- 0

	fmt.Println("Finished starting amplifiers")

	fmt.Println("Waiting for amplifiers to finish")
	wg.Wait()
	fmt.Println("Done waiting!")

	return finalAnswer
}

func runComputer(label string, inputChannel, outputChannel chan int, original []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// handle the terminal case, when we try to send output on a closed channel for amplifier E, that means we are done
	defer func() {
		if r := recover(); r != nil {
			if label == "E" {
				fmt.Printf("Amplifier %s. Expected for Amplifier E. Recovered from: %v\n", label, r)
				close(inputChannel)
			} else {
				fmt.Printf("Amplifier %s. Unexpected panic!!!: %v\n", label, r)
			}
		}
	}()

	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	numToSkip := 0
	var output int
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
			fmt.Printf("Amplifier %s, Receiving Input\n", label)
			input := <-inputChannel
			fmt.Printf("Amplifier %s, Received Input %d\n", label, input)
			array[array[i+1]] = input
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := array[array[i+1]]
			output = a
			fmt.Printf("Amplifier %s, Sending Output %d\n", label, output)
			if label == "E" {
				finalAnswer = output
			}
			outputChannel <- output
			fmt.Printf("Amplifier %s, Really sent Output %d\n", label, output)
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

	fmt.Printf("Amplifier %s, Closing input channel.\n", label)
	close(inputChannel)
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
