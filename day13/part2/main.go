package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var data = []int{}

var ball *computerOutput
var paddle *computerOutput

// The software draws tiles to the screen with output instructions: every three output instructions specify the x
// position (distance from the left), y position (distance from the top), and tile id. The tile id is interpreted
// as follows:
type computerOutput struct {
	// x position (distance from the left)
	x int

	// y position (distance from the top)
	y int

	// 0 is an empty tile. No game object appears in this tile.
	// 1 is a wall tile. Walls are indestructible barriers.
	// 2 is a block tile. Blocks can be broken by the ball.
	// 3 is a horizontal paddle tile. The paddle is indestructible.
	// 4 is a ball tile. The ball moves diagonally and bounces off objects.
	tileID int
}

func (c *computerOutput) tileChar() string {
	var s string
	if c.tileID == 0 {
		s = " "
	} else if c.tileID == 1 {
		s = "🔥"
	} else if c.tileID == 2 {
		s = "❎"
	} else if c.tileID == 3 {
		s = "🍗"
	} else if c.tileID == 4 {
		s = "🌕"
	}
	return s
}

func main() {
	readData()

	// The game didn't run because you didn't put in any quarters. Unfortunately, you did not bring any quarters.
	// Memory address 0 represents the number of quarters that have been inserted; set it to 2 to play for free.
	data[0] = 2

	// Start the computer
	wg := &sync.WaitGroup{}
	wg.Add(1)
	output := make(chan *computerOutput)
	go runComputer(output, data, wg)

	// Process each output as it comes
	outputs := make([]*computerOutput, 0)
	clearTerminal()
	for output := range output {
		if output.x == -1 && output.y == 0 {
			if output.tileID > 0 {
				// Check how close we are to finishing
				numLeft := numBlocksLeft(outputs)

				writeCharacterToPosition(25, 0, fmt.Sprintf("Score %d", output.tileID))
				writeCharacterToPosition(26, 0, fmt.Sprintf("Blocks remaining %3d", numLeft))

				//fmt.Printf("Number left: %d. Last Score: %d\n", numLeft, output.tileID)
				if numLeft == 0 {
					fmt.Printf("\nFinal score: %d\n", output.tileID)
					os.Exit(0)
				}
			}
		} else {
			// set the position of the ball
			if output.tileID == 4 {
				ball = output
			} else if output.tileID == 3 {
				paddle = output
			}

			// update the list of outputs
			if existingOutput := outputForPosition(output.x, output.y, outputs); existingOutput == nil {
				outputs = append(outputs, output)
			} else {
				existingOutput.tileID = output.tileID
			}

			// draw the output
			writeCharacterToPosition(output.y, output.x, output.tileChar())
		}
	}

	wg.Wait()
}

func numBlocksLeft(outputs []*computerOutput) int {
	var numLeft int
	for _, output := range outputs {
		if output.tileID == 2 {
			numLeft++
		}
	}
	return numLeft
}

func outputForPosition(x, y int, outputs []*computerOutput) *computerOutput {
	for _, output := range outputs {
		if output.x == x && output.y == y {
			return output
		}
	}
	return nil
}

func clearTerminal() {
	cmd := exec.Command("tput", "-S")
	cmd.Stdin = bytes.NewBufferString("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func writeCharacterToPosition(x, y int, char string) {
	cmd := exec.Command("tput", "-S")
	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("cup %d %d", x, y))
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf(char)
}

// The computer's available memory should be much larger than the initial program. Memory beyond the initial program
// starts with the value 0 and can be read or written like any other memory. (It is invalid to try to access memory at a
// negative address, though.)
func readData() {
	data = make([]int, 100000000)

	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day13/part2/input.txt")
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

func runComputer(outputChannel chan *computerOutput, original []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	numToSkip := 0
	relativeModeBase := 0
	var output *computerOutput
	var numOutputs int
	for i := 0; i < len(array); i += numToSkip {
		opcodeRaw := array[i]
		modes := parseModes(opcodeRaw)
		opcode := modes.opcode
		//fmt.Printf("Opcode: %d, Mode: %+v\n", opcode, modes)
		if opcode == 99 {
			break
		} else if opcode == 1 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a + b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			numToSkip = 4
		} else if opcode == 2 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a * b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			numToSkip = 4
		} else if opcode == 3 {
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.

			// Move the paddle where it should go
			// var input int
			// if paddle.x > ball.x {
			// 	input = -1
			// } else if paddle.x < ball.x {
			// 	input = 1
			// }

			// Randomly move the paddle
			// var input int
			// r := rand.Intn(3)
			// if r == 0 {
			// 	input = -1
			// } else if r == 1 {
			// 	input = 0
			// } else {
			// 	input = 1
			// }

			// Read user keystrokes
			reader := bufio.NewReader(os.Stdin)
			keystroke, _, _ := reader.ReadRune()
			var input int
			if keystroke == 'a' {
				input = -1
			} else if keystroke == 'd' {
				input = 1
			} else {
				input = 0
			}

			setParameterForMode(modes.modeFirst, i+1, relativeModeBase, input, array)
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			if numOutputs == 0 {
				output = &computerOutput{
					x: a,
				}
				numOutputs++
			} else if numOutputs == 1 {
				output.y = a
				numOutputs++
			} else if numOutputs == 2 {
				output.tileID = a
				numOutputs = 0
				outputChannel <- output
				output = nil
			}
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
			numToSkip = 2
		}
	}

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
	} else if mode == 2 { // relative mode
		a[a[position]+relativeModeBase] = value
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
