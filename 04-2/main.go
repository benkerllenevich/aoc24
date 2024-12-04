package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("04-2/data.txt")
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
	for i := 1; i < len(lines)-1; i++ {
		line := lines[i]
		for j := 1; j < len(line)-1; j++ {
			if line[j] != 'A' {
				continue
			}

			if ((check(i+1, j-1, 'M') && check(i-1, j+1, 'S')) || (check(i+1, j-1, 'S') && check(i-1, j+1, 'M'))) &&
				((check(i-1, j-1, 'M') && check(i+1, j+1, 'S')) || (check(i-1, j-1, 'S') && check(i+1, j+1, 'M'))) {
				count++
			}
		}
	}

	fmt.Println(count)
}
