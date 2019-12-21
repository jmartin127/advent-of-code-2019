package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	hasAstroid bool
	x          int
	y          int
}

type astroidMap struct {
	data [][]*point
}

func (am *astroidMap) print() {
	for _, row := range am.data {
		for _, v := range row {
			p := "."
			if v.hasAstroid {
				p = "#"
			}
			fmt.Printf(p)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	am := readMap()
	am.print()

	// answer := hasClearLineOfSight(am.data[2][4], am.data[0][4], am)
	// fmt.Printf("Answer :%t\n", answer)

	// answer := determineNumAstroidsInSight(am.data[2][4], am)
	// fmt.Printf("Answer :%d\n", answer)

	maxNum := 0
	maxX := 0
	maxY := 0
	for i, row := range am.data {
		for j := range row {
			num := determineNumAstroidsInSight(am.data[i][j], am)
			fmt.Printf("%d", num)
			if num > maxNum {
				maxNum = num
				maxX = j
				maxY = i
			}
		}
		fmt.Println()
	}

	fmt.Printf("Max: %d, from %d,%d\n", maxNum, maxX, maxY)
}

func determineNumAstroidsInSight(p *point, am *astroidMap) int {
	// we can only put the monitoring station where there is an astroid
	if !p.hasAstroid {
		return 0
	}

	// fmt.Printf("Checking point at %d, %d\n", p.x, p.y)

	// compare to all other points
	var numInSight int
	for i, row := range am.data {
		for j := range row {
			// if the other point doesn't have an astroid, then don't check it
			otherPoint := am.data[i][j]
			if !otherPoint.hasAstroid {
				continue
			}

			// don't check against itself
			if (otherPoint.x == p.x) && (otherPoint.y == p.y) {
				continue
			}

			// fmt.Printf("Now checking %d,%d\n", otherPoint.x, otherPoint.y)
			// check if there are any other astroids blocking the line of sight
			if hasClearLineOfSight(p, otherPoint, am) {
				// fmt.Printf("\thas clear line of site to %d, %d\n", otherPoint.x, otherPoint.y)
				numInSight++
			}
		}
	}

	return numInSight
}

// checks if there is a clear line of sight between points a/b
func hasClearLineOfSight(a *point, b *point, am *astroidMap) bool {
	for i, row := range am.data {
		for j := range row {
			// if the other point doesn't have an astroid, then don't check it
			c := am.data[i][j]
			if !c.hasAstroid {
				continue
			}

			// don't check against themselves
			if c == a {
				continue
			}
			if c == b {
				continue
			}

			// fmt.Printf("Comparing %d,%d\n", c.x, c.y)
			// fmt.Printf("\t\tA: %d,%d\n", a.x, a.y)
			// fmt.Printf("\t\tB: %d,%d\n", b.x, b.y)
			// fmt.Printf("\t\tC: %d,%d\n", c.x, c.y)
			// fmt.Printf("\tDistance AB %f\n", distAB)
			// fmt.Printf("\tDistance AC %f\n", distAC)
			// fmt.Printf("\tCollinear? %t\n", pointsAreCollinear(a, b, c))

			if pointsAreCollinear(a, b, c) && pointCIsBetweenAB(a, b, c) {
				// fmt.Printf("\t\t\tReturning false\n")
				return false
			}
		}
	}

	// fmt.Printf("\t\t\tReturning true\n")
	return true
}

func distanceBetweenPoints(a *point, b *point) float64 {
	ax := float64(a.x)
	bx := float64(b.x)
	ay := float64(a.y)
	by := float64(b.y)
	return math.Sqrt(((ax - bx) * (ax - bx)) + ((ay - by) * (ay - by)))
}

func pointsAreCollinear(a *point, b *point, c *point) bool {
	slopeAB := slope(a, b)
	slopeBC := slope(b, c)
	slopeAC := slope(a, c)
	// fmt.Printf("\t\tSlope AB: %f\n", slopeAB)
	// fmt.Printf("\t\tSlope BC: %f\n", slopeBC)
	// fmt.Printf("\t\tSlope AC: %f\n", slopeAC)
	return (slopeAB == slopeBC) && (slopeBC == slopeAC)
}

func pointCIsBetweenAB(a *point, b *point, c *point) bool {
	dotproduct := (c.x-a.x)*(b.x-a.x) + (c.y-a.y)*(b.y-a.y)
	if dotproduct < 0 {
		return false
	}

	squaredlengthba := (b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y)
	if dotproduct > squaredlengthba {
		return false
	}

	return true
}

func slope(a *point, b *point) float64 {
	bot := float64(b.x - a.x)
	if bot == 0 {
		return -10000000 // no slope
	}
	m := float64(b.y-a.y) / bot
	return m
}

func readMap() *astroidMap {
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lineNum int
	data := make([][]*point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		a := strings.Split(l, "")
		row := make([]*point, 0)
		for i, v := range a {
			hasAstroid := false
			if v == "#" {
				hasAstroid = true
			}
			p := &point{
				hasAstroid: hasAstroid,
				x:          i,
				y:          lineNum,
			}
			row = append(row, p)
		}
		lineNum++
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &astroidMap{data: data}
}
