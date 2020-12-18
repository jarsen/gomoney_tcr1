package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Dollar struct {
	Amount uint64
}

func (d *Dollar) times(multiplier uint64) *Dollar {
	return d
}
