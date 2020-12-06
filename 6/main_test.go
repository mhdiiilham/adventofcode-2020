package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputFromFile(t *testing.T) {
	input := readFromFile("test.txt")
	assert.Equal(t, 5, len(input), "Length of input should be equal to five")
	assert.Equal(t, "abc", input[0], "The first index should be equal to 'abc'")
	assert.Equal(t, "b", input[len(input)-1], "The last index should be equal to 'b'")
}

func TestReadFromEmptyFile(t *testing.T) {
	input := readFromFile("empty.txt")
	assert.Equal(t, 0, len(input), "Length of input should be equal to zero")
}

func TestCountEachGroup(t *testing.T) {
	input := readFromFile("test.txt")
	yesAnsert := countYesAnswer(input[2])
	assert.Equal(t, 3, yesAnsert, "The sum of group 3 that answer yes should be eqaual to 3")
}

func TestCountAllAnswer(t *testing.T) {
	input := readFromFile("test.txt")
	yes := countYes(input)
	assert.Equal(t, 11, yes, "Sum of yes from the input groups should be equal to 11")
}
