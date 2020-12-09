package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readFromFile("test.txt")
	acc := accumulatorValue(input)
	fmt.Println("Part one asnwer->", acc)
	fmt.Println(isContain([]int{}, 0))
}

func accumulatorValue(bootCodes []string) int {
	accumulator := 0
	visitedIndex := []int{}
	currentIndex := 0

	for currentIndex < len(bootCodes)-1 {
		extractCode := strings.Split(bootCodes[currentIndex], " ")
		operation := extractCode[0]
		argument := strings.Split(extractCode[1], "")[0]
		argumentValue := strings.ReplaceAll(extractCode[1], "+", "")
		argumentValue = strings.ReplaceAll(extractCode[1], "-", "")
		argumentValueToInt, _ := strconv.Atoi(argumentValue)

		if isContain(visitedIndex, currentIndex) {
			return accumulator
		}

		visitedIndex = append(visitedIndex, currentIndex)

		if operation == "acc" {
			switch argument {
			case "-":
				accumulator -= argumentValueToInt
				break
			case "+":
				accumulator += argumentValueToInt
				break
			}
			currentIndex++
		} else if operation == "jmp" {
			switch argument {
			case "-":
				currentIndex -= argumentValueToInt
				break
			case "+":
				currentIndex += argumentValueToInt
				break
			}
		} else {
			currentIndex++
		}

	}

	return accumulator
}

func readFromFile(file string) []string {
	raws, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(raws), "\n")
}

func isContain(s []int, t int) bool {
	for _, e := range s {
		if e == t {
			return true
		}
	}
	return false
}
