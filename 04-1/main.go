package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("04-1/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check := func(i int, j int, char byte) bool {
		line := lines[i]
		if line[j] == char {
			return true
		}
		return false
	}

	count := 0

	searchUp := false
	searchDown := true
	for i, line := range lines {
		if i == 3 {
			searchUp = true
		}
		if i == len(lines)-3 {
			searchDown = false
		}

		searchLeft := false
		searchRight := true
		for j, char := range line {
			if j == 3 {
				searchLeft = true
			}
			if j == len(line)-3 {
				searchRight = false
			}

			if char != 'X' {
				continue
			}

			if searchUp {
				if check(i-1, j, 'M') && check(i-2, j, 'A') && check(i-3, j, 'S') {
					count += 1
				}

				if searchLeft && check(i-1, j-1, 'M') && check(i-2, j-2, 'A') && check(i-3, j-3, 'S') {
					count += 1
				}

				if searchRight && check(i-1, j+1, 'M') && check(i-2, j+2, 'A') && check(i-3, j+3, 'S') {
					count += 1
				}
			}

			if searchDown {
				if check(i+1, j, 'M') && check(i+2, j, 'A') && check(i+3, j, 'S') {
					count += 1
				}

				if searchLeft && check(i+1, j-1, 'M') && check(i+2, j-2, 'A') && check(i+3, j-3, 'S') {
					count += 1
				}

				if searchRight && check(i+1, j+1, 'M') && check(i+2, j+2, 'A') && check(i+3, j+3, 'S') {
					count += 1
				}
			}

			if searchLeft && check(i, j-1, 'M') && check(i, j-2, 'A') && check(i, j-3, 'S') {
				count += 1
			}

			if searchRight && check(i, j+1, 'M') && check(i, j+2, 'A') && check(i, j+3, 'S') {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
