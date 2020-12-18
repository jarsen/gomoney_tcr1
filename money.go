package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Money struct {
	amount   uint64
	currency string
}

type Dollar struct {
	*Money
}

type Franc struct {
	*Money
}

func NewDollar(amount uint64) *Money {
	return &Money{amount: amount, currency: "USD"}
}

func NewFranc(amount uint64) *Money {
	return &Money{amount: amount, currency: "CHF"}
}

func (m *Money) Times(multiplier uint64) *Money {
	return &Money{amount: m.amount * multiplier, currency: m.currency}
}
