package main

import (
	"example/hello/src/aocutils"
	"fmt"
	"regexp"
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

type op rune

const (
	mul op = '*'
	add op = '+'
)

type calc struct {
	num        []string
	length     int
	startIndex int
	result     int64
	operation  op
}

var calculations []*calc

/* Do some puzzle initialization */

func initializePuzzle() {

	r := regexp.MustCompile(`[\*\+]{1}\s+`)
	operatorsAsString := inputData[len(inputData)-1]
	m := r.FindAllStringIndex(operatorsAsString, -1)
	calculations = make([]*calc, len(m))

	for i, ops := range m {
		switch operatorsAsString[ops[0] : ops[0]+1] {
		case "*":
			calculations[i] = &calc{
				num:        make([]string, len(inputData)-1),
				startIndex: ops[0],
				length:     ops[1] - ops[0],
				result:     1,
				operation:  mul,
			}
		case "+":
			calculations[i] = &calc{
				num:        make([]string, len(inputData)-1),
				startIndex: ops[0],
				length:     ops[1] - ops[0],
				result:     0,
				operation:  add,
			}
		default:
			continue
		}
		if i != len(m)-1 {
			calculations[i].length--
		}
	}

	for j, data := range inputData[:len(inputData)-1] {

		for _, c := range calculations {
			num := data[c.startIndex : c.startIndex+c.length]
			c.num[j] = num
			numAsInt := aocutils.CString2Int(strings.Trim(num, " "))
			switch c.operation {
			case mul:
				c.result *= int64(numAsInt)
			case add:
				c.result += int64(numAsInt)
			}
		}
	}

}

/* Solve here */

func part1() {
	_ = inputData
	var gt int64 = 0
	for _, c := range calculations {
		gt += c.result
	}
	fmt.Printf("Solution for part 1: %d\n", gt)
}

func part2() {
	var gt int64 = 0
	for _, c := range calculations {

		r := int64(0)
		if c.operation == mul {
			r = 1
		}

		for i := 0; i < c.length; i++ {
			num := ""

			for _, n := range c.num {
				num = fmt.Sprintf("%s%s", num, string(n[i]))
			}
			numAsInt := aocutils.CString2Int(strings.Trim(num, " "))
			switch c.operation {
			case mul:
				r *= int64(numAsInt)
			case add:
				r += int64(numAsInt)
			}

		}

		gt += r

	}

	fmt.Printf("Solution for part 2: %d\n", gt)
}
