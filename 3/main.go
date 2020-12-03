package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := loadInput("input.txt")
	result := countThree(input)
	fmt.Println(result)
}

func countThree(input [][]string) int {
	var result int

	posX := 3
	posY := 1
	maxRow := len(input[posY]) - 1

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
		posX = posX + 3
		posY++

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
