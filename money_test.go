package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{Amount: 5}
	assert.EqualValues(t, five.Amount, 5)
}
