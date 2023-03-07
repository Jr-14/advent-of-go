package tuningtrouble

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func TuningTrouble(fileName string) int {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	stride := 1024
	reader := bufio.NewReader(file)
	buffer := make([]byte, 0, stride)
	for {
		n, err := io.ReadFull(reader, buffer[:cap(buffer)])
		buffer = buffer[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}
		signal := string(buffer[:n])
		fmt.Println(signal)
		fmt.Println("")
	}
	return 0
}
