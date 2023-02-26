package rps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RockPaperScissors struct {
	choice string
	points int
}

var Rock = RockPaperScissors{"Rock", 1}
var Paper = RockPaperScissors{"Paper", 2}
var Scissors = RockPaperScissors{"Scissors", 3}

func TestHelloWorld() {
	fmt.Println("Hello World!")
}

func parse(str string) (string, string) {
	strArray := strings.Split(str, " ")
	return strArray[0], strArray[1]
}

func mapChoice(rpsChoice string) RockPaperScissors {
	switch rpsChoice {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic("oh no")
	}
}

func determineMyChoice(opponentChoice RockPaperScissors, winningCondition string) RockPaperScissors {
	switch winningCondition {
	case "lose":
		switch opponentChoice {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		default:
			panic("oops")
		}
	case "win":
		switch opponentChoice {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		default:
			panic("oops")
		}
	case "draw":
		return opponentChoice
	default:
		panic("oops")
	}
}

func determineStrategy(encryptedStrategy string) string {
	switch encryptedStrategy {
	case "X":
		return "lose"
	case "Y":
		return "draw"
	case "Z":
		return "win"
	default:
		panic("I can't win")
	}
}

func calculateRoundScorePartTwo(opponentChoice string, myChoice string) int {
	opponentPlay := mapChoice(opponentChoice)
	strategy := determineStrategy(myChoice)
	toPlay := determineMyChoice(opponentPlay, strategy)
	switch strategy {
	case "lose":
		return 0 + toPlay.points
	case "draw":
		return 3 + toPlay.points
	case "win":
		return 6 + toPlay.points
	default:
		panic("no points")
	}
}

func calculateRoundScore(opponentChoice string, myChoice string) int {
	opponentPlay := mapChoice(opponentChoice)
	myPlay := mapChoice(myChoice)
	if opponentPlay == myPlay {
		return 3 + myPlay.points
	}
	if myPlay.choice == "Rock" && opponentPlay.choice == "Scissors" {
		return 6 + myPlay.points
	}
	if myPlay.choice == "Paper" && opponentPlay.choice == "Rock" {
		return 6 + myPlay.points
	}
	if myPlay.choice == "Scissors" && opponentPlay.choice == "Paper" {
		return 6 + myPlay.points
	}
	return 0 + myPlay.points
}

func PlayRockPaperScissors(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		opponentChoice, myChoice := parse(text)
		totalScore += calculateRoundScorePartTwo(opponentChoice, myChoice)
	}
	return totalScore
}
