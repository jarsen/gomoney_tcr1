package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, *five.Times(2), Dollar{Amount: 10})
	assert.Equal(t, *five.Times(3), Dollar{Amount: 15})
}
