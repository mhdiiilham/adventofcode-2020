package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	input := readFromFile("test.txt")

	assert.Equal(t, 9, len(input), "Length of input should be equal to 9")
}
