package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(first int, second int) int {
	if first > second {
		return second
	} else {
		return first
	}
}

func countTaps(value int) int {
	return value/4 + min(value%4, 2)
}

func main() {
	var countStrings int

	fmt.Scan(&countStrings)

	scanner := bufio.NewScanner(os.Stdin)

	result := 0

	for i := 0; i < countStrings; i++ {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		num, _ := strconv.Atoi(line)

		result += countTaps(num)
	}

	fmt.Print(result)
}
