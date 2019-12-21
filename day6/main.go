package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	label  string
	parent *node
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	nodes := make(map[string]*node)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chunks := strings.Split(line, ")")
		orbited := chunks[0]
		orbiter := chunks[1]

		// add the nodes to the map if we have never seen them
		if _, ok := nodes[orbited]; !ok {
			nodes[orbited] = &node{label: orbited}
		}
		if _, ok := nodes[orbiter]; !ok {
			nodes[orbiter] = &node{label: orbiter}
		}

		// connect the nodes
		orbitedNode := nodes[orbited]
		orbiterNode := nodes[orbiter]
		orbiterNode.parent = orbitedNode

		fmt.Printf("Orbited: %s, Orbiter: %s\n", orbited, orbiter)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Now, loop through the nodes and count the parents
	totalCount := 0
	for _, node := range nodes {
		nodeCount := countOrbits(node, nodes)
		totalCount += nodeCount
	}
	fmt.Printf("Final Count: %d\n", totalCount)

	// Part 2
	youParents := getParentLabels(nodes["YOU"], nodes)
	santaParents := getParentLabels(nodes["SAN"], nodes)

	fmt.Printf("You parents: %v\n", youParents)
	fmt.Printf("Santa parents: %v\n", santaParents)

	// Find the number of labels unique to each list and add them
	youUnique := numUniqueElements(youParents, santaParents)
	santaUnique := numUniqueElements(santaParents, youParents)
	fmt.Printf("Result: %d and %d\n", youUnique, santaUnique)
}

func countOrbits(n *node, nodes map[string]*node) int {
	var count int
	for current := n; current != nil; count++ {
		current = current.parent
	}
	return count - 1
}

func getParentLabels(n *node, nodes map[string]*node) []string {
	parentLabels := make([]string, 0)
	var count int
	for current := n; current != nil; count++ {
		if current.parent != nil {
			parentLabels = append(parentLabels, current.parent.label)
		}
		current = current.parent
	}
	return parentLabels
}

func numUniqueElements(a []string, b []string) int {
	uniqueCount := 0
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
			}
		}
		if !found {
			uniqueCount++
		}
	}
	return uniqueCount
}
