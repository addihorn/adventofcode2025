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

type tile struct {
	X, Y int
}

var tileSet []*tile

/* Do some puzzle initialization */

func initializePuzzle() {
	tileSet = make([]*tile, len(inputData))
	for i, data := range inputData {
		coord := strings.Split(data, ",")
		tileSet[i] = &tile{
			X: aocutils.CString2Int(coord[0]),
			Y: aocutils.CString2Int(coord[1]),
		}
	}
}

/* Solve here */

func part1() {
	maxRect := 0
	for i, t1 := range tileSet {
		for _, t2 := range tileSet[i+1:] {
			area := (aocutils.Abs(t1.X-t2.X) + 1) * (aocutils.Abs(t1.Y-t2.Y) + 1)
			maxRect = aocutils.Max([]int{maxRect, area})
		}
	}

	fmt.Printf("Solution for part 1: %d\n", maxRect)
}

func part2() {
	_ = inputData
	type edge struct {
		fixed, from, to int
	}
	vEdge := []edge{}
	hEdge := []edge{}

	// build up the edge set
	for i, t1 := range tileSet {

		var t2 *tile
		switch i {
		case 0:
			t2 = tileSet[len(tileSet)-1]
		default:
			t2 = tileSet[i-1]
		}

		//vertical edge
		if t1.X == t2.X {

			vEdge = append(vEdge, edge{
				fixed: t1.X,
				from:  aocutils.Min([]int{t1.Y, t2.Y}),
				to:    aocutils.Max([]int{t1.Y, t2.Y}),
			})
		} else {
			hEdge = append(hEdge, edge{
				fixed: t1.Y,
				from:  aocutils.Min([]int{t1.X, t2.X}),
				to:    aocutils.Max([]int{t1.X, t2.X}),
			})
		}

	}

	maxRect := 0

	for i, t1 := range tileSet {
		for _, t2 := range tileSet[i+1:] {

			upperLeftX := aocutils.Min([]int{t1.X, t2.X})
			upperLeftY := aocutils.Min([]int{t1.Y, t2.Y})

			bottomRightX := aocutils.Max([]int{t1.X, t2.X})
			bottomRightY := aocutils.Max([]int{t1.Y, t2.Y})

			f1, s1, e1 := 0, 0, 0
			f2, s2, e2 := 0, 0, 0
			valid := true

			//vertical left
			f1 = upperLeftX
			s1 = upperLeftY
			e1 = bottomRightY

			// vertical right
			f2 = bottomRightX
			s2 = upperLeftY
			e2 = bottomRightY

			for _, e := range hEdge {
				// cuts left vertical?
				if e.fixed > s1 && e.fixed < e1 {
					// line could cut
					if e.from <= f1 && e.to > f1 {
						// found a cutting line
						// so this is a non valid edge
						valid = false
						break
					}
				}
				// cuts right vertical?
				if e.fixed > s2 && e.fixed < e2 {
					if e.from < f2 && e.to >= f2 {
						// found a cutting line
						// so this is a non valid edge
						valid = false
						break
					}
				}
			}
			// end the check, if we do not have a valid square
			if !valid {
				continue
			}

			//horizontal top
			f1 = upperLeftY
			s1 = upperLeftX
			e1 = bottomRightX

			// horizontal bottom
			f2 = bottomRightY
			s2 = upperLeftX
			e2 = bottomRightX

			for _, e := range vEdge {
				// cuts top horizontal?
				if e.fixed > s1 && e.fixed < e1 {
					// line could cut
					if e.from <= f1 && e.to > f1 {
						// found a cutting line
						// so this is a non valid edge
						valid = false
						break
					}
				}
				// cuts bottom horozontal?
				if e.fixed > s2 && e.fixed < e2 {
					if e.from < f2 && e.to >= f2 {
						// found a cutting line
						// so this is a non valid edge
						valid = false
						break
					}
				}
			}
			// end the check, if we do not have a valid square
			if !valid {
				continue
			}

			area := (aocutils.Abs(t1.X-t2.X) + 1) * (aocutils.Abs(t1.Y-t2.Y) + 1)
			maxRect = aocutils.Max([]int{maxRect, area})

		}
	}

	fmt.Printf("Solution for part 2: %d\n", maxRect)
}
