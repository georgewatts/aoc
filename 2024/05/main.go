package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rules map[int][]int

func (r Rules) append(page int, suffix int) {
	r[page] = append(r[page], suffix)
}

func (r Rules) validatePage(page int, suffix int) bool {
	return slices.Contains(r[page], suffix)
}

func (r Rules) validateUpdate(update string) int {
	pages := strings.Split(update, ",")
	pageNumbers := []int{}

	for _, v := range pages {
		page, _ := strconv.Atoi(v)
		pageNumbers = append(pageNumbers, page)
	}

	for _, v := range pageNumbers {
		successive := getSucceedingInts(pageNumbers, v)
		for _, v2 := range successive {
			if !r.validatePage(v, v2) {
				return 0
			}
		}
	}

	return pageNumbers[(len(pageNumbers)-1)/2]
}

func getPrecedingInts(arr []int, intToFind int) []int {
	return arr[:slices.Index(arr, intToFind)]
}

func getSucceedingInts(arr []int, intToFind int) []int {
	return arr[slices.Index(arr, intToFind)+1:]
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	rules := Rules{}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			first, _ := strconv.Atoi(pages[0])
			second, _ := strconv.Atoi(pages[1])
			rules.append(first, second)
		}

		if strings.Contains(line, ",") {
			total += rules.validateUpdate(line)
		}
	}

	fmt.Printf("total: %v\n", total)
}
