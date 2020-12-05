package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func yearValidate(byr string, digit, min, max int) bool {
	if len(byr) != digit {
		return false
	}

	byrToInt, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}

	if byrToInt < min || byrToInt > max {
		return false
	}

	return true
}

func heightValidate(height string) bool {
	isCM, _ := regexp.MatchString(`cm`, height)
	isIN, _ := regexp.MatchString(`in`, height)

	if isCM {
		removeCM := strings.ReplaceAll(height, "cm", "")
		value, err := strconv.Atoi(removeCM)
		if err != nil {
			return false
		}
		if value < 150 || value > 193 {
			return false
		}
	}

	if isIN {
		removeIN := strings.ReplaceAll(height, "in", "")
		value, err := strconv.Atoi(removeIN)
		if err != nil {
			return false
		}
		if value < 59 || value > 76 {
			return false
		}
	}

	return true
}

func validateColor(clr string) bool {
	ok, err := regexp.MatchString(`^#(?:[0-9a-fA-F]{3}){1,2}$`, clr)
	if err != nil {
		return false
	}

	return ok
}

func validateEcl(ecl string) bool {
	switch ecl {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func validatePassportID(pid string) bool {
	pidArray := strings.Split(pid, "")

	if len(pidArray) != 9 || pidArray[0] != "0" {
		return false
	}

	return true
}
