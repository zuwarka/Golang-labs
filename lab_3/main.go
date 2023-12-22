package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Лабораторная работа №3, задание №3: В потоке символов сосчитать среднюю длину слова.
*/

const (
	IN  = 1
	OUT = 0
)

func countAverageFromReader() float64 {
	lttr := 0
	wrd := 0
	state := OUT

	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}

		if c == ' ' || c == '.' || c == '\n' || c == ',' || c == '\t' || c == '!' || c == '?' {
			state = OUT
		} else {
			if state == OUT {
				wrd++
				lttr++
				state = IN
			} else {
				lttr++
			}
		}
	}

	return float64(lttr) / float64(wrd)
}

func main() {
	var avg float64

	avg = countAverageFromReader()
	fmt.Printf("The average sum is %g", avg)
}
