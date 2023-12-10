package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func pointsForCard(s string) int {
	numbers := strings.Split(s, " | ")

	// winningNumbers := strings.Split(strings.TrimSpace(numbers[0]), " ")
	// chosenNumbers := strings.Split(strings.TrimSpace(numbers[1]), " ")
	winningNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[0], -1)
	chosenNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[1], -1)

	fmt.Printf("winningNumbers: %v\n", winningNumbers)
	fmt.Printf("chosenNumbers: %v\n", chosenNumbers)

	points := 0

	for _, v := range chosenNumbers {
		for _, v2 := range winningNumbers {
			if v == v2 {
				if points == 0 {
					points = 1
				} else {
					points = points + points
				}
			}
		}
	}

	return points
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	sum := 0

	for len(line) != 0 {
		// card := regexp.MustCompile(`Card\s\d+:\s([\d+[:blank:]]+)`).FindAllStringSubmatch(line, -1)
		card := strings.Split(line, ": ")

		sum += pointsForCard(card[1])

		scanner.Scan()
		line = scanner.Text()
	}

	fmt.Printf("sum: %v\n", sum)
}
