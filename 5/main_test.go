package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeatID(t *testing.T) {
	boardingPass := "FBFBBFFRLR"
	seatID := getSeatID(boardingPass)
	assert.Equal(t, seatID, 357)
}

func TestGetMySeatID(t *testing.T) {
	_, listSeatID := getHighestSeatID("input.txt")
	mySeatID := getMySeatID(listSeatID)
	assert.Equal(t, 615, mySeatID)
}

func TestGetHighestSeatID(t *testing.T) {
	highestSeatID, listSeatID := getHighestSeatID("input.txt")
	assert.Equal(t, 953, highestSeatID)
	assert.Equal(t, 908, len(listSeatID), "Sum of boarding pass should be 908")
}

func TestFrontFormula(t *testing.T) {
	currentRange := [2]int{32, 63}
	FRange := frontFormula(currentRange)
	fmt.Println(FRange)
	assert.Equal(t, 32, FRange[0], "First Index should be equal to 32")
	assert.Equal(t, 47, FRange[1], "Last Index should be equal to 47")
}

func TestBackFormula(t *testing.T) {
	currentRange := [2]int{32, 47}
	FRange := backFormula(currentRange)
	assert.Equal(t, 40, FRange[0], "First Index should be equal to 40")
	assert.Equal(t, 47, FRange[1], "Last Index should be equal to 47")
}

func TestCalculateColumn(t *testing.T) {
	row := calculateCol("FBFBBFFRLR")
	assert.Equal(t, 5, row)
}
