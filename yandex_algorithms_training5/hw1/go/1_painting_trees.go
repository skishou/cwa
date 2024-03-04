package main

import (
	"fmt"
)

func countMax(firstDigit int, secondDigit int) int {
	if firstDigit < secondDigit {
		return secondDigit
	} else if firstDigit > secondDigit {
		return firstDigit
	} else {
		return secondDigit
	}
}

func countMin(firstDigit int, secondDigit int) int {
	if firstDigit < secondDigit {
		return firstDigit
	} else if firstDigit > secondDigit {
		return secondDigit
	} else {
		return firstDigit
	}
}

func countTrees(boyTreeNumber int, boyStep int, girlTreeNumber int, girlStep int) int {
	allVariant := 2 * (boyStep + girlStep + 1)
	maxLow := countMax(boyTreeNumber-boyStep, girlTreeNumber-girlStep)
	minHigh := countMin(boyTreeNumber+boyStep, girlTreeNumber+girlStep)

	result := allVariant - countMax(0, minHigh-maxLow+1)

	return result
}

func main() {
	var p, v, q, m int
	fmt.Scan(&p, &v, &q, &m)

	trees := countTrees(p, v, q, m)

	fmt.Print(trees)
}
