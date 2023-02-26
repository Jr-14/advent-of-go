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

func calculateRuckSackPriorityPartTwo(fSack, sSack, tSack string) int {
	fSackPriority := [53]int{0}
	sSackPriority := [53]int{0}
	tSackPriority := [53]int{0}
	var commonItemPriority int
	for i := 0; i < len(fSack); i++ {
		character := string(fSack[i])
		priority := calculateCharacterPriority(character)
		if fSackPriority[priority] < 1 {
			fSackPriority[priority]++
		}
	}

	for i := 0; i < len(sSack); i++ {
		character := string(sSack[i])
		priority := calculateCharacterPriority(character)
		if sSackPriority[priority] < 1 {
			sSackPriority[priority]++
		}
	}

	for i := 0; i < len(tSack); i++ {
		character := string(tSack[i])
		priority := calculateCharacterPriority(character)
		if tSackPriority[priority] < 1 {
			tSackPriority[priority]++
		}
	}

	for i := 0; i < len(fSackPriority); i++ {
		fp := fSackPriority[i]
		sp := sSackPriority[i]
		tp := tSackPriority[i]
		if fp+sp+tp == 3 {
			commonItemPriority = i
		}
	}
	return commonItemPriority
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

func CalculateRuckSackPartTwo(fileName string) (sum int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineCounter := 0
	elfGroup := make([]string, 3)
	for scanner.Scan() {
		rucksack := scanner.Text()
		elfGroup[lineCounter] = rucksack
		lineCounter += 1
		if lineCounter == 3 {
			sum += calculateRuckSackPriorityPartTwo(elfGroup[0], elfGroup[1], elfGroup[2])
			lineCounter = 0
			elfGroup = make([]string, 3)
		}
	}
	return
}
