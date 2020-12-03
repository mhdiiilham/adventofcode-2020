package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := loadInput("input.txt")
	second := countTree(input, 3, 1)
	fmt.Println("second:", second)
}

func countTree(input [][]string, right, down int) int {
	result := 0

	posX := right
	posY := down
	maxRow := len(input[0]) - 1

	for posY < len(input) {
		if posX > maxRow {
			xTemp := posX
			posX = xTemp - maxRow - 1
		}

		if input[posY][posX] == "#" {
			result++
			input[posY][posX] = "X"
		} else {
			input[posY][posX] = "O"
		}
		posX = posX + right
		posY = posY + down

	}

	return result
}

func loadInput(fileName string) [][]string {
	var result [][]string

	raws, err := ioutil.ReadFile(fileName)

	if err != nil {
		return result
	}

	for _, raw := range strings.Split(string(raws), "\n") {
		result = append(result, strings.Split(raw, ""))
	}

	return result
}
