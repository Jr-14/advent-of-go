package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/tuningtrouble"
)

func main() {
	// const example = "[Q] [J]                         [H]"
	signal := tuningtrouble.TuningTrouble("tuningtrouble/input.txt")
	fmt.Printf("The crates on top of the stacks are %d", signal)
	fmt.Println("")
}
