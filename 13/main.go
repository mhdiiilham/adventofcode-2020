package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readFromFile("input.txt")
	partOne := getEarliestDepartPossible(input)
	fmt.Println("Part One:", partOne)
}

func getEarliestDepartPossible(input []string) int {
	listPossibleDepart := []int{}
	minDepart, err := strconv.Atoi(input[0])
	matchDepartTime := map[int]int{}
	if err != nil {
		return 0
	}

	for _, busID := range strings.Split(input[1], ",") {
		if busID != "x" {
			busIDInt, err := strconv.Atoi(busID)
			if err != nil {
				return 0
			}
			possibleDepart := closerToMinDepart(busIDInt, minDepart)
			matchDepartTime[possibleDepart] = busIDInt
			listPossibleDepart = append(listPossibleDepart, possibleDepart)
		}
	}

	sort.Ints(listPossibleDepart)
	return matchDepartTime[listPossibleDepart[0]] * (listPossibleDepart[0] - minDepart)
}

func closerToMinDepart(busID, minDepartTime int) int {
	busDepart := 0
	for busDepart < minDepartTime {
		busDepart += busID
	}
	return busDepart
}

func multiplyingTime(busID, minDepartTime int) int {
	minWaitTime := 0
	busDepart := 0

	for busDepart < minDepartTime {
		busDepart += busID
	}

	minWaitTime = busDepart - minDepartTime

	return minWaitTime * busID
}

func readFromFile(fileName string) []string {
	raws, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(raws), "\n")
}
