package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessLoadFile(t *testing.T) {
	input := loadInputFromFile("input.txt")

	assert.Equal(t, 1000, len(input))
}

func TestSuccessLoadFileFromUnExist(t *testing.T) {
	input := loadInputFromFile("inputtt.txt")

	assert.Equal(t, 0, len(input))
}

func TestGetMinMax(t *testing.T) {
	min, max := extractMinMax("16-17 t: tttttttttttttttmt")

	assert.Equal(t, 16, min, "Minimal Should be 16")
	assert.Equal(t, 17, max, "Maximal Should be 17")
}

func TestGetTarget(t *testing.T) {
	target := extractTarget("16-17 t: tttttttttttttttmt")

	assert.Equal(t, "t", target, "Target should be equal to 't'")
}

func TestGetStringPasswd(t *testing.T) {
	passwd := extractPasswd("16-17 t: tttttttttttttttmt")

	assert.Equal(t, "tttttttttttttttmt", passwd)
}

func TestInputIsEmptySliceOfString(t *testing.T) {
	input := []string{}
	total := countInput(input)

	assert.Equal(t, 0, total)
}

func TestSuccessCounting(t *testing.T) {
	input := loadInputFromFile("test.txt")
	total := countInput(input)

	assert.Equal(t, 2, total)
}
