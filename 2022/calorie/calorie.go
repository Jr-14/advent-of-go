package calorie

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCalories(fileName string) [][]int {
	// Get file size
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileInformation, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The file is %d bytes long\n", fileInformation.Size())

	scanner := bufio.NewScanner(file)

	calorieArray := make([]int, 0)
	calorieCountArray := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			val, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			calorieArray = append(calorieArray, val)
		} else {
			// Push the latest array of number to the entire
			calorieCountArray = append(calorieCountArray, calorieArray)
			calorieArray = make([]int, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Calorie Array: %v", calorieCountArray)
	return calorieCountArray
}

func sumCalories(calorieArray []int) (calories int) {
	for _, e := range calorieArray {
		calories += e
	}
	return
}

func CountMaxCalorie(fileName string) int {
	calories := readCalories(fileName)
	maxCalorie := 0
	for _, calorieArr := range calories {
		newCalorie := sumCalories(calorieArr)
		if newCalorie >= maxCalorie {
			maxCalorie = newCalorie
		}
	}
	return maxCalorie
}
