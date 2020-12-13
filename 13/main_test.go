package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnExistInput(t *testing.T) {
	input := readFromFile("notdefine.txt")
	assert.Equal(t, 0, len(input), "Input length should be equal to zero")
}

func TestReadInputSuccess(t *testing.T) {
	input := readFromFile("test.txt")
	assert.Equal(t, 2, len(input), "Input length should be equal to two")
	assert.Equal(t, "939", input[0], "Input index 0 should be equal to 939")
	assert.Equal(t, "7,13,x,x,59,x,31,19", input[1], "Input index 1 should be equal to 7,13,x,x,59,x,31,19")
}

func TestMultipliedDepartAndWait(t *testing.T) {
	busID := 59
	minDepartTime := 939
	multipliedDepartAndWaitTime := multiplyingTime(busID, minDepartTime)
	assert.Equal(t, 295, multipliedDepartAndWaitTime, "The multiplied of minutes waiting and bus id should be equal to 295")
}
