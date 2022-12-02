package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	maxCals := 0
	currentElfCals := 0

	// scan the elves
	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		// add snack calories to the elf total
		currentElfCals += snack

		if err != nil {
			if currentElfCals > maxCals {
				maxCals = currentElfCals
			}
			// found blank line, start with new elf
			currentElfCals = 0
		}
	}
	fmt.Println(maxCals)
}
