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
	five := NewFranc(5)
	assert.Equal(t, five.Times(2), NewFranc(10))
	assert.Equal(t, five.Times(3), NewFranc(15))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.NotEqual(t, NewDollar(5), NewDollar(6))
}
