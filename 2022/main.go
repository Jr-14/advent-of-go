package main

import (
	// "github.com/Jr-14/advent-of-go/2022/rps"
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/rucksack"
)

func main() {
	// score := rps.PlayRockPaperScissors("rps/input.txt")
	// fmt.Printf("My Max Score is %d.", score)
	priorities := rucksack.CalculateRuckSack("rucksack/input.txt")
	fmt.Printf("Sum of priorities in Rucksack is %d.", priorities)
}
