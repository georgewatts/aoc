package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var conversionMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3ree",
	"four":  "f4ur",
	"five":  "f5ve",
	"six":   "s6x",
	"seven": "s7ven",
	"eight": "e8ght",
	"nine":  "n9ne",
}

// var conversionMap = map[string]string{
// 	"one":   "1",
// 	"two":   "2",
// 	"three": "3",
// 	"four":  "4",
// 	"five":  "5",
// 	"six":   "6",
// 	"seven": "7",
// 	"eight": "8",
// 	"nine":  "9",
// }

// FYI maps are already a reference type underneath?
func parseIntFromString(s string, count int) (int, error) {
	for k, v := range conversionMap {
		s = strings.ReplaceAll(s, k, v)
	}
	//|one|two|three|four|five|six|seven|eight|nine
	foundChars := regexp.MustCompile(`\d`).FindAllString(s, -1)
	// strings.
	// fmt.Printf("%s \n", foundChars)

	first := foundChars[0]
	last := foundChars[len(foundChars)-1]

	// if val, ok := conversionMap[first]; ok {
	// 	first = val
	// }

	// if val, ok := conversionMap[last]; ok {
	// 	last = val
	// }

	doubleDigit := first + last

	println(count, first, last, doubleDigit)

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

	sum, count := 0, 1

	for len(line) != 0 {
		doubledigit, err := parseIntFromString(line, count)
		if err != nil {
			println("error", err)
		}
		// print(" ", doubledigit, "\n")
		sum += doubledigit
		scanner.Scan()
		line = scanner.Text()
		count++
	}

	println("sum: ", sum)
}
