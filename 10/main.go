package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
}

func readFromFile(filePath string) []int {
	input := []int{}
	raws, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []int{}
	}
	for _, jolt := range strings.Split(string(raws), "\n") {
		joltToInt, _ := strconv.Atoi(jolt)
		input = append(input, joltToInt)
	}
	return input
}
