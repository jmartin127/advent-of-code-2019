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

func (p *point) string() string {
	return fmt.Sprintf("%d,%d has=%t", p.x, p.y, p.hasAstroid)
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

	// find the best location for the monitoring station
	bestLocation := findBestMonitoringStationLocation(am)

	// find all of the astroids that are in sight
	astroidsInSight := findOtherAstroidsInSight(bestLocation, am)
	fmt.Printf("Number of astroids in sight: %d\n", len(astroidsInSight))

	// determine which astroid is closest to >= 90 degrees from the monitoring station
	// refDegree := float64(90)
	refDegree := float64(-1)
	var nextAstroid *point
	var count int
	for true {
		nextAstroid, refDegree = nextAstroidClosestToRefDegree(refDegree, bestLocation, astroidsInSight)
		count++
		if nextAstroid == nil {
			break
		}
		nextAstroid.hasAstroid = false
		fmt.Printf("Count=%d, Next one %s. New Ref %f\n", count, nextAstroid.string(), refDegree)
	}
}

func nextAstroidClosestToRefDegree(refDegree float64, monitorStation *point, astroidsInSight []*point) (*point, float64) {
	minDist := float64(100000000)
	var nextAstroid *point
	var nextAstroidAngle float64
	for _, astroidInSight := range astroidsInSight {
		deg := findDegreesFromPoint(monitorStation, astroidInSight)
		dist := deg - refDegree
		if dist > 0 && dist < minDist {
			minDist = dist
			nextAstroid = astroidInSight
			nextAstroidAngle = deg
		}
	}
	return nextAstroid, nextAstroidAngle
}

func findDegreesFromPoint(a *point, b *point) float64 {
	x := b.x - a.x
	y := a.y - b.y

	radians := math.Atan2(float64(y), float64(x))
	deg := radians * (180 / math.Pi)

	deg = deg * -1  // increasing degrees as we go clockwise
	deg = deg - 270 // center on 90 degrees

	// ensure we aren't negative
	if deg < 0 {
		deg = deg + float64(360)
	}
	if deg < 0 {
		deg = deg + float64(360)
	}
	// fmt.Printf("From %s to %s is deg %f, radians %f\n", a.string(), b.string(), deg, radians)

	return deg
}

// Your job is to figure out which asteroid would be the best place to build a new monitoring station. A monitoring station can detect any asteroid to which it has direct line of sight
func findBestMonitoringStationLocation(am *astroidMap) *point {
	maxNum := 0
	var bestLocation *point
	for i, row := range am.data {
		for j := range row {
			otherAstroidsInSight := findOtherAstroidsInSight(am.data[i][j], am)
			num := len(otherAstroidsInSight)
			fmt.Printf("%d", num)
			if num > maxNum {
				maxNum = num
				bestLocation = am.data[i][j]
			}
		}
		fmt.Println()
	}

	fmt.Printf("Max: %d, from %d,%d\n", maxNum, bestLocation.x, bestLocation.y)

	return bestLocation
}

func findOtherAstroidsInSight(p *point, am *astroidMap) []*point {
	// we can only put the monitoring station where there is an astroid
	if !p.hasAstroid {
		return []*point{}
	}

	// fmt.Printf("Checking point at %d, %d\n", p.x, p.y)

	// compare to all other points
	otherAstroidsInSight := make([]*point, 0)
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
				otherAstroidsInSight = append(otherAstroidsInSight, otherPoint)
			}
		}
	}

	return otherAstroidsInSight
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
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day10_part2/input.txt")
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
