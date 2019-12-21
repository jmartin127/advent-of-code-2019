package main

import (
	"fmt"
	"math"
)

type moon struct {
	x         int
	y         int
	z         int
	xVelocity int
	yVelocity int
	zVelocity int
}

type system struct {
	moons []*moon
}

func (m *moon) print() {
	fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x= %d, y= %d, z= %d>\n", m.x, m.y, m.z, m.xVelocity, m.yVelocity, m.zVelocity)
}

func (m *moon) hash() string {
	return fmt.Sprintf("%d%d", m.z, m.zVelocity)
}

func (m *moon) potentialEnergy() int {
	return absoluteValue(m.x) + absoluteValue(m.y) + absoluteValue(m.z)
}

func (m *moon) kineticEnergy() int {
	return absoluteValue(m.xVelocity) + absoluteValue(m.yVelocity) + absoluteValue(m.zVelocity)
}

func (m *moon) totalEnergy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

func absoluteValue(v int) int {
	return int(float64(math.Abs(float64(v))))
}

func (s *system) printMoons() {
	for _, moon := range s.moons {
		moon.print()
	}
	fmt.Println()
}

func (s *system) hash() string {
	return fmt.Sprintf("%s%s%s%s", s.moons[0].hash(), s.moons[1].hash(), s.moons[2].hash(), s.moons[3].hash())
}

func (s *system) totalEnergy() int {
	totalEnergy := 0
	for _, moon := range s.moons {
		totalEnergy += moon.totalEnergy()
	}
	return totalEnergy
}

// Final Answer:
// 186028 (x unit of time to orbit)
// 231614 (y unit of time to orbit)
// 102356 (z unit of time to orbit)
// LCM = 551,272,644,867,044

// Input:
// <x=5, y=4, z=4>
// <x=-11, y=-11, z=-3>
// <x=0, y=7, z=0>
// <x=-13, y=2, z=10>

func main() {

	moons := []*moon{}
	moons = append(moons, &moon{x: 5, y: 4, z: 4})
	moons = append(moons, &moon{x: -11, y: -11, z: -3})
	moons = append(moons, &moon{x: 0, y: 7, z: 0})
	moons = append(moons, &moon{x: -13, y: 2, z: 10})

	// moons = append(moons, &moon{x: -1, y: 0, z: 2})
	// moons = append(moons, &moon{x: 2, y: -10, z: -7})
	// moons = append(moons, &moon{x: 4, y: -8, z: 8})
	// moons = append(moons, &moon{x: 3, y: 5, z: -1})

	// moons = append(moons, &moon{x: -8, y: -10, z: 0})
	// moons = append(moons, &moon{x: 5, y: 5, z: 10})
	// moons = append(moons, &moon{x: 2, y: -7, z: 3})
	// moons = append(moons, &moon{x: 9, y: -8, z: -3})
	sys := &system{moons: moons}

	priorHashes := make(map[string]struct{})
	priorHashes[sys.hash()] = struct{}{}
	for i := 0; true; i++ {
		if i%100000 == 0 {
			fmt.Printf("After %d steps:\n", i)
			// sys.printMoons()
			// fmt.Printf("Total energy: %d\n", sys.totalEnergy())
			p := float64(i) / 4686774924 * 100
			fmt.Printf("Percent: %f\n", p)
		}
		applyGravity(sys.moons)
		applyVelocity(sys.moons)
		hash := sys.hash()
		if _, ok := priorHashes[hash]; ok {
			fmt.Printf("Reached prior state after %d steps\n", i+1)
			break
		}
		priorHashes[hash] = struct{}{}
	}

}

func applyVelocity(moons []*moon) {
	for _, moon := range moons {
		moon.x += moon.xVelocity
		moon.y += moon.yVelocity
		moon.z += moon.zVelocity
	}
}

// To apply gravity, consider every pair of moons.
func applyGravity(moons []*moon) {

	for i := 0; i < len(moons); i++ {
		for j := 0; j < len(moons); j++ {
			if i < j { // don't compare against itself
				applyGravityToPair(moons[i], moons[j])
			}
		}
	}

}

// On each axis (x, y, and z), the velocity of each moon changes
// by exactly +1 or -1 to pull the moons together. For example, if Ganymede has an x position of 3, and Callisto
// has a x position of 5, then Ganymede's x velocity changes by +1 (because 5 > 3) and Callisto's x velocity
// changes by -1 (because 3 < 5). However, if the positions on a given axis are the same, the velocity on that
// axis does not change for that pair of moons.
func applyGravityToPair(a *moon, b *moon) {
	if a.x < b.x {
		a.xVelocity++
		b.xVelocity--
	} else if a.x > b.x {
		a.xVelocity--
		b.xVelocity++
	}

	if a.y < b.y {
		a.yVelocity++
		b.yVelocity--
	} else if a.y > b.y {
		a.yVelocity--
		b.yVelocity++
	}

	if a.z < b.z {
		a.zVelocity++
		b.zVelocity--
	} else if a.z > b.z {
		a.zVelocity--
		b.zVelocity++
	}
}
