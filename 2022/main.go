package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/tuningtrouble"
)

func main() {
	tuning := tuningtrouble.TuningTrouble("tuningtrouble/input.txt")
	fmt.Printf("Tuning, %d\n", tuning)
}
