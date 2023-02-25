package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/rps"
)

func main() {
	score := rps.PlayRockPaperScissors("rps/input.txt")
	fmt.Printf("My Max Score is %d.", score)
}
