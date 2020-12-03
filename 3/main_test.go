package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInputFail(t *testing.T) {
	input := loadInput("testi.txt")
	assert.Equal(t, 0, len(input))
}

func TestLoadInputSuccess(t *testing.T) {
	input := loadInput("test.txt")
	assert.Equal(t, 11, len(input))
}

func TestCountThreeSuccess(t *testing.T) {
	input := loadInput("test.txt")
	counted := countThree(input)

	assert.Equal(t, 7, counted)
}
