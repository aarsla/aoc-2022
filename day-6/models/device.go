package models

import (
	"aoc/reader"
	"bufio"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Device struct {
	InputFile   string
	InputString string
}

func NewDevice() Device {
	return Device{}
}

func (d Device) WithInputFile(i string) Device {
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), "../"+i)

	scanner := reader.ReadFile(filePath)
	for scanner.Scan() {
		d.InputString += scanner.Text()
	}

	return Device{
		InputFile:   filePath,
		InputString: d.InputString,
	}
}

func (d Device) Marker(markerType Marker) (string, int) {
	scanner := reader.ReadFile(d.InputFile)
	scanner.Split(bufio.ScanRunes)

	var counter int
	var letters []string

	for scanner.Scan() {
		counter++

		letters = append(letters, scanner.Text())

		if hasDuplicates(letters) {
			letters = pop(letters)
		} else {
			if len(letters) == markerType.EnumIndex() {
				break
			}
		}

	}

	return strings.Join(letters, ""), counter
}

func hasDuplicates(s []string) bool {
	frequency := make(map[string]int)

	for _, letter := range s {
		_, exist := frequency[letter]
		if exist {
			frequency[letter] += 1
			if frequency[letter] > 1 {
				return true
			}
		} else {
			frequency[letter] = 1
		}
	}

	return false
}

func pop(s []string) []string {
	return s[1:]
}

func (d Device) String() string {
	return fmt.Sprintf("Input: %s \n", d.InputFile)
}
