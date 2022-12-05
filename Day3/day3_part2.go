package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

var (
	upperAlphabet []string
	lowerAlphabet []string
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	totalPriorities := 0

	// populate alphabets
	var i byte
	for i = 'A'; i <= 'Z'; i++ {
		upperAlphabet = append(upperAlphabet, string(i))
	}

	var x byte
	for x = 'a'; x <= 'z'; x++ {
		lowerAlphabet = append(lowerAlphabet, string(x))
	}

	groupCount := 0
	var groupRucksacks []string

	for sc.Scan() {
		groupRucksacks = append(groupRucksacks, sc.Text())
		groupCount += 1

		if groupCount%3 == 0 {
			rucksack1 := groupRucksacks[0]
			rucksack2 := groupRucksacks[1]
			rucksack3 := groupRucksacks[2]

			var found string
			for _, c := range rucksack1 {
				if strings.Index(rucksack2, string(c)) != -1 && strings.Index(rucksack3, string(c)) != -1 {
					found = string(c)
					break
				}
			}
			groupRucksacks = nil
			totalPriorities += getPriority(found)
		}
	}
	fmt.Printf("%d\n", totalPriorities)
}

func getPriority(s string) int {
	res := slices.Index(upperAlphabet, s)
	if res != -1 {
		return res + 27
	}
	res = slices.Index(lowerAlphabet, s)
	return res + 1

}
