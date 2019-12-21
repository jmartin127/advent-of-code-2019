package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wide = 25
var tall = 6

type layer struct {
	data [][]int
}

func (l *layer) numDigit(digit int) int {
	var numDigit int
	for _, row := range l.data {
		for _, v := range row {
			if v == digit {
				numDigit++
			}
		}
	}
	return numDigit
}

func (l *layer) print() {
	for _, row := range l.data {
		for _, v := range row {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func main() {
	layers := readLayers()

	var minZerosLayer *layer
	minZeros := 10000000000000
	for _, layer := range layers {
		layer.print()
		fmt.Println()
		numZeros := layer.numDigit(0)
		if numZeros < minZeros {
			minZeros = numZeros
			minZerosLayer = layer
		}
	}

	answer := minZerosLayer.numDigit(1) * minZerosLayer.numDigit(2)

	fmt.Printf("Answer: %d\n", answer)
}

func readLayers() []*layer {
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		l := scanner.Text()
		input = strings.Split(l, "")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	intArray := make([]int, 0)
	for _, v := range input {
		i, _ := strconv.Atoi(v)
		intArray = append(intArray, i)
	}

	layers := make([]*layer, 0)
	l := &layer{data: make([][]int, 0)}
	for i := 0; i < len(intArray); i += wide {
		row := intArray[i : i+wide]
		if len(l.data) < tall {
			l.data = append(l.data, row)
		} else {
			layers = append(layers, l)
			l = &layer{data: make([][]int, 0)}
			l.data = append(l.data, row)
		}
	}
	layers = append(layers, l)

	return layers
}
