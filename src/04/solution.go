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

type paper struct {
	p          rune
	X, Y       int
	accessible bool
}

var room = make(map[[2]int]*paper)
var gridsize *aocutils.Gridsize

// for visualization
var grid = make(map[[2]int]rune)

const (
	paperroll = '@'
	freeSpot  = '.'
)

/* Do some puzzle initialization */

func initializePuzzle() {
	_ = inputData

	gridsize = &aocutils.Gridsize{MinX: 0, MinY: 0, MaxX: len(inputData[0]) - 1, MaxY: len(inputData) - 1}

	for y, data := range inputData {
		for x, spot := range data {
			if spot == freeSpot {
				continue
			}
			grid[[2]int{x, y}] = spot
			room[[2]int{x, y}] =
				&paper{
					p:          spot,
					X:          x,
					Y:          y,
					accessible: false}

		}
	}

}

/* Solve here */

func part1() {
	_ = inputData

	fmt.Printf("Solution for part 1: %d\n", findAccessible())
}

func part2() {
	_ = inputData
	a := 0
	for f := findAccessible(); f > 0; f = findAccessible() {
		a += f
		removeAccessible()
	}

	fmt.Printf("Solution for part 2: %d\n", a)
}

func findAccessible() (a int) {
	a = 0

	for _, p := range room {
		neighbors := 0
		for x := p.X - 1; x < p.X+2; x++ {
			for y := p.Y - 1; y < p.Y+2; y++ {
				if x == p.X && y == p.Y {
					continue
				}
				if _, ok := room[[2]int{x, y}]; ok {
					neighbors++
				}

			}
		}
		if neighbors < 4 {
			a++
			p.accessible = true
		}
	}
	return a
}
func removeAccessible() {
	for _, p := range room {
		if p.accessible {
			delete(room, [2]int{p.X, p.Y})
		}
	}
}
