package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	winningNumbers []string
	chosenNumbers  []string
	copies         int
}

func createCard(s string) Card {
	numbers := strings.Split(s, " | ")

	winningNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[0], -1)
	chosenNumbers := regexp.MustCompile(`\d+`).FindAllString(numbers[1], -1)
	copies := 1

	return Card{
		winningNumbers,
		chosenNumbers,
		copies,
	}
}

func (card Card) countWinningNumbers() int {
	count := 0

	for _, v := range card.chosenNumbers {
		for _, v2 := range card.winningNumbers {
			if v == v2 {
				count++
			}
		}
	}

	return count
}

func processCards(cards []Card) []Card {
	cardsCopy := make([]Card, len(cards))
	copy(cardsCopy, cards)

	// totalCollection := make([][]Card, len(cards))

	for i, c := range cardsCopy {
		winningCount := c.countWinningNumbers()

		fmt.Printf("winningCount: %v\n", winningCount)

		for j := 0; j < c.copies; j++ {
			for i2 := i + 1; i2 < i+1+winningCount; i2++ {
				cardsCopy[i2].copies++
			}
		}
	}

	return cardsCopy
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	cards := []Card{}

	for len(line) != 0 {
		card := strings.Split(line, ": ")

		cards = append(cards, createCard(card[1]))

		scanner.Scan()
		line = scanner.Text()
	}

	processedCards := processCards(cards)

	sum := 0
	for _, c := range processedCards {
		sum += c.copies
	}

	fmt.Printf("sum: %v\n", sum)
}
