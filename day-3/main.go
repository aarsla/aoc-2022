package main

import (
	"aoc/day_3/models"
	"aoc/reader"
	"fmt"
)

func main() {
	scanner := reader.ReadFile("./input.txt")

	counter := 1
	var rucksacks []models.Rucksack
	var priorityScore int

	for scanner.Scan() {
		rucksack := new(models.Rucksack).FromInput(scanner.Text())
		rucksacks = append(rucksacks, rucksack)

		if counter%3 == 0 {
			group := new(models.Group).FromRucksacks(rucksacks)
			priorityScore += group.Priority()

			rucksacks = nil
		}

		//fmt.Println(rucksack)
		//priorityScore += rucksack.Priority()
		counter++
	}

	fmt.Println(priorityScore)
	fmt.Println("-------------")
}
