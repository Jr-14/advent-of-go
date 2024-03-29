package cleanup

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	low  int
	high int
}

func isFullyOverlapping(firstSection Section, secondSection Section) bool {
	if firstSection.low >= secondSection.low && firstSection.high <= secondSection.high {
		return true
	}
	if secondSection.low >= firstSection.low && secondSection.high <= firstSection.high {
		return true
	}
	return false
}

func isPartiallyOverlapping(firstSection Section, secondSection Section) bool {
	if firstSection.high < secondSection.low {
		return false
	}
	if firstSection.low > secondSection.high {
		return false
	}
	return true
}

func parseSection(section string) Section {
	substring := strings.Split(section, "-")
	low, err := strconv.Atoi(substring[0])
	if err != nil {
		panic(err)
	}
	high, err := strconv.Atoi(substring[1])
	if err != nil {
		panic(err)
	}
	return Section{low, high}
}

func parseAssignmentPairs(assignments string) (string, string) {
	substring := strings.Split(assignments, ",")
	return substring[0], substring[1]
}

func retrieveAssignmentPairs(assignments string) (Section, Section) {
	assignmentOne, assignmentTwo := parseAssignmentPairs(assignments)
	elfSectionOne := parseSection(assignmentOne)
	elfSectionTwo := parseSection(assignmentTwo)
	return elfSectionOne, elfSectionTwo
}

func CountFullyOverlappingAssignmentsFromFile(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return CountFullOverlappingAssignments(file)
}

func CountFullOverlappingAssignments(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		elfAssignmentOne, elfAssignmentTwo := retrieveAssignmentPairs(text)
		if isFullyOverlapping(elfAssignmentOne, elfAssignmentTwo) {
			sum += 1
		}
	}
	return sum
}

func CountPartiallyOverlappingAssignmentsFromFile(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return CountPartialOverlappingAssignments(file)
}

func CountPartialOverlappingAssignments(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		elfAssignmentOne, elfAssignmentTwo := retrieveAssignmentPairs(text)
		if isPartiallyOverlapping(elfAssignmentOne, elfAssignmentTwo) {
			sum += 1
		}
	}
	return sum
}
