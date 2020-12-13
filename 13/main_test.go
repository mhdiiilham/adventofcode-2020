package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnExistInput(t *testing.T) {
	input := readFromFile("notdefine.txt")
	assert.Equal(t, 0, len(input), "Input length should be equal to zero")
}
