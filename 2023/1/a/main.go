package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func parseIntFromString(s string, count int) (int, error) {
	foundChars := regexp.MustCompile(`\d`).FindAllString(s, -1)

	first := foundChars[0]
	last := foundChars[len(foundChars)-1]

	println(count, first, last)

	doubleDigit := first + last

	return strconv.Atoi(doubleDigit)
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	sum, count := 0, 0

	for len(line) != 0 {
		doubledigit, _ := parseIntFromString(line, count)
		// print(" ", doubledigit, "\n")
		sum += doubledigit
		scanner.Scan()
		line = scanner.Text()
		count++
	}

	println("sum: ", sum)
}
