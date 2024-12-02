package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type StrArr []string

func (arr StrArr) Convert() []int {
	retVal := make([]int, len(arr))
	for i, v := range arr {
		retVal[i], _ = strconv.Atoi(v)
	}
	return retVal
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	leftInts, rightInts := StrArr{}, StrArr{}

	for len(line) != 0 {
		ints := strings.Split(line, "   ")

		leftInts = append(leftInts, ints[0])
		rightInts = append(rightInts, ints[1])

		scanner.Scan()
		line = scanner.Text()
	}

	leftSide := leftInts.Convert()
	rightSide := rightInts.Convert()

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	totalDistance := 0

	for i := 0; i < len(leftSide); i++ {
		totalDistance += int(math.Abs(float64(leftSide[i]) - float64(rightSide[i])))
	}

	fmt.Printf("totalDistance: %v\n", totalDistance)
}
