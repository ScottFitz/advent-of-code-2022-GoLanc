package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

var (
	ucAlphabet []string
	lcAlphabet []string
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	totalPriorities := 0

	// populate alphabets
	var i byte
	for i = 'A'; i <= 'Z'; i++ {
		ucAlphabet = append(ucAlphabet, string(i))
	}

	var x byte
	for x = 'a'; x <= 'z'; x++ {
		lcAlphabet = append(lcAlphabet, string(x))
	}

	for sc.Scan() {
		compartments := splitTheString(sc.Text())
		compartment1 := compartments[0]
		compartment2 := compartments[1]
		// find element in both
		var found string
		for _, c := range compartment1 {
			if strings.Index(compartment2, string(c)) != -1 {
				found = string(c)
				break
			}
		}
		totalPriorities += getPriorityValue(found)
	}
	fmt.Printf("%d\n", totalPriorities)
}

func splitTheString(s string) []string {
	var compartments []string
	runes := []rune(s)
	if len(runes) == 0 {
		return []string{s}
	}
	size := len(s) / 2
	for i := 0; i < len(runes); i += size {
		nn := i + size
		if nn > len(runes) {
			nn = len(runes)
		}
		compartments = append(compartments, string(runes[i:nn]))
	}
	return compartments
}

func getPriorityValue(s string) int {
	res := slices.Index(ucAlphabet, s)
	if res != -1 {
		return res + 27
	} else {
		res = slices.Index(lcAlphabet, s)
		return res + 1
	}
}
