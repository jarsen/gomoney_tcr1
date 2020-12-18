package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, five.Times(2), NewDollar(10))
	assert.Equal(t, five.Times(3), NewDollar(15))
}

func TestFrancMultiplication(t *testing.T) {
	_ = NewFranc(5)
}
