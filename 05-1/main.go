package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var file, err = os.Open("05-1/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parsingRules := true
	rules := make([][]int, 100)
	var safe [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			first, _ := strconv.Atoi(line[:2])
			second, _ := strconv.Atoi(line[3:])

			rules[first] = append(rules[first], second)
			continue
		}

		stringSet := strings.Split(line, ",")
		var checked []int
		passed := true

		for _, val := range stringSet {
			i, _ := strconv.Atoi(val)
			currentRules := rules[i]

			if passed {
				for _, rule := range currentRules {
					if slices.Contains(checked, rule) {
						passed = false
						break
					}
				}
			}

			checked = append(checked, i)
		}

		if passed {
			safe = append(safe, checked)
		}
	}

	total := 0
	for _, safeSet := range safe {
		total += safeSet[len(safeSet)/2]
	}
	fmt.Println(total)
}
