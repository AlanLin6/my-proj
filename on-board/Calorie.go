package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func countSum() (int, string) {
	sum := 0

	for {
		text, err := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n") // Consume the next line character
		if text == "" {
			break
		}
		if err == io.EOF {
			return sum, "end"
		}
		num, _ := strconv.Atoi(text)
		sum += num
	}

	return sum, "continue"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader = bufio.NewReader(file)
	maximum := 0
	for {
		sum, status := countSum()
		if status == "end" {
			break
		}
		maximum = max(maximum, sum)
	}
	fmt.Println(maximum)
}
