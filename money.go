package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Dollar struct {
	Amount uint64
}

func NewDollar(amount uint64) *Dollar {
	return &Dollar{Amount: amount}
}

func (d *Dollar) Times(multiplier uint64) *Dollar {
	return &Dollar{Amount: d.Amount * multiplier}
}
