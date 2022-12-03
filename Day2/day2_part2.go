package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	rock := "A"
	paper := "B"
	scissors := "C"

	lose := "X"
	draw := "Y"
	win := "Z"
	totalScore := 0

	for sc.Scan() {
		fmt.Printf("%s\n", sc.Text())
		strategy := strings.Split(sc.Text(), " ")
		if len(strategy) > 1 {
			theirChoice := strategy[0]
			endResult := strategy[1]

			switch theirChoice {
			case rock:
				switch endResult {
				case lose:
					totalScore += 3
				case draw:
					totalScore += 4
				case win:
					totalScore += 8
				}
			case paper:
				switch endResult {
				case lose:
					totalScore += 1
				case draw:
					totalScore += 5
				case win:
					totalScore += 9
				}
			case scissors:
				switch endResult {
				case lose:
					totalScore += 2
				case draw:
					totalScore += 6
				case win:
					totalScore += 7
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", totalScore)
}
