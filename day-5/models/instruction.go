package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Input  string
	Pieces int
	From   int
	To     int
}

func NewInstruction(input string) Instruction {
	instructions := strings.Split(input, " ")
	pieces, _ := strconv.Atoi(instructions[1])
	from, _ := strconv.Atoi(instructions[3])
	to, _ := strconv.Atoi(instructions[5])

	return Instruction{
		Input:  input,
		Pieces: pieces,
		From:   from,
		To:     to,
	}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s: \n", i.Input)
}
