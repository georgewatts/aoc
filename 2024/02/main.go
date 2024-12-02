package main

import (
	"bufio"
	"fmt"
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

func processReport(report []int) bool {
	processed := slices.Compact(report)

	if len(report) != len(processed) {
		return false
	}

	ascending := processed[0] < processed[len(processed)-1]

	for i := 1; i < len(processed); i++ {
		prev := processed[i-1]
		curr := processed[i]
		diff := 0
		if ascending {
			diff = curr - prev
		} else {
			diff = prev - curr
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	report := StrArr{}
	data := [][]int{}

	for len(line) != 0 {
		report = strings.Split(line, " ")
		convertedReport := report.Convert()

		data = append(data, convertedReport)

		scanner.Scan()
		line = scanner.Text()
	}

	count := 0
	for _, v := range data {
		if processReport(v) {
			count++
		}
	}

	fmt.Printf("count: %v\n", count)
}
