package main

import (
	"aoc/day_1/models"
	"aoc/reader"
	"fmt"
	"github.com/goombaio/namegenerator"
	"log"
	"sort"
	"strconv"
	"time"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	scanner := reader.ReadFile("./input.txt")

	e := models.Elf{Name: nameGenerator.Generate()}
	var elves []models.Elf

	for scanner.Scan() {
		if scanner.Text() == "" {
			//fmt.Println("----------------")
			elves = append(elves, e)
			e = models.Elf{Name: nameGenerator.Generate()}
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			} else {
				//fmt.Println(scanner.Text())
				e.Snacks = append(e.Snacks, calories)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxIndex := 0
	maxCalories := 0

	for index, e := range elves {
		calories := e.TotalCalories()
		//fmt.Printf("Elf at index %v has %v total calories\n", index, calories)

		if calories > maxCalories {
			maxIndex = index
			maxCalories = calories
		}
	}

	sort.Sort(sort.Reverse(models.ByTotalCalories(elves)))
	//sort.Slice(elves, func(i, j int) bool {
	//	return elves[i].Total() > elves[j].Total()
	//})

	topElves := elves[0:3]

	fmt.Println(elves)
	fmt.Printf("Elf at index %v has %v total calories\n", maxIndex, maxCalories)

	fmt.Printf("Top %v elves\n", 3)
	fmt.Println(topElves)

	var sum int
	for _, e := range topElves {
		sum = sum + e.TotalCalories()
	}

	fmt.Printf("Total calories %v\n", sum)
}
