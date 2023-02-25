package main

import (
	"fmt"

	"github.com/Jr-14/advent-of-go/2022/calorie"
)

func main() {
	maxCalorie := calorie.CountCalories("calorie/1-input.txt")
	fmt.Printf("Max Calorie is %d.\n", maxCalorie)
}
