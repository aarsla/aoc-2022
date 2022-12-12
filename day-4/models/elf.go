package models

import (
	"fmt"
	"github.com/goombaio/namegenerator"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Elf struct {
	Name         string
	Slice        []int
	SectionOne   []int
	SectionTwo   []int
	Intersection []int
}

func NewElf() Elf {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}

	return Elf{
		Slice: arr,
		Name:  nameGenerator.Generate(),
	}
}

func (e Elf) WithSections(i string) Elf {
	sections := strings.Split(i, ",")
	sectionOne := strings.Split(sections[0], "-")
	sectionTwo := strings.Split(sections[1], "-")

	sectionOneStart, _ := strconv.Atoi(sectionOne[0])
	sectionOneEnd, _ := strconv.Atoi(sectionOne[1])

	sectionTwoStart, _ := strconv.Atoi(sectionTwo[0])
	sectionTwoEnd, _ := strconv.Atoi(sectionTwo[1])

	e.SectionOne = e.Slice[sectionOneStart : sectionOneEnd+1]
	e.SectionTwo = e.Slice[sectionTwoStart : sectionTwoEnd+1]

	return e.intersection()
}

func (e Elf) intersection() Elf {
	var c []int
	m := make(map[int]bool)

	for _, item := range e.SectionOne {
		m[item] = true
	}

	for _, item := range e.SectionTwo {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}

	e.Intersection = c
	return e
}

func (e Elf) Intersects() bool {
	if reflect.DeepEqual(e.Intersection, e.SectionOne) || reflect.DeepEqual(e.Intersection, e.SectionTwo) {
		return true
	}

	return false
}

func (e Elf) Overlaps() bool {
	if len(e.Intersection) > 0 {
		return true
	}

	return false
}

func (e Elf) String() string {
	return fmt.Sprintf("%s: \n", e.Name)
}
