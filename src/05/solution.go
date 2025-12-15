package main

import (
	"example/hello/src/aocutils"
	"fmt"
	"strings"
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
type (
	storage struct {
		start, stop int
	}
	ingredient struct {
		id      int
		isFresh bool
	}
)

var fridge = make(map[int]*storage)
var ingredients = make(map[int]*ingredient)

func initializePuzzle() {
	_ = inputData
	initFridge := true

	for i, data := range inputData {
		if data == "" {
			initFridge = false
			continue
		}

		if initFridge {
			ing := strings.Split(data, "-")
			s := &storage{
				start: aocutils.CString2Int(ing[0]),
				stop:  aocutils.CString2Int(ing[1]),
			}
			fridge[i] = s
		} else {
			id := aocutils.CString2Int(data)
			ingredients[id] = &ingredient{
				id:      id,
				isFresh: checkFreshness(id),
			}
		}

	}

}

func checkFreshness(id int) bool {

	for _, fr := range fridge {
		if id >= fr.start && id <= fr.stop {
			return true
		}
	}
	return false

}

/* Solve here */

func part1() {
	_ = inputData
	f := 0
	for _, i := range ingredients {
		if i.isFresh {
			f++
		}
	}

	fmt.Printf("Solution for part 1: %d\n", f)
}

func part2() {

	for !consolidate() {
		// consolidate the frige list until the ranges are non overlapping
	}

	ing := 0
	for _, f := range fridge {
		ing += f.stop - f.start + 1
	}

	fmt.Printf("Solution for part 2: %d\n", ing)
}

func consolidate() bool {
	lengthBefore := len(fridge)
	newFridge := make(map[int]*storage)
	init := &storage{
		start: fridge[0].start,
		stop:  fridge[0].stop,
	}
	newFridge[0] = init
	i := 1

	for _, f := range fridge {
		newRange := true
		for _, fn := range newFridge {
			if (f.start <= fn.start && f.stop >= fn.start) || (f.start <= fn.stop && f.stop >= fn.stop) || (f.start >= fn.start && f.stop <= fn.stop) {
				// overlap found
				fn.start = aocutils.Min([]int{fn.start, f.start})
				fn.stop = aocutils.Max([]int{fn.stop, f.stop})
				newRange = false
				break
			}
		}

		if newRange {
			newFridge[i] = &storage{
				start: f.start,
				stop:  f.stop,
			}
			i++
		}
	}
	fridge = newFridge

	return len(fridge) == lengthBefore
}
