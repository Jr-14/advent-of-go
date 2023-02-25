package calorie

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func BufferReadPrintCalorier(fileName string) {
	// Get file size
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileInformation, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The file is %d bytes long", fileInformation.Size())

	scanner := bufio.NewScanner(file)
	newLineCounter := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			newLineCounter += 1
			fmt.Printf("Hello I am the %d newline\n", newLineCounter)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
