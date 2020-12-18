package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := Dollar{Amount: 5}
	ten := five.Times(2)
	assert.EqualValues(t, ten.Amount, 10)
}
