package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/cleanup"
)

func main() {
	overlapping := cleanup.CountPartialOverlappingAssignments("cleanup/input.txt")
	fmt.Printf("Sum of any overlapping assignment is %d.", overlapping)
}
