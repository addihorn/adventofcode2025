package main

import (
	"example/hello/src/aocutils"
	"fmt"
)

const inputFile = "input_test.txt"

var inputData []string

var tachyonBeams map[int]bool

const (
	splitter = '^'
)

func main() {
	inputData = aocutils.ReadInput(inputFile)
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	tachyonBeams = make(map[int]bool)
	for i, p := range inputData[0] {
		if p == 'S' {
			tachyonBeams[i] = true
		}
	}
}

/* Solve here */

func part1() {
	initializePuzzle()
	split := 0
	for _, row := range inputData[1:] {
		for i := range tachyonBeams {
			if row[i] == splitter {
				split++
				delete(tachyonBeams, i)
				tachyonBeams[i-1] = true
				tachyonBeams[i+1] = true
			}
		}
	}

	fmt.Printf("Solution for part 1: %d\n", split)
}

func part2() {
	initializePuzzle()
	timelines := 0
	for _, row := range inputData[1:] {
		split := false
		for i := range tachyonBeams {
			if row[i] == splitter {
				split = true
				delete(tachyonBeams, i)
				tachyonBeams[i-1] = true
				tachyonBeams[i+1] = true
			}
		}
		if split {
			timelines += len(tachyonBeams)
		}
	}

	fmt.Printf("Solution for part 1: %d\n", timelines)
}
