package main

import (
	"cmp"
	"example/hello/src/aocutils"
	"fmt"
	"math"
	"slices"
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

type junktionBox struct {
	X, Y, Z              int
	connectedTo, network []*junktionBox
	shortestUnconnected  struct {
		box      *junktionBox
		distance float64
	}
}

/* Do some puzzle initialization */
var jb []*junktionBox

func initializePuzzle() {
	_ = inputData
	jb = make([]*junktionBox, len(inputData))
	for i, box := range inputData {
		coord := strings.Split(box, ",")
		jb[i] = &junktionBox{
			X:           aocutils.CString2Int(coord[0]),
			Y:           aocutils.CString2Int(coord[1]),
			Z:           aocutils.CString2Int(coord[2]),
			connectedTo: []*junktionBox{},
			network:     []*junktionBox{},
			shortestUnconnected: struct {
				box      *junktionBox
				distance float64
			}{
				box:      nil,
				distance: math.MaxFloat64,
			},
		}
		jb[i].connectedTo = []*junktionBox{jb[i]}
		jb[i].network = []*junktionBox{jb[i]}
	}
	for _, box := range jb {
		findNeighbors(box)
	}

}

func findNeighbors(box *junktionBox) {
	for _, neigh := range jb {
		if box == neigh {
			continue
		}
		if slices.Contains(box.connectedTo, neigh) {
			continue
		}
		distance := math.Sqrt(
			float64((box.X-neigh.X)*(box.X-neigh.X)) +
				float64((box.Y-neigh.Y)*(box.Y-neigh.Y)) +
				float64((box.Z-neigh.Z)*(box.Z-neigh.Z)))
		if distance < box.shortestUnconnected.distance {
			box.shortestUnconnected.box = neigh
			box.shortestUnconnected.distance = distance
		}
	}
}

/* Solve here */

func part1() {
	_ = inputData
	initializePuzzle()
	for range 1000 {
		// find shortest path
		min := slices.MinFunc(jb, func(a, b *junktionBox) int {
			return cmp.Compare(a.shortestUnconnected.distance, b.shortestUnconnected.distance)
		})
		// connect the nodes
		n := min.shortestUnconnected.box
		min.connectedTo = append(min.connectedTo, n)
		n.connectedTo = append(n.connectedTo, min)
		newNetwork := func(n1, n2 []*junktionBox) []*junktionBox {
			switch slices.Equal(n1, n2) {
			case true:
				return n1
			default:
				return slices.Concat(n1, n2)
			}
		}(min.network, n.network)
		// recalculate the network
		for _, b := range newNetwork {
			b.network = newNetwork
			b.shortestUnconnected.box = nil
			b.shortestUnconnected.distance = math.MaxFloat64
			findNeighbors(b)
		}

	}

	// sort the vertices
	slices.SortStableFunc(jb, func(a, b *junktionBox) int {
		return cmp.Compare(len(a.network), len(b.network)) * -1
	})
	i := 0
	lc := 1
	for range 3 {
		lc *= len(jb[i].network)
		i += len(jb[i].network)
	}

	fmt.Printf("Solution for part 1: %d\n", lc)
}

func part2() {

	var a, b *junktionBox
	initializePuzzle()
	for len(jb) != len(jb[0].connectedTo) {
		// find shortest path
		min := slices.MinFunc(jb, func(a, b *junktionBox) int {
			return cmp.Compare(a.shortestUnconnected.distance, b.shortestUnconnected.distance)
		})
		// connect the nodes
		n := min.shortestUnconnected.box

		newNetwork := slices.Concat(min.connectedTo, n.connectedTo)
		for _, b := range newNetwork {
			b.connectedTo = newNetwork
			b.shortestUnconnected.box = nil
			b.shortestUnconnected.distance = math.MaxFloat64
			findNeighbors(b)
		}
		a = min
		b = n
	}
	fmt.Printf("Solution for part 2: %d\n", a.X*b.X)
}
