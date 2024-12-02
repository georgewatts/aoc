package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	StrArr []string
	IntArr []int
)

func (arr StrArr) Convert() []int {
	retVal := make([]int, len(arr))
	for i, v := range arr {
		retVal[i], _ = strconv.Atoi(v)
	}
	return retVal
}

func CountElement(arr []int, count int) int {
	retVal := 0
	for _, v := range arr {
		if v == count {
			retVal++
		}
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

	total := 0

	for _, v := range leftSide {
		count := CountElement(rightSide, v)
		total += v * count
	}

	fmt.Printf("total: %v\n", total)
}
