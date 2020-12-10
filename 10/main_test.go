package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	input := readFromFile("test.txt")
	assert.Equal(t, 31, len(input), "Length of test input should be equal to 31")
}
