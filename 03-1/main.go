package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("03-1/data.txt")
	if err != nil {
		panic(err)
	}

	str := string(content)
	total := 0

	fail := func(index int) {
		str = str[index:]
	}

	for {
		i := strings.Index(str, "mul(")
		if i == -1 {
			break
		}

		num1, nextI := parseNumber(str, i + 4)
		if nextI >= len(str) {
			break
		}

		if num1 == -1 || str[nextI] != ',' {
			fail(nextI - 1)
			continue
		}

		num2, nextI := parseNumber(str, nextI + 1)
		if nextI >= len(str) {
			break
		}

		if num2 == -1 || str[nextI] != ')' {
			fail(nextI - 1)
			continue
		}

		total += num1 * num2
		fail(nextI + 1)
	}

	fmt.Println(total)
}

func checkNum(input string) bool {
	if _, err := strconv.Atoi(input); err == nil {
		return true
	}
	return false
}

func parseNumber(str string, i int) (int, int) {
	if !checkNum(string(str[i])) {
		return -1, i + 1
	}

	numVal := ""
	for j := i; j < len(str); j++ {
		current := string(str[j])
		if checkNum(current) {
			numVal = numVal + current
		} else {
			break
		}
	}

	val, err := strconv.Atoi(numVal)
	if err != nil {
		panic(err)
	}

	return val, i + len(numVal)
}
