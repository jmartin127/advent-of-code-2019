package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	count := 0
	for i := 197487; i <= 673251; i++ {
		if twoAdjacentSame(i) && onlyIncrease(i) {
			fmt.Printf("Meets: %d\n", i)
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func twoAdjacentSame(num int) bool {
	s := strconv.Itoa(num)
	prev := ""
	numAdjacent := 0
	for _, r := range s {
		if prev != "" {
			//fmt.Printf("Comparing %s and %s\n", prev, string(r))
			if prev == string(r) {
				numAdjacent++
			} else {
				if numAdjacent == 1 {
					return true
				}
				numAdjacent = 0
			}
		}
		prev = string(r)
	}

	return numAdjacent == 1
}

func charIsPartOfLargerGroup(char string, num int) bool {
	invalidSeq := fmt.Sprintf("%s%s%s", char, char, char)
	s := strconv.Itoa(num)
	if strings.Contains(s, invalidSeq) {
		return true
	}

	return false
}

func onlyIncrease(num int) bool {
	s := strconv.Itoa(num)
	prev := ""
	prevInt := -1
	for _, r := range s {
		currInt, _ := strconv.Atoi(string(r))
		if prev != "" && currInt < prevInt {
			return false
		}
		prev = string(r)
		prevInt, _ = strconv.Atoi(prev)
	}
	return true
}
