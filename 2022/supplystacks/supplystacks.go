package supplystacks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MAX_STACK_INDEX = 9

func createCrateAndAddToInitialStack(lineStack string, cargoStack *[][]string, rowIndex int) {
	// fmt.Printf("Current line stack: %s", lineStack)
	// fmt.Println("")
	// cargoRow := (*cargoStack)[rowIndex]
	for i := 1; i < len(lineStack); i += 4 {
		character := string(lineStack[i])
		if character != " " {
			crateCargoStack := (*cargoStack)[i/4]
			crateCargoStack = append(crateCargoStack, character)
			(*cargoStack)[i/4] = crateCargoStack
		}
	}
}

func buildInitialStack(file *os.File, scanner *bufio.Scanner) [][]string {
	cargoStack := make([][]string, 9)
	lineStack := make([]string, 0)

	counter := 0
	for scanner.Scan() {
		counter++
		if counter > MAX_STACK_INDEX {
			break
		}
		text := scanner.Text()
		// fmt.Println(text)
		lineStack = append(lineStack, text)
	}
	// fmt.Printf("line Stack:\n %v", lineStack)
	// fmt.Println("")

	row := 0
	for i := len(lineStack) - 2; i >= 0; i-- {
		createCrateAndAddToInitialStack(lineStack[i], &cargoStack, row)
		row++
	}
	// fmt.Printf("%v", cargoStack)
	return cargoStack
}

func move(amount, from, to int, cargoStack *[][]string) {
	auxStack := make([]string, 0)
	fromStack := (*cargoStack)[from-1]
	// fmt.Printf("%v", fromStack)
	for i := 0; i < amount; i++ {
		last := len(fromStack) - 1
		lastCargo := fromStack[last]
		auxStack = append(auxStack, lastCargo)
		fromStack = fromStack[:last]
	}
	(*cargoStack)[from-1] = fromStack
	// fmt.Printf("%v", fromStack)

	toStack := (*cargoStack)[to-1]
	for i := amount - 1; i >= 0; i-- {
		topAuxStack := auxStack[i]
		toStack = append(toStack, topAuxStack)
	}
	(*cargoStack)[to-1] = toStack
}

func parseText(text string) (amount, to, from int) {
	arr := strings.Split(text, " ")
	amount, err := strconv.Atoi(arr[1])
	if err != nil {
		panic(err.Error())
	}

	from, err = strconv.Atoi(arr[3])
	if err != nil {
		panic(err.Error())
	}

	to, err = strconv.Atoi(arr[5])
	if err != nil {
		panic(err.Error())
	}

	return
}

func printLastCargoOnStack(cargoStack *[][]string) {
	for i := 0; i < 9; i++ {
		last := len((*cargoStack)[i]) - 1
		fmt.Printf("%s", (*cargoStack)[i][last])
	}
	fmt.Println("")
}

func SupplyStack(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	cargoStack := buildInitialStack(file, scanner)
	lineCount := 0
	fmt.Printf("%v", cargoStack)
	for scanner.Scan() {
		lineCount += 1
		text := scanner.Text()
		amount, to, from := parseText(text)
		// fmt.Printf("%d, %d, %d", amount, to, from)
		move(amount, from, to, &cargoStack)
		// fmt.Println(text)
		// if lineCount > 2 {
		// 	break
		// }
	}
	printLastCargoOnStack(&cargoStack)
}
