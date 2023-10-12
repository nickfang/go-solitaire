package main

import "fmt"

type stacks [][]card

func newStacks() stacks {
	stacks := stacks{}
	for i := 0; i < 4; i++ {
		stacks = append(stacks, []card{})
	}
	return stacks
}

func (s stacks) display() {
	fmt.Println(s)
}
