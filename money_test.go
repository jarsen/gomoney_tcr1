package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	assert.Equal(t, NewDollar(5).Times(2), NewDollar(10))
	assert.Equal(t, NewDollar(5).Times(3), NewDollar(15))
	assert.Equal(t, NewFranc(5).Times(2), NewFranc(10))
	assert.Equal(t, NewFranc(5).Times(3), NewFranc(15))
}

func TestAddition(t *testing.T) {
	assert.Equal(t, NewDollar(5).Plus(NewDollar(5)), NewDollar(10))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.NotEqual(t, NewDollar(5), NewDollar(6))
	assert.Equal(t, NewFranc(5), NewFranc(5))
	assert.NotEqual(t, NewFranc(5), NewFranc(6))
	assert.NotEqual(t, NewDollar(5), NewFranc(5))
}
