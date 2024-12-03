package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Launch a goroutine to run in the background
	//var wg sync.WaitGroup

	var file, err = os.Open("02-2/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var outputChannel = make(chan int)

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		go processLine(scanner.Text(), outputChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(outputChannel)
	}()

	var total = 0
	for num := range outputChannel {
		total += num
	}

	fmt.Println(total)
}

func processLine(line string, output chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	stringArr := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
	var arr = []int{}

	for _, str := range stringArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}

	if !verify(arr) {
		found := false
		for i := 0; i < len(arr); i++ {
			if verify(remove(arr, i)) {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	output <- 1
}

func remove(arr []int, i int) []int {
	new := make([]int, 0, len(arr)-1)
	new = append(new, arr[:i]...)
	new = append(new, arr[i+1:]...)
	return new
}

func verify(arr []int) bool {
	max := len(arr) - 1
	var increasing bool
	safe := true

	for i := 0; i < len(arr); i++ {
		if i == max {
			continue
		}

		current := arr[i]
		next := arr[i+1]
		diff := current - next

		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff < 1 {
			safe = false
			break
		}

		increased := diff > 0
		if i == 0 {
			increasing = diff > 0
		}

		if increasing != increased {
			safe = false
			break
		}
	}

	return safe
}
