package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	rock     = "A"
	paper    = "B"
	scissors = "C"

	myRock     = "X"
	myPaper    = "Y"
	myScissors = "Z"
)

func main() {
	input, _ := os.Open("input.csv")
	defer input.Close()
	sc := bufio.NewScanner(input)

	totalScore := 0

	for sc.Scan() {
		fmt.Printf("%s\n", sc.Text())
		strategy := strings.Split(sc.Text(), " ")
		if len(strategy) > 1 {
			theirChoice := strategy[0]
			myChoice := strategy[1]

			switch theirChoice {
			case rock:
				switch myChoice {
				case myRock:
					totalScore += 4
				case myPaper:
					totalScore += 8
				case myScissors:
					totalScore += 3
				}
			case paper:
				switch myChoice {
				case myRock:
					totalScore += 1
				case myPaper:
					totalScore += 5
				case myScissors:
					totalScore += 9
				}
			case scissors:
				switch myChoice {
				case myRock:
					totalScore += 7
				case myPaper:
					totalScore += 2
				case myScissors:
					totalScore += 6
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", totalScore)
}
