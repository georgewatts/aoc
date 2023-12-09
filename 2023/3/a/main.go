package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Matches[T int | string] struct {
	indexes [][]int
	matches []T
}

type Line struct {
	index         int
	symbolMatches Matches[string]
	digitMatches  Matches[int]
}

func determinePositions(index int, s string) *Line {
	symbolRegex := regexp.MustCompile(`[^[:alnum:][:blank:]]`)
	symbolIndexes := symbolRegex.FindAllStringIndex(s, -1)
	symbolMatches := symbolRegex.FindAllString(s, -1)

	digitRegex := regexp.MustCompile(`\d+`)
	digitIndexes := digitRegex.FindAllStringIndex(s, -1)
	digitMatches := digitRegex.FindAllString(s, -1)

	convertedDigitMatches := []int{}

	for _, v := range digitMatches {
		converted, _ := strconv.Atoi(v)
		convertedDigitMatches = append(convertedDigitMatches, converted)
	}

	return &Line{
		index: index,
		symbolMatches: Matches[string]{
			indexes: symbolIndexes,
			matches: symbolMatches,
		},
		digitMatches: Matches[int]{
			indexes: digitIndexes,
			matches: convertedDigitMatches,
		},
	}
}

func (l Line) adjacentParts() int {
	numbers := []int{}

	for i, v := range l.digitMatches.indexes {
		for _, v2 := range l.symbolMatches.indexes {
			if v[0] == v2[len(v2)-1] || v[len(v)-1] == v2[0] {
				fmt.Printf("l.digitMatches.matches[i]: %v\n", l.digitMatches.matches[i])
				numbers = append(numbers, l.digitMatches.matches[i])
			}
		}
	}

	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func verticallyAdjacentPartsTo(l Line, l2 Line) int {
	numbers := []int{}

	for _, v := range l.symbolMatches.indexes {
		// fmt.Printf("v: %v\n", v)
		for i, v2 := range l2.digitMatches.indexes {
			// fmt.Printf("v2: %v\n", v2)
			if (v[0] >= v2[0] && v[0] <= v2[len(v2)-1]) || (v[1] >= v2[0] && v[1] <= v2[len(v2)-1]) {
				// fmt.Printf("symbols check: %v\n", l2.digitMatches.matches[i])
				fmt.Printf("l2.digitMatches.matches[i]: %v\n", l2.digitMatches.matches[i])
				numbers = append(numbers, l2.digitMatches.matches[i])
			}
		}
	}

	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()
	lineIndex := 0
	lines := []*Line{}

	for len(line) != 0 {
		line = strings.ReplaceAll(line, ".", " ")

		fmt.Printf("line %d: %v\n", lineIndex, line)
		lines = append(lines, determinePositions(lineIndex, line))

		scanner.Scan()
		line = scanner.Text()
		lineIndex++
	}

	sum := 0

	for i, l := range lines {
		sum += l.adjacentParts()

		if i > 0 {
			sum += verticallyAdjacentPartsTo(*l, *lines[i-1])
		}

		if i < len(lines)-1 {
			sum += verticallyAdjacentPartsTo(*l, *lines[i+1])
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
