package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func countSumHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("numbers")

	maximum, err := calculateMaximum(text)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, maximum)
}

func calculateMaximum(text string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	sum := 0
	maximum := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			maximum = max(maximum, sum)
			sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %v", err)
		}

		sum += num
		maximum = max(maximum, sum)
	}

	maximum = max(maximum, sum)

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error during scanning: %v", err)
	}

	return maximum, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Sum Calculator</title>
		</head>
		<body>
			<h1>Sum Calculator</h1>
			<form method="post" action="/countSum">
				<label for="numbers">Enter numbers (separated by newline):</label><br>
				<textarea name="numbers" rows="10" cols="30"></textarea><br>
				<input type="submit" value="Calculate Sum">
			</form>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	http.HandleFunc("/countSum", countSumHandler)
	http.HandleFunc("/", indexHandler)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
