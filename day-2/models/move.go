package models

import "fmt"

type Move struct {
	Name        string
	Instruction string
}

func (m Move) FromString(c string) Move {
	input := Moves[c]
	return Move{Name: input}
}

func (m Move) WithInstructions(i string) Move {
	m.Instruction = Instructions[i]
	return m
}

func (m Move) CounterMove() Move {
	return Move{Name: Rules[m.Name+" "+m.Instruction]}
}

func (m Move) String() string {
	return fmt.Sprintf("Name: %s, %s \n", m.Name, m.Instruction)
}

func (m Move) shapeScore() int {
	return ShapeScore[m.Name]
}
