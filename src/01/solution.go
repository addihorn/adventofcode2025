package main

import (
	"example/hello/src/aocutils"
	"fmt"
	"regexp"
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
	data := inputData
	start := 50
	r := regexp.MustCompile(`([L|R]{1})(\d*)`)
	password := 0

	for _, d := range data {

		a := r.FindAllStringSubmatch(d, -1)
		if len(a) == 0 {
			break
		}
		delta := aocutils.CString2Int(a[0][2])
		delta = delta % 100
		switch a[0][1] {
		case "L":
			start = circle(start - delta)
		case "R":
			start = circle(start + delta)
		}
		if start == 0 {
			password++
		}
		//fmt.Println(delta, start, password)
	}
	fmt.Println("Solution for part 1: ", password)
}

func part2() {
	data := inputData
	start := 50
	r := regexp.MustCompile(`([L|R]{1})(\d*)`)
	password := 0

	for _, d := range data {
		a := r.FindAllStringSubmatch(d, -1)
		if len(a) == 0 {
			break
		}
		delta := aocutils.CString2Int(a[0][2])
		password += delta / 100
		delta = delta % 100
		s := start
		switch a[0][1] {
		case "L":
			start = circle(start - delta)
			if start > s && s != 0 {
				password++
			}
		case "R":
			start = circle(start + delta)
			if start < s && start != 0 {
				password++
			}
		}
		if start == 0 {
			password++
		}
		//fmt.Println(s, delta, start, password)
	}
	fmt.Printf("Solution for part 2: %d\n", password)
}

func circle(n int) int {
	if n > 99 {
		return n - 100
	}
	if n < 0 {
		return n + 100
	}
	return n
}
