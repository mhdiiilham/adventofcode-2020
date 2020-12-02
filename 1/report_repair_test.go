package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessConvertTXTFile(t *testing.T) {
	input := readInputFile("input.txt")
	assert.Equal(t, 200, len(*input), "Input length should be equal to 200")
}

func TestSuccessConvertTXTFileReturningNil(t *testing.T) {
	input := readInputFile("inputtt.txt")
	assert.Nil(t, input, "Input should be nil")
}

func TestSuccessGetResult(t *testing.T) {
	input := []int{1010, 6543, 7346, 334, 1010, 2934}
	reportResult := reportRepair(input)

	assert.Equal(t, 1020100, reportResult, "Report Repair should be equal to 1020100")
}

func TestNotFound(t *testing.T) {
	input := []int{1010, 6543, 7346, 334, 2934}
	reportResult := reportRepair(input)
	assert.Equal(t, 0, reportResult, "Report Repair should be equal to 0 if none sum is equal to 2020")
}

func TestSuccessGetResult3Sum(t *testing.T) {
	input := readInputFile("input.txt")
	reportResult := reportRepair3Sum(*input)

	assert.Equal(t, 262738554, reportResult, "Report Repair should be equal to 262738554")
}
