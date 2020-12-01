package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readInputFile(filepath string) *[]int {
	var result []int

	rawInput, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil
	}

	rawString := string(rawInput)
	resultString := strings.Split(rawString, "\n")

	for _, v := range resultString {
		n, _ := strconv.Atoi(v)
		result = append(result, n)
	}

	return &result
}

func reportRepair(input []int) int {
	var result int
	sort.Ints(input)

	endIndex := len(input) - 1
	startIndex := endIndex - 1

	for endIndex >= 1 {
		toFind := 2020 - input[endIndex]

		if input[startIndex] == toFind {
			return input[endIndex] * input[startIndex]
		}

		if startIndex == 0 {
			endIndex--
			startIndex = endIndex - 1
		} else {
			startIndex--
		}
	}

	return result
}
