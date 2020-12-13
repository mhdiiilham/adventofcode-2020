package main

import (
	"io/ioutil"
	"strings"
)

func main() {

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
