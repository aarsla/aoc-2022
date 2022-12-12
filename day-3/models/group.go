package models

import (
	"fmt"
	"strings"
)

type Group struct {
	Rucksacks []Rucksack
	Badge     string
}

func (g Group) FromRucksacks(r []Rucksack) Group {
	return Group{
		Rucksacks: r,
	}.badge()
}

func (g Group) badge() Group {
	var firstMatch string
	var secondMatch string

	for _, ch := range g.Rucksacks[0].Input {
		firstMatch = match(string(ch), g.Rucksacks[1].Input)
		if firstMatch != "" {
			secondMatch = match(firstMatch, g.Rucksacks[2].Input)
			if secondMatch != "" {
				return Group{
					Rucksacks: g.Rucksacks,
					Badge:     secondMatch,
				}
			}
		} else {
			continue
		}
	}

	return g
}

func match(letter string, input string) string {
	match := strings.Contains(input, letter)

	if match {
		return letter
	}

	return ""
}

func (g Group) Priority() int {
	return Priority[g.Badge]
}

func (g Group) String() string {
	return fmt.Sprintf("Badge: %v \n", g.Badge)
}
