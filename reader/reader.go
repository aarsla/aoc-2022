package reader

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileName string) *bufio.Scanner {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}
