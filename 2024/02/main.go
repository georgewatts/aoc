package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type StrArr []string

func Convert(arr []string) []int {
	retVal := make([]int, len(arr))
	for i, v := range arr {
		retVal[i], _ = strconv.Atoi(v)
	}
	return retVal
}

func Abs(integer int) int {
	if integer < 0 {
		return integer * -1
	}
	return integer
}

func calcValues(report []int, ascending bool) bool {
	for i := 1; i < len(report); i++ {
		prev := report[i-1]
		curr := report[i]
		if prev == 0 || curr == 0 {
			panic("oops")
		}
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

func processReport(report []int) bool {
	// processed := slices.Compact(report)
	processed := slices.Clone(report)

	// if len(report) != len(processed) {
	// 	return false
	// }

	ascending := processed[0] < processed[len(processed)-1]

	result := calcValues(processed, ascending)

	if result {
		return true
	}

	if !result {
		index := 0
		for index < len(processed) {
			reportRetry := slices.Clone(processed)
			reportRetry = slices.Delete(reportRetry, index, index+1)
			result := calcValues(reportRetry, ascending)
			if result {
				return true
			}
			index++
		}
	}

	log.Println("unsafe", processed)
	return false
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	data := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		convertedReport := Convert(report)

		data = append(data, convertedReport)
	}

	count := 0
	for _, v := range data {
		if processReport(v) {
			count++
		}
	}

	fmt.Printf("count: %v\n", count)
}
