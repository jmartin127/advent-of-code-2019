package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var data = []int{}

type computerOutput struct {
	// First, it will output a value indicating the color to paint the panel the robot is over:
	//   0 means to paint the panel black, and
	//   1 means to paint the panel white.
	color int

	// Second, it will output a value indicating the direction the robot should turn: 0 means it
	// should turn left 90 degrees, and 1 means it should turn right 90 degrees.
	direction int // up=0, right=1, down=2, left=3
}

type panel struct {
	x     int
	y     int
	color int
}

type robot struct {
	facing int
	panel  *panel
}

func (r *robot) turn(direction int) {
	if r.facing == 0 { // up
		if direction == 0 {
			r.facing = 3
		} else {
			r.facing = 1
		}
	} else if r.facing == 1 { // right
		if direction == 0 {
			r.facing = 0
		} else {
			r.facing = 2
		}
	} else if r.facing == 2 { // down
		if direction == 0 {
			r.facing = 1
		} else {
			r.facing = 3
		}
	} else if r.facing == 3 { // left
		if direction == 0 {
			r.facing = 2
		} else {
			r.facing = 0
		}
	}
}

func (r *robot) moveForward() (x, y int) {
	if r.facing == 0 { // up
		return r.panel.x, r.panel.y + 1
	} else if r.facing == 1 { // right
		return r.panel.x + 1, r.panel.y
	} else if r.facing == 2 { // down
		return r.panel.x, r.panel.y - 1
	} else if r.facing == 3 { // left
		return r.panel.x - 1, r.panel.y
	}

	return 0, 0
}

// 0=black
// 1=white
func main() {
	readData()

	fmt.Println("Starting computer")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	input := make(chan int, 4)
	input <- 0 // (All of the panels are currently black.)
	output := make(chan *computerOutput, 4)

	// initialize the robot and panels
	panels := make([]*panel, 0)
	robot := &robot{
		facing: 0, // start's facing up.
		panel:  &panel{x: 0, y: 0},
	}

	go runComputer(input, output, data, wg)

	var counter int
	for output := range output {
		counter++
		if counter < 20 {
			fmt.Printf("\tReceived output %+v\n", output)
		}

		// set the color
		currentPanel := findPanel(robot.panel.x, robot.panel.y, panels)
		if currentPanel == nil {
			currentPanel = &panel{x: robot.panel.x, y: robot.panel.y}
			fmt.Printf("Adding panel for current! %+v\n", currentPanel)
			panels = append(panels, currentPanel)
		}
		currentPanel.color = output.color
		if counter < 20 {
			fmt.Printf("\tSetting color to %d\n", output.color)
		}

		// turn the robot
		if counter < 20 {
			fmt.Printf("\tRobot was facing %d\n", robot.facing)
		}
		robot.turn(output.direction)
		if counter < 20 {
			fmt.Printf("\tturning %d\n", output.direction)
			fmt.Printf("\tRobot is now facing %d\n", robot.facing)
		}

		// move forward one panel
		newX, newY := robot.moveForward()
		if counter < 20 {
			fmt.Printf("New X: %d.  New Y: %d\n", newX, newY)
		}
		newPanel := findPanel(newX, newY, panels)
		if newPanel == nil {
			newPanel = &panel{x: newX, y: newY}
			fmt.Printf("Adding panel for new! %+v\n", newPanel)
			panels = append(panels, newPanel)
		}
		robot.panel = newPanel
		if counter < 20 {
			fmt.Printf("Setting panel for robot to %+v\n", newPanel)
		}

		// The Intcode program will serve as the brain of the robot. The program uses input instructions to access the
		// robot's camera: provide 0 if the robot is over a black panel or 1 if the robot is over a white panel.
		fmt.Printf("Sending %d. Num Panels: %d\n", robot.panel.color, len(panels))
		input <- robot.panel.color
	}

	fmt.Println("Waiting for computer to finish")
	wg.Wait()
	fmt.Println("Done waiting!")

}

func findPanel(x, y int, panels []*panel) *panel {
	for _, p := range panels {
		if p.x == x && p.y == y {
			return p
		}
	}
	return nil
}

// The computer's available memory should be much larger than the initial program. Memory beyond the initial program
// starts with the value 0 and can be read or written like any other memory. (It is invalid to try to access memory at a
// negative address, though.)
func readData() {
	fmt.Println("Making very large array")
	data = make([]int, 100000000)
	fmt.Println("Done making array")

	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day11/input.txt")
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

func runComputer(inputChannel chan int, outputChannel chan *computerOutput, original []int, wg *sync.WaitGroup) {
	//fmt.Printf("Running for input %d\n", input)
	defer wg.Done()

	// copy the array so we do not edit the original
	array := make([]int, len(original))
	copy(array, original)

	numToSkip := 0
	relativeModeBase := 0
	var output *computerOutput
	for i := 0; i < len(array); i += numToSkip {
		opcodeRaw := array[i]
		modes := parseModes(opcodeRaw)
		opcode := modes.opcode
		// fmt.Printf("Opcode: %d, Mode: %+v\n", opcode, modes)
		if opcode == 99 {
			fmt.Println("Received HALT signal!")
			break
		} else if opcode == 1 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a + b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			// fmt.Printf("Added %d + %d and set to correct place\n", a, b)
			numToSkip = 4
		} else if opcode == 2 {
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			b := getParameterForMode(modes.modeSecond, i+2, relativeModeBase, array)
			c := a * b
			setParameterForMode(modes.modeThird, i+3, relativeModeBase, c, array)
			numToSkip = 4
		} else if opcode == 3 {
			// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
			input := <-inputChannel
			setParameterForMode(modes.modeFirst, i+1, relativeModeBase, input, array)
			numToSkip = 2
		} else if opcode == 4 {
			// Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
			a := getParameterForMode(modes.modeFirst, i+1, relativeModeBase, array)
			if output == nil {
				output = &computerOutput{
					color: a,
				}
			} else {
				output.direction = a
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

	//fmt.Printf("Output: %d\n", input)
	close(inputChannel)
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
