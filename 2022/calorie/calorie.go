package calorie

import (
	"bufio"
	"os"
	"strconv"
)

func readCaloriesFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	calorieArray := make([]int, 0)
	totalCalories := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			calorie, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			calorieArray = append(calorieArray, calorie)
		} else {
			totalCalories = append(totalCalories, calorieArray)
			calorieArray = make([]int, 0)
		}
	}
	return totalCalories
}

func sum(array []int) (sum int) {
	for _, element := range array {
		sum += element
	}
	return
}

func CountCalories(fileName string) int {
	elfCalories := readCaloriesFile(fileName)
	maxCalorie := 0
	for _, calories := range elfCalories {
		countedCalorie := sum(calories)
		if countedCalorie >= maxCalorie {
			maxCalorie = countedCalorie
		}
	}
	return maxCalorie
}
