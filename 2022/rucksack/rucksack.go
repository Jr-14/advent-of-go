package rucksack

import (
	"bufio"
	"os"
)

func calculateCharacterPriority(character string) int {
	asciiRep := int(character[0])
	// lowercase characters
	var priority int
	if asciiRep >= 97 && asciiRep <= 122 {
		priority = asciiRep - 96
	} else {
		priority = asciiRep - 38
	}
	return priority
}

func splitToCompartments(items string) (string, string) {
	return items[:len(items)/2], items[len(items)/2:]
}

func calculateRucksackPriority(rucksack string) int {
	firstCompartment, secondCompartment := splitToCompartments(rucksack)
	firstCompartmentArray := [53]int{0}
	var priority int
	for i := 0; i < len(firstCompartment); i++ {
		fcCharacter := string(firstCompartment[i])
		fcPriority := calculateCharacterPriority(fcCharacter)
		firstCompartmentArray[fcPriority]++
	}
	for i := 0; i < len(secondCompartment); i++ {
		scCharacter := string(secondCompartment[i])
		scPriority := calculateCharacterPriority(scCharacter)
		if firstCompartmentArray[scPriority] > 0 {
			priority = scPriority
			break
		}
	}
	return priority
}

func CalculateRuckSack(fileName string) (sum int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rucksack := scanner.Text()
		sum += calculateRucksackPriority(rucksack)
	}
	return
}
