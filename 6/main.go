package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := readFromFile("input.txt")
	sumYes := countYes(input)
	allYes := everyoneYes(input)
	fmt.Println("Part One ->", sumYes)
	fmt.Println("Part Two ->", allYes)

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

func countAllYes(groupAnswer string) int {
	allYes := 0
	splitByPerson := strings.Split(groupAnswer, "\n")
	totalPerson := len(splitByPerson)
	answerRecord := map[string]int{}

	for _, byPerson := range splitByPerson {
		byPersonSplit := strings.Split(byPerson, "")
		for _, byPersonAnswer := range byPersonSplit {
			if answerRecord[byPersonAnswer] == 0 {
				answerRecord[byPersonAnswer] = 1
			} else {
				answerRecord[byPersonAnswer]++
			}
		}
	}
	for _, record := range answerRecord {
		if record == totalPerson {
			allYes++
		}
	}
	return allYes
}

func everyoneYes(listInput []string) int {
	counted := 0
	for _, input := range listInput {
		counted += countAllYes(input)
	}
	return counted
}

func countYes(listInput []string) int {
	counted := 0
	for _, input := range listInput {
		counted += countYesAnswer(input)
	}
	return counted
}
