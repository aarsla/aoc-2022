package main

import (
	"aoc/day_5/models"
	"aoc/reader"
	"fmt"
)

func main() {
	crane := move()
	fmt.Println("-------------")
	fmt.Println(crane.TopOfStack())

	crane9001 := move9001()
	fmt.Println("-------------")
	fmt.Println(crane9001.TopOfStack())
}

func move() models.Crane {
	scanner := reader.ReadFile("./input.txt")
	crane := models.NewCrane()

	for scanner.Scan() {
		instruction := models.NewInstruction(scanner.Text())
		crane.Move(instruction)
	}

	return crane
}

func move9001() models.Crane {
	scanner := reader.ReadFile("./input.txt")
	crane := models.NewCrane()

	for scanner.Scan() {
		instruction := models.NewInstruction(scanner.Text())
		crane.Move9001(instruction)
	}

	return crane
}
