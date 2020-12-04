package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	testInput := loadFromFile("input.txt")
	validPassword := countValidPassport(testInput)
	fmt.Println(validPassword)
}

func loadFromFile(file string) []string {
	var input []string

	raws, err := ioutil.ReadFile(file)
	if err != nil {
		return input
	}

	input = strings.Split(string(raws), "\n\n")

	return input
}

func countValidPassport(input []string) int {
	var validPassport int

	for _, passport := range input {
		byr := bytes.Contains([]byte(passport), []byte("byr"))
		iyr := bytes.Contains([]byte(passport), []byte("iyr"))
		eyr := bytes.Contains([]byte(passport), []byte("eyr"))
		hgt := bytes.Contains([]byte(passport), []byte("hgt"))
		hcl := bytes.Contains([]byte(passport), []byte("hcl"))
		ecl := bytes.Contains([]byte(passport), []byte("ecl"))
		pid := bytes.Contains([]byte(passport), []byte("pid"))

		if byr && iyr && eyr && hgt && hcl && ecl && pid {
			validPassport++
		}
	}

	return validPassport
}
