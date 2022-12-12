package models

import "fmt"

type Carrier interface {
	TotalCalories() int
}

type Elf struct {
	Carrier
	Name   string
	Snacks []int
}

func (e Elf) TotalCalories(...Carrier) int {
	arrSum := 0

	for _, a := range e.Snacks {
		arrSum = arrSum + a
	}

	return arrSum
}

func (e Elf) String() string {
	return fmt.Sprintf("%s: \t %d \n", e.Name, e.TotalCalories())
}

type ByTotalCalories []Elf

func (a ByTotalCalories) Len() int           { return len(a) }
func (a ByTotalCalories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTotalCalories) Less(i, j int) bool { return a[i].TotalCalories() < a[j].TotalCalories() }
