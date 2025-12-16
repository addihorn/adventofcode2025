package main

import (
	"example/hello/src/aocutils"
	"fmt"
)

const inputFile = "input.txt"

var inputData []string

var tachyonBeams map[int]int

const (
	splitter = '^'
)

var split = 0

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	tachyonBeams = make(map[int]int)
	for i, p := range inputData[0] {
		if p == 'S' {
			tachyonBeams[i] = 1
		}
	}

	for _, row := range inputData[1:] {
		for i, s := range tachyonBeams {
			if row[i] == splitter {
				split++
				delete(tachyonBeams, i)
				tachyonBeams[i-1] += s
				tachyonBeams[i+1] += s
			}
		}
	}
}

/* Solve here */

func part1() {

	fmt.Printf("Solution for part 1: %d\n", split)
}

func part2() {
	timelines := 0
	for _, s := range tachyonBeams {
		timelines += s
	}
	fmt.Printf("Solution for part 1: %d\n", timelines)
}
