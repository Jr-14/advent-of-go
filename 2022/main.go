package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/cleanup"
)

func main() {
	overlapping := cleanup.CountFullOverlappingAssignments("cleanup/input.txt")
	fmt.Printf("Sum of priorities in Rucksack is %d.", overlapping)
}
