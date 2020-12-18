package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Dollar struct {
	amount uint64
}

type Franc struct {
	amount uint64
}

func NewDollar(amount uint64) *Dollar {
	return &Dollar{amount: amount}
}

func NewFranc(amount uint64) *Franc {
	return &Franc{amount: amount}
}

func (d *Dollar) Times(multiplier uint64) *Dollar {
	return &Dollar{amount: d.amount * multiplier}
}

func (d *Franc) Times(multiplier uint64) *Franc {
	return &Franc{amount: d.amount * multiplier}
}
