package main

import (
	"fmt"
)

const (
	north int = 1
	south int = 2
	west  int = 3
	east  int = 4
)

// 0: The repair droid hit a wall. Its position has not changed.
// 1: The repair droid has moved one step in the requested direction.
// 2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.
const (
	wall              int = 0
	moved             int = 1
	movedOxygenSystem int = 2
)

type point struct {
	x int
	y int
}

type cell struct {
	position            *point
	visited             bool
	oxygenSystem        bool
	directionsFromStart []int
}

type robot struct {
	position            *point
	directionsFromStart []int
}

var comp *computer

var allCells []*cell

func main() {
	comp = NewComputer("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day15/input.txt")
	shortestDistance := findShortestDistanceToOxygenSystem()
	fmt.Printf("Shortest distance: %d\n", shortestDistance)
}

// The remote control program executes the following steps in a loop forever:

// Accept a movement command via an input instruction.
// Send the movement command to the repair droid.
// Wait for the repair droid to finish the movement operation.
// Report on the status of the repair droid via an output instruction.
func findShortestDistanceToOxygenSystem() int {
	fmt.Println("Finding the shortest distance")

	r := &robot{directionsFromStart: []int{}, position: &point{}}

	queue := []*cell{}
	queue = append(queue, &cell{position: &point{x: 0, y: 0}, visited: true, directionsFromStart: []int{}})
	for len(queue) > 0 {
		// pop front node from the queue
		c := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// fmt.Printf("Popped cell from queue %+v\n", c)

		// if the popped node is the destination, return the distance
		if c.oxygenSystem {
			return len(c.directionsFromStart)
		}

		// go to the cell that was popped
		r.returnToStartLocation()
		r.followDirections(c.directionsFromStart)

		// for each of the 4 adjacent cells, enqueue each valid cell with a +1 distance and mark them visited
		queue = r.exploreAdjacentCellAndQueue(north, queue)
		queue = r.exploreAdjacentCellAndQueue(south, queue)
		queue = r.exploreAdjacentCellAndQueue(west, queue)
		queue = r.exploreAdjacentCellAndQueue(east, queue)

		fmt.Printf("New queue size: %d\n", len(queue))
	}

	panic("Never found the oxygen sensor")
}

func (r *robot) followDirections(directions []int) {
	for _, direction := range directions {
		r.moveForwardInDirection(direction)
	}
}

func (r *robot) moveForwardInDirection(direction int) int {
	_, output := comp.runComputer(direction)
	// fmt.Printf("\t\toutput %d\n", output)
	if output == wall {
		return output
	}

	// actually move the robot as well
	r.directionsFromStart = append(r.directionsFromStart, direction)

	// update the x/y coordinates of the robot
	r.updateXY(direction)

	// fmt.Printf("\trobot moved in direction %d\n", direction)
	return output
}

func (r *robot) exploreAdjacentCellAndQueue(direction int, queue []*cell) []*cell {
	cell := r.exploreAdjacentCell(direction)
	if cell != nil && !hasCellBeenVisited(cell) {
		// fmt.Printf("Adding cell to queue!  %+v\n", cell)
		queue = append(queue, cell)
		allCells = append(allCells, cell)
	}
	return queue
}

func hasCellBeenVisited(c *cell) bool {
	for _, o := range allCells {
		if c.position.x == o.position.x && c.position.y == o.position.y {
			return o.visited
		}
	}
	return false
}

// returns the adjacent cell if it is not a wall, otherwise returns nil
func (r *robot) exploreAdjacentCell(direction int) *cell {
	// fmt.Printf("\texploring cell in direction: %d\n", direction)
	// attempt to move the robot in the given direction
	output := r.moveForwardInDirection(direction)
	if output == wall {
		return nil
	}

	// check if we found the oxygen system
	var isOxygenSystem bool
	if output == movedOxygenSystem {
		isOxygenSystem = true
	}

	// record the directions to this cell from the starting point
	directionsFromStart := make([]int, len(r.directionsFromStart))
	copy(directionsFromStart, r.directionsFromStart)

	// create the cell
	cell := &cell{
		position:            &point{x: r.position.x, y: r.position.y}, // robot is currently in this cell
		visited:             true,
		oxygenSystem:        isOxygenSystem,
		directionsFromStart: directionsFromStart,
	}

	// move the robot back
	r.goBackOne()

	return cell
}

func (r *robot) returnToStartLocation() {
	atStart := r.goBackOne()
	for !atStart {
		atStart = r.goBackOne()
	}
}

// north=1, south=2, west=3, east=4
func (r *robot) goBackOne() (atStart bool) {
	if len(r.directionsFromStart) == 0 {
		return true
	}

	// remove the last direction
	priorDirection := r.directionsFromStart[len(r.directionsFromStart)-1]
	r.directionsFromStart = r.directionsFromStart[:len(r.directionsFromStart)-1]

	// determine which direction to move in to go back
	var newDirection int
	switch priorDirection {
	case 1:
		newDirection = 2
	case 2:
		newDirection = 1
	case 3:
		newDirection = 4
	case 4:
		newDirection = 3
	default:
		panic("Hmm... this ain't good")
	}

	// update the x/y
	r.updateXY(newDirection)

	// fmt.Printf("\trobot moved back to direction: %d\n", newDirection)

	ended, output := comp.runComputer(newDirection)
	if output != 1 && output != 2 {
		panic(fmt.Sprintf("Bad assumption on the output. Output was %d. Ended: %t", output, ended))
	}

	return false
}

func (r *robot) updateXY(direction int) {
	// update the x/y coordinates of the robot
	switch direction {
	case north:
		r.position.y = r.position.y - 1
	case south:
		r.position.y = r.position.y + 1
	case west:
		r.position.x = r.position.x - 1
	case east:
		r.position.x = r.position.x + 1
	}
}
