package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	totalPairs := 0

	for sc.Scan() {
		assignmentStr := strings.Split(sc.Text(), ",")
		assignments := CreateSliceOfInts(assignmentStr)

		// populate assignment ranges
		var assign1Range []int
		var r int
		for r = assignments[0]; r <= assignments[1]; r++ {
			assign1Range = append(assign1Range, r)
		}

		var assign2Range []int
		var t int
		for t = assignments[2]; t <= assignments[3]; t++ {
			assign2Range = append(assign2Range, t)
		}

		// check for full containment of one or the other
		if fullyContains(assign1Range, assign2Range) || fullyContains(assign2Range, assign1Range) {
			totalPairs += 1
		}

	}
	fmt.Printf("%d\n", totalPairs)
}

func containsElement(i []int, e int) bool {
	for _, a := range i {
		if a == e {
			return true
		}
	}
	return false
}

func fullyContains(s1 []int, s2 []int) bool {
	if len(s1) > len(s2) {
		return false
	}
	for _, element := range s1 {
		if !containsElement(s2, element) {
			return false
		}
	}
	return true
}

func CreateSliceOfInts(s []string) []int {
	var r []int
	for _, v := range s {
		theRange := strings.Split(v, "-")
		low, _ := strconv.Atoi(theRange[0])
		r = append(r, low)
		high, _ := strconv.Atoi(theRange[1])
		r = append(r, high)
	}
	return r
}
