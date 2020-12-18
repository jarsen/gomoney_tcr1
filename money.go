package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Money struct {
	amount uint64
}

type Dollar struct {
	*Money
}

type Franc struct {
	*Money
}

func NewDollar(amount uint64) *Dollar {
	return &Dollar{Money: &Money{amount: amount}}
}

func NewFranc(amount uint64) *Franc {
	return &Franc{Money: &Money{amount: amount}}
}

func (d *Dollar) Times(multiplier uint64) *Dollar {
	return &Dollar{Money: &Money{amount: d.amount * multiplier}}
}

func (d *Franc) Times(multiplier uint64) *Franc {
	return &Franc{Money: &Money{amount: d.amount * multiplier}}
}
