package main

import (
	"example/hello/src/aocutils"
	"fmt"
)

const inputFile = "input.txt"

var inputData []string

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	_ = inputData
}

/* Solve here */

func part1() {
	_ = inputData
	sum := 0
	for _, data := range inputData {
		sum += findJoltage(data, 2)
	}

	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {
	_ = inputData
	sum := 0
	for _, data := range inputData {
		s := findJoltage(data, 12)
		//fmt.Println(s)
		sum += s
	}

	fmt.Printf("Solution for part 2: %d\n", sum)
}

func findJoltage(battery string, length int) int {

	max := battery[:length]

	for _, r := range battery[length:] {
		nm := aocutils.CString2Int(max)
		for i := range max {
			candidate := fmt.Sprintf("%s%s%s", max[:i], max[i+1:], string(r))
			nm = aocutils.Max([]int{nm, aocutils.CString2Int(candidate)})
		}

		newMax := aocutils.Max([]int{
			aocutils.CString2Int(max),
			nm,
		})

		//fmt.Println(newMax)

		max = fmt.Sprint(newMax)

	}
	return aocutils.CString2Int(max)

}
