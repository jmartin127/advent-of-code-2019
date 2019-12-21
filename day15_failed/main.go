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

	fuelEq := equationsByOutputChemical["CXFTF"]
	fmt.Printf("Fuel Equation: %+v\n", fuelEq)
	allMaterials := determineRawMaterialsForChemicalAmount(fuelEq.output, equationsByOutputChemical)

	totalOre := 0
	for chemical, amount := range allMaterials {
		fmt.Printf("Raw material %d %s\n", amount, chemical)
		eq := equationsByOutputChemical[chemical]
		goalAmount := amount
		outputAmount := eq.output.amount

		var multiplier int
		if amount < 0 {
			goalAmount = goalAmount * -1
			test := int(math.Floor(float64(goalAmount) / float64(outputAmount)))
			fmt.Printf("Num batches to REMOVE!!! %d %s. Goal %d, Output %d\n", test, chemical, goalAmount, outputAmount)
		} else {
			multiplier = int(math.Ceil(float64(goalAmount) / float64(outputAmount)))
		}

		totalOre += (multiplier * eq.inputs[0].amount)
		fmt.Printf("\tmultiplier %d input %d, adding %d ore to the final answer\n", multiplier, eq.inputs[0].amount, (multiplier * eq.inputs[0].amount))
	}

	fmt.Printf("Answer: %d\n", totalOre)
}

func determineRawMaterialsForChemicalAmount(ca *chemicalAmount, equations map[string]equation) map[string]int {
	fmt.Printf("Determining amount of ore to make: %s, %d\n", ca.chemical, ca.amount)

	// load the equation
	equation := equations[ca.chemical]

	// exit if this is a raw material
	if len(equation.inputs) == 1 && equation.inputs[0].chemical == "ORE" {
		m := make(map[string]int)
		m[equation.output.chemical] = ca.amount
		return m
	}

	// determine the multiplier (how many times we need to run the equation)
	m := int(math.Ceil(float64(ca.amount) / float64(equation.output.amount)))

	// initialize the output
	allMaterials := make(map[string]int)

	// check if there is extra
	extra := (equation.output.amount * m) - ca.amount
	fmt.Printf("EXTRA chemical: %s %d\n", ca.chemical, extra)

	// remove the extra
	if _, ok := allMaterials[ca.chemical]; ok {
		allMaterials[ca.chemical] = allMaterials[ca.chemical] - extra
	} else {
		allMaterials[ca.chemical] = -1 * extra
	}

	// apply the multiplier to all of the inputs
	for _, inputChemicalAmount := range equation.inputs {
		multipliedChemicalAmount := &chemicalAmount{
			chemical: inputChemicalAmount.chemical,
			amount:   inputChemicalAmount.amount * m,
		}
		rawMaterialsNeeded := determineRawMaterialsForChemicalAmount(multipliedChemicalAmount, equations)
		fmt.Printf("__Need %v for %d %s\n", rawMaterialsNeeded, multipliedChemicalAmount.amount, multipliedChemicalAmount.chemical)

		for k, v := range rawMaterialsNeeded {
			fmt.Printf("Adding %s %d\n", k, v)
			if _, ok := allMaterials[k]; ok {
				allMaterials[k] = allMaterials[k] + v
			} else {
				allMaterials[k] = v
			}
		}
	}

	fmt.Printf("Need %v for %d %s\n", allMaterials, ca.amount, ca.chemical)
	return allMaterials
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
