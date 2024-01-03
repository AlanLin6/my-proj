package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func countSum(scanner *bufio.Scanner) (int, string) {
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			return sum, "continue"
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, "invalid number"
		}

		sum += num
	}

	return sum, "end"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text := `1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000`

	maximum := 0
	scanner := bufio.NewScanner(strings.NewReader(text))
	sum := 0
	
	end := "continue"

	for end == "continue" {
		maximum = max(maximum, sum)
		sum, end = countSum(scanner)
	}

	if end == "invalid number" {
		fmt.Println("Invalid number")
		return
	}

	fmt.Println(maximum)
}
