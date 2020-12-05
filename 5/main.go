package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type rowFunc func(x [2]int) [2]int

func main() {
	boarding := "BFFFBBFRRR"
	highestSeatID := getHighestSeatID("input.txt")
	fmt.Println("SeatID->", getSeatID(boarding))
	fmt.Println("Hihest Seat ID->", highestSeatID)
}

func getHighestSeatID(file string) int {
	seatID := 0
	raws, err := ioutil.ReadFile(file)
	if err != nil {
		return seatID
	}

	splitByLine := strings.Split(string(raws), "\n")
	for _, boarding := range splitByLine {
		currentSeatID := getSeatID(boarding)
		if currentSeatID > seatID {
			seatID = currentSeatID
		}
	}

	return seatID
}

func getSeatID(boarding string) int {
	seatID := 0
	col := calculateCol(boarding)
	row := calculateRow(boarding)
	seatID = (row * 8) + col
	return seatID
}

func calculateCol(boarding string) int {
	col := 0
	colRanges := [2]int{0, 7}
	boarding = boarding[7:10]
	boardings := strings.Split(boarding, "")
	checkLetter := map[string]rowFunc{
		"L": frontFormula,
		"R": backFormula,
	}

	for i, letter := range boardings {
		if i == len(boardings)-1 {
			if letter == "L" {
				return colRanges[0]
			}

			if letter == "R" {
				return colRanges[1]
			}
		}
		colRanges = checkLetter[letter](colRanges)
	}

	return col
}

func calculateRow(boarding string) int {
	rows := 0
	rowRangs := [2]int{0, 127}
	boarding = boarding[0:7]
	boardings := strings.Split(boarding, "")
	checkLetter := map[string]rowFunc{
		"F": frontFormula,
		"B": backFormula,
	}

	for i, letter := range boardings {
		if i == len(boardings)-1 {
			if letter == "F" {
				return rowRangs[0]
			}

			if letter == "B" {
				return rowRangs[1]
			}
		}
		rowRangs = checkLetter[letter](rowRangs)
	}

	return rows
}

func frontFormula(input [2]int) [2]int {
	rowRanges := input
	min := input[0]
	max := input[1]
	rowRanges[1] = ((max - min) / 2) + min
	return rowRanges
}

func backFormula(input [2]int) [2]int {
	rowRanges := input
	min := input[0]
	max := input[1]
	rowRanges[0] = (((max - min) / 2) + 1) + min
	return rowRanges
}
