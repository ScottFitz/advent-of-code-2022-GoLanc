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

	maxCals1 := 0
	maxCals2 := 0
	maxCals3 := 0
	currentElfCals := 0

	// scan the elves
	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		// add snack calories to the elf total
		currentElfCals += snack

		if err != nil {
			if currentElfCals > maxCals3 {
				maxCals3 = currentElfCals
			}
			if maxCals3 > maxCals2 {
				maxCals3, maxCals2 = maxCals2, maxCals3
			}
			if maxCals2 > maxCals1 {
				maxCals2, maxCals1 = maxCals1, maxCals2
			}
			// found blank line, start with new elf
			currentElfCals = 0
		}
	}
	fmt.Println(maxCals1 + maxCals2 + maxCals3)
}
