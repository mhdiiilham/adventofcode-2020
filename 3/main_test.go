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

func TestCountTreeSuccess(t *testing.T) {
	input := loadInput("test.txt")
	counted := countTree(input, 3, 1)

	assert.Equal(t, 7, counted)
}

func TestCountTreePartTwoSuccess(t *testing.T) {
	input := loadInput("test.txt")
	first := countTree(input, 1, 1)
	second := countTree(input, 3, 1)
	third := countTree(input, 5, 1)
	fourth := countTree(input, 7, 1)
	fifth := countTree(input, 1, 2)

	assert.Equal(t, 336, first*second*third*fourth*fifth)
}
