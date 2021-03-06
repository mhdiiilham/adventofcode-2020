package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := loadInputFromFile("input.txt")
	validByPosPassword := validateByIndex(input)
	validByRange := countInput(input)
	fmt.Println(fmt.Sprintf(
		"Valida by range: %v, valid by exactly one position %v",
		validByRange,
		validByPosPassword,
	))
}

func loadInputFromFile(filePath string) []string {
	raws, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []string{}
	}
	return strings.Split(string(raws), "\n")
}

func countInput(input []string) int {
	var result int

	for _, passwd := range input {
		min, max := extractMinMax(passwd)
		target := extractTarget(passwd)
		passwdToCheck := extractPasswd(passwd)

		count := bytes.Count([]byte(passwdToCheck), []byte(target))

		if count >= min && count <= max {
			result++
		}

	}

	return result
}

func validateByIndex(input []string) int {
	var result int

	for _, passwd := range input {
		firstPos, secondPos := extractMinMax(passwd)
		target := extractTarget(passwd)
		passwdToCheck := extractPasswd(passwd)
		passwdToArray := strings.Split(passwdToCheck, "")

		if passwdToArray[firstPos-1] == target && passwdToArray[secondPos-1] != target {
			result++
		} else if passwdToArray[firstPos-1] != target && passwdToArray[secondPos-1] == target {
			result++
		}
	}

	return result
}

func extractMinMax(passwd string) (int, int) {

	splitted := strings.Split(passwd, " ")
	splitWithDash := strings.Split(splitted[0], "-")

	min, _ := strconv.Atoi(splitWithDash[0])
	max, _ := strconv.Atoi(splitWithDash[1])

	return min, max
}

func extractTarget(passwd string) string {
	splitted := strings.Split(passwd, " ")[1]

	return strings.Split(splitted, ":")[0]
}

func extractPasswd(passwd string) string {
	splitted := strings.Split(passwd, " ")
	return splitted[len(splitted)-1]
}
