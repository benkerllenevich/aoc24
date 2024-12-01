package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var file, err = os.Open("01-1/data.txt")
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

	sort.Ints(first)
	sort.Ints(second)

	total := 0
	for i := 0; i < len(first); i++ {
		diff := first[i] - second[i]
		total += int(math.Abs(float64(diff)))
	}

	fmt.Println(total)
}
