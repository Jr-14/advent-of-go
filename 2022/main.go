package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/supplystacks"
)

func main() {
	// const example = "[Q] [J]                         [H]"
	crates := supplystacks.SupplyStack("supplystacks/input.txt")
	fmt.Printf("The crates on top of the stacks are %s", crates)
	fmt.Println("")
}
