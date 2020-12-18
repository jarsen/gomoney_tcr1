package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{Amount: 5}
	assert.Equal(t, *five.Times(2), Dollar{Amount: 10})
	assert.EqualValues(t, five.Times(3).Amount, 15)
}
