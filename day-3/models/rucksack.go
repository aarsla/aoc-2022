package models

import (
	"errors"
	"fmt"
	"strings"
)

type Rucksack struct {
	Input string
	One   string
	Two   string
}

func (r Rucksack) FromInput(i string) Rucksack {
	length := len(i) / 2
	firstCompartment := i[0:length]
	secondCompartment := i[length:]

	return Rucksack{
		Input: i,
		One:   firstCompartment,
		Two:   secondCompartment,
	}
}

func (r Rucksack) Priority() int {
	letter, _ := r.duplicate()
	return Priority[letter]
}

func (r Rucksack) duplicate() (string, error) {
	for _, ch := range r.One {
		letter := string(ch)
		duplicate := strings.Contains(r.Two, letter)

		if duplicate {
			return letter, nil
		}
	}

	return "", errors.New("no duplicates")
}

func (r Rucksack) String() string {
	return fmt.Sprintf(
		"%v:, %v (%v) - %v (%v) \n",
		r.Input,
		r.One, len(r.One),
		r.Two, len(r.Two))
}
