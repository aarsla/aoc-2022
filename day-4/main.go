package main

import (
	"aoc/day_4/models"
	"aoc/reader"
	"fmt"
)

func main() {
	scanner := reader.ReadFile("./input.txt")

	var intersectionCounter int
	var overlapCounter int

	for scanner.Scan() {
		elf := models.NewElf().WithSections(scanner.Text())

		if elf.Intersects() {
			intersectionCounter++
		}

		if elf.Overlaps() {
			overlapCounter++
		}
	}

	fmt.Println(intersectionCounter)
	fmt.Println(overlapCounter)
	fmt.Println("-------------")
}
