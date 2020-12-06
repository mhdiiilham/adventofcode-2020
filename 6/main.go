package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readFromFile("input.txt")
	yesAnswer := countYesAnswer(input[1])
	sumYes := countYes(input)
	fmt.Println("Yes ->", yesAnswer)
	fmt.Println("Sum YES ->", sumYes)

}

func readFromFile(filename string) []string {
	raws, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(raws), "\n\n")
}

func countYesAnswer(answers string) int {
	answered := []string{}
	record := map[string]int{}
	actualAnswer := strings.ReplaceAll(answers, "\n", "")
	answersInArray := strings.Split(actualAnswer, "")
	for _, q := range answersInArray {
		if record[q] == 0 {
			record[q] = 1
			answered = append(answered, q)
		}
	}
	return len(answered)
}

func countYes(listInput []string) int {
	counted := 0

	for _, input := range listInput {
		counted += countYesAnswer(input)
	}

	return counted
}
