package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type chemicalAmount struct {
	chemical string
	amount   int
}

type equation struct {
	inputs []*chemicalAmount
	output *chemicalAmount
}

func (e *equation) multiply(multiplier int) {
	e.output.amount = e.output.amount * multiplier
	for i := range e.inputs {
		e.inputs[i].amount = e.inputs[i].amount * multiplier
	}
}

var totalOre int

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/Documents/advent-of-code/day15/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	equationsByOutputChemical := make(map[string]equation)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		outputChemicalAmount, inputsChemicalAmounts := parseLine(l)
		eq := equation{
			inputs: inputsChemicalAmounts,
			output: outputChemicalAmount,
		}
		equationsByOutputChemical[outputChemicalAmount.chemical] = eq
		// fmt.Printf("Output %+v, Inputs: %+v\n", outputChemicalAmount, inputsChemicalAmounts)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fuelEq := equationsByOutputChemical["FUEL"]
	fuelEq.multiply(1993284)
	fmt.Printf("Fuel Equation: %+v, Output %+v\n", fuelEq, fuelEq.output)

	surplus := make(map[string]int)
	resourcesRequred := make(map[string]*chemicalAmount)
	resourcesRequred[fuelEq.output.chemical] = fuelEq.output
	for i := 0; len(resourcesRequred) > 0; i++ {
		fmt.Printf("Resources requred: %+v\n", resourcesRequred)

		for _, requiredResource := range resourcesRequred {
			aquireResource(requiredResource, surplus, resourcesRequred, equationsByOutputChemical)
		}
	}

	fmt.Printf("TOTAL ORE: %d\n", totalOre)
	if totalOre > 1000000000000 {
		fmt.Println("You spent too much fuel!")
	}
}

func aquireResource(ca *chemicalAmount, surplus map[string]int, resourcesRequired map[string]*chemicalAmount, equations map[string]equation) {
	amountNeeded := ca.amount
	currentlyHave := surplus[ca.chemical]
	fmt.Printf("Chemical %s, Have: %d, Need: %d\n", ca.chemical, currentlyHave, amountNeeded)

	// if we already have enough in surplus, just use that and remove this as a required resource
	if currentlyHave >= amountNeeded {
		surplus[ca.chemical] = surplus[ca.chemical] - amountNeeded
		delete(resourcesRequired, ca.chemical)
		return
	}

	amountStillNeeded := amountNeeded - currentlyHave

	// Determine how many times to run the equation
	eq := equations[ca.chemical]
	m := int(math.Ceil(float64(amountStillNeeded) / float64(eq.output.amount)))

	// Determine how much was produced
	amountProduced := eq.output.amount * m

	// Determine how much surplus there now is
	newSurplus := amountProduced - amountStillNeeded
	surplus[ca.chemical] = newSurplus

	// Determine how much we now need in additional resources
	for _, input := range eq.inputs {
		newAmountOfResourceRequred := input.amount * m

		// for ORE, don't add it to the required resource list, just increment the counter
		if input.chemical == "ORE" {
			totalOre += newAmountOfResourceRequred
			continue
		}

		if _, ok := resourcesRequired[input.chemical]; ok {
			resourcesRequired[input.chemical].amount = resourcesRequired[input.chemical].amount + newAmountOfResourceRequred
		} else {
			newChemicalAmount := &chemicalAmount{chemical: input.chemical, amount: newAmountOfResourceRequred}
			resourcesRequired[input.chemical] = newChemicalAmount
		}
	}

	// Remove this resource from the list of required resources
	delete(resourcesRequired, ca.chemical)
}

// e.g., 11 RWLD, 3 BNKVZ, 4 PXTS, 3 XTRQC, 5 LSDX, 5 LMHL, 36 MGRTM => 4 ZCSB
func parseLine(l string) (*chemicalAmount, []*chemicalAmount) {
	s := strings.Split(l, "=>")
	inputs := s[0]
	output := s[1]

	outputChemicalAmount := parseChemicalAmount(output)

	inputsArray := strings.Split(inputs, ", ")
	inputsChemicalAmounts := make([]*chemicalAmount, 0)
	for _, input := range inputsArray {
		inputsChemicalAmounts = append(inputsChemicalAmounts, parseChemicalAmount(input))
	}

	return outputChemicalAmount, inputsChemicalAmounts
}

func parseChemicalAmount(s string) *chemicalAmount {
	s = strings.TrimSpace(s)
	a := strings.Split(s, " ")
	amount, err := strconv.Atoi(a[0])
	if err != nil {
		panic(err)
	}
	return &chemicalAmount{
		chemical: a[1],
		amount:   amount,
	}
}
