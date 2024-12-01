package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var file, err = os.Open("01-2/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var first []int
	var second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")

		firstVal, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		first = append(first, firstVal)

		secondVal, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		second = append(second, secondVal)
	}

	amounts := make(map[int]int)
	for i := 0; i < len(second); i++ {
		val := second[i];
		amounts[val] += 1;
	}

	total := 0
	for i := 0; i < len(first); i++ {
		val := first[i]
		total += val * amounts[val]
	}

	fmt.Println(total)
}
