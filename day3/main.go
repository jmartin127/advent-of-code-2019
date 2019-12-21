package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x          int
	y          int
	wire1      bool
	wire2      bool
	wire1Steps int
	wire2Steps int
}

type segment struct {
	direction string
	length    int
}

var matrixSize = 25000
var middle = matrixSize / 2

func main() {
	// first dimension is "x", second is "y"
	// initialize the matrix to be arbitrarily large
	matrix := make([][]*point, matrixSize)
	for i := range matrix {
		matrix[i] = make([]*point, matrixSize)
		for j := range matrix[i] {
			matrix[i][j] = &point{x: i, y: j}
		}
	}
	fmt.Println("Done initializing")

	// Add the wires to the matrix
	wire1Input, wire2Input := loadWireInputs("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day3/input-final.txt")
	fmt.Println("Adding wire 1")
	addWireToMatrix(matrix, wire1Input, true)
	fmt.Println("Adding wire 2")
	addWireToMatrix(matrix, wire2Input, false)

	// Determine where they cross
	minDistance := 100000000
	for i := range matrix {
		for j := range matrix[i] {
			p := matrix[i][j]
			if p.wire1 && p.wire2 {
				fmt.Printf("Wires cross at: %d, %d\n", i, j)
				xDistance := middle - i
				if xDistance < 0 {
					xDistance = i - middle
				}
				yDistance := middle - j
				if yDistance < 0 {
					yDistance = j - middle
				}
				//distance := xDistance + yDistance
				//fmt.Printf("\tdistance: %d, x %d, y %d\n", distance, xDistance, yDistance)

				// part 2
				xTotalSteps := p.wire1Steps
				yTotalSteps := p.wire2Steps
				totalSteps := xTotalSteps + yTotalSteps
				distance := totalSteps
				fmt.Printf("\tdistance: %d, x %d, y %d\n", distance, xTotalSteps, yTotalSteps)

				if distance < minDistance && distance != 0 {
					minDistance = distance
				}
			}
		}
	}
	fmt.Printf("Min distance: %d\n", minDistance)
}

func addWireToMatrix(matrix [][]*point, wireInput []string, isWire1 bool) {
	startPoint := &point{x: middle, y: middle}
	totalSteps := 0
	for _, input := range wireInput {
		seg := parseSegment(input)
		startPoint = populateSegment(matrix, startPoint, isWire1, seg, totalSteps)
		fmt.Printf("New start point: %+v. Total steps: %d\n", startPoint, totalSteps)
		totalSteps = totalSteps + seg.length
	}
}

func populateSegment(matrix [][]*point, start *point, isWire1 bool, seg *segment, totalSteps int) *point {
	if seg.direction == "U" {
		for i := 0; i < seg.length; i++ {
			// fmt.Printf("X %d, Y %d\n", start.x, start.y+i)
			p := matrix[start.x][start.y+i]
			updatePoint(p, isWire1, totalSteps+i)
		}
		return matrix[start.x][start.y+seg.length]
	} else if seg.direction == "D" {
		for i := 0; i < seg.length; i++ {
			// fmt.Printf("X %d, Y %d\n", start.x, start.y-i)
			p := matrix[start.x][start.y-i]
			updatePoint(p, isWire1, totalSteps+i)
		}
		return matrix[start.x][start.y-seg.length]
	} else if seg.direction == "L" {
		for i := 0; i < seg.length; i++ {
			// fmt.Printf("X %d, Y %d\n", start.x-i, start.y)
			p := matrix[start.x-i][start.y]
			updatePoint(p, isWire1, totalSteps+i)
		}
		return matrix[start.x-seg.length][start.y]
	} else if seg.direction == "R" {
		for i := 0; i < seg.length; i++ {
			// fmt.Printf("X %d, Y %d\n", start.x+i, start.y)
			p := matrix[start.x+i][start.y]
			updatePoint(p, isWire1, totalSteps+i)
		}
		return matrix[start.x+seg.length][start.y]
	}

	return nil
}

func updatePoint(p *point, isWire1 bool, totalSteps int) {
	if isWire1 {
		p.wire1 = true
		if p.wire1Steps == 0 {
			p.wire1Steps = totalSteps
		}
	} else {
		p.wire2 = true
		if p.wire2Steps == 0 {
			p.wire2Steps = totalSteps
		}
	}
}

func parseSegment(input string) *segment {
	if strings.HasPrefix(input, "U") {
		return &segment{direction: "U", length: parseSegmentLength(input)}
	} else if strings.HasPrefix(input, "D") {
		return &segment{direction: "D", length: parseSegmentLength(input)}
	} else if strings.HasPrefix(input, "L") {
		return &segment{direction: "L", length: parseSegmentLength(input)}
	} else if strings.HasPrefix(input, "R") {
		return &segment{direction: "R", length: parseSegmentLength(input)}
	}

	panic("shouldn't get here")
}

func parseSegmentLength(input string) int {
	s := input[1:]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func loadWireInputs(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wire1Input []string
	var wire2Input []string
	if scanner.Scan() {
		line := scanner.Text()
		wire1Input = readLine(line)
	}
	if scanner.Scan() {
		line := scanner.Text()
		wire2Input = readLine(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return wire1Input, wire2Input
}

func readLine(line string) []string {
	return strings.Split(line, ",")
}
