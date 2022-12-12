package main

import (
	"aoc/day_6/models"
	"fmt"
)

func main() {
	inputFileName := "input.txt"

	device := models.NewDevice().
		WithInputFile(inputFileName)

	packetMarker := models.Marker(models.Packet)
	fmt.Println(device.Marker(packetMarker))
	fmt.Println("-------------")

	messageMarker := models.Marker(models.Message)
	fmt.Println(device.Marker(messageMarker))
	fmt.Println("-------------")
}
