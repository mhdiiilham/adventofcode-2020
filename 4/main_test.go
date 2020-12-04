package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInputSuccess(t *testing.T) {
	input := loadFromFile("test.txt")

	assert.Equal(t, 4, len(input), "Lenght of input should be equal to 4")
}

func TestLoadInputFail(t *testing.T) {
	input := loadFromFile("testt.txt")

	assert.Equal(t, 0, len(input), "Lenght of input should be equal to 0")
}

func TestCounValidPasswordSuccess(t *testing.T) {
	input := loadFromFile("test.txt")
	validPassport := countValidPassport(input)

	assert.Equal(t, 2, validPassport, "Number of valid passport should be 2")
}
