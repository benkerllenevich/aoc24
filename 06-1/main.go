package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func main() {
	var file, err = os.Open("06-1/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var parsedMap [][]bool
	var guardX, guardY int
	direction := UP

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		var currentLine []bool

		for j, char := range line {
			currentLine = append(currentLine, char == '#')
			if char == '^' {
				guardX = j
				guardY = i
			}
		}

		parsedMap = append(parsedMap, currentLine)
		i++
	}

	maxX := len(parsedMap[0]) - 1
	maxY := len(parsedMap) - 1
	visited := make([][]int, maxY+1)

	visited[guardY] = append(visited[guardY], guardX)

	for {
		nextX := guardX
		nextY := guardY

		switch direction {
		case UP:
			nextY -= 1
		case RIGHT:
			nextX += 1
		case DOWN:
			nextY += 1
		case LEFT:
			nextX -= 1
		}

		if nextX > maxX || nextX < 0 || nextY > maxY || nextY < 0 {
			break
		}

		if parsedMap[nextY][nextX] {
			direction += 1
			if direction > LEFT {
				direction = UP
			}
			continue
		}

		guardX = nextX
		guardY = nextY

		if !slices.Contains(visited[guardY], guardX) {
			visited[guardY] = append(visited[guardY], guardX)
		}
	}

	total := 0
	for _, arr := range visited {
		total += len(arr)
	}
	fmt.Println(total)
}
