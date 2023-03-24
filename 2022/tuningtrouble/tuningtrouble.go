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
	nPos := 0
	markerPos := 0
	signalPrefix := ""
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

		markerPos = processSignal(signal, signalPrefix, n)
		if markerPos > 0 {
			str := string(signal[markerPos-3:markerPos + 1])
			fmt.Printf("Signal: %s\n", str)
			return nPos + markerPos + 2
		}
		nPos += n
		signalPrefix = string(signal[n-4:n])
	}
	return 0
}

func processSignal(signal, signalPrefix string, n int) (int) {
	m := make(map[string]int)

	if signalPrefix == "" {
		for i := 0; i < 4; i++ {
			character := string(signal[i])
			_, ok := m[character]
			if ok {
				m[character] += 1
			} else {
				m[character] = 1
			}
		}
	} else {
		for i := 0; i < 4; i++ {
			character := string(signalPrefix[i])
			_, ok := m[character]
			if ok {
				m[character] += 1
			} else {
				m[character] = 1
			}
		}
	}

	j := 0
	earlyReturn := true
	for _, v := range m {
		if v > 1 {
			earlyReturn = false
			break
		}
		j += 1
	}
	if earlyReturn {
		return j + 1
	}

	for i := 4; i < n; i++ {
		charToRemove := string(signal[i-4])
		m[charToRemove] -= 1

		character := string(signal[i])
		_, ok := m[character]
		if ok {
			m[character] += 1
		} else {
			m[character] = 1
		}
		for _, v := range m {
			if v > 1 {
				earlyReturn = false
				break
			}
		}
		if earlyReturn {
			return i + 1
		}
	}
	// fmt.Printf("Signal length %d, and n: %d", len(signal), n)
	return 0
}
