package main

import (
	"example/hello/src/aocutils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var inputData []string

type idRange struct {
	start, end int
}

var ids []idRange

func main() {
	inputData = aocutils.ReadInputWithDelimeter(inputFile, ",")
	initializePuzzle()
	part1()
	part2()
}

/* Do some puzzle initialization */

func initializePuzzle() {
	ids = make([]idRange, len(inputData))
	r := regexp.MustCompile(`(\d*)-(\d*)`)

	for i, d := range inputData {
		m := r.FindAllStringSubmatch(d, -1)
		start := m[0][1]
		end := m[0][2]
		ids[i] = idRange{
			start: aocutils.CString2Int(start),
			end:   aocutils.CString2Int(end),
		}
	}

}

/* Solve here */

func part1() {

	sum := 0
	for _, id := range ids {
		for i := id.start; i <= id.end; i++ {
			iAsString := strconv.Itoa(i)
			l := len(iAsString)
			// length is odd, so there is no way
			// no repeating number possible
			if l%2 != 0 {
				continue
			}
			if iAsString[:l/2] == iAsString[l/2:] {
				sum += i
			}

		}
	}

	fmt.Printf("Solution for part 1: %d\n", sum)
}

func part2() {
	sum := 0
	for _, id := range ids {
		for i := id.start; i <= id.end; i++ {
			iAsString := strconv.Itoa(i)
			l := len(iAsString)

			for index := 1; index < l/2+1; index++ {

				if iAsString == strings.Repeat(iAsString[:index], l/index) {
					sum += i
					break
				}
			}
		}
	}
	fmt.Printf("Solution for part 2: %d\n", sum)
}
