package main

import (
	"aoc/day_2/models"
	"aoc/reader"
	"fmt"
)

func main() {
	scanner := reader.ReadFile("./input.txt")

	var playerScore int
	var opponentScore int

	for scanner.Scan() {
		firstCharacter := scanner.Text()[0:1]
		secondCharacter := scanner.Text()[2:3]

		firstMove := new(models.Move).
			FromString(firstCharacter).
			WithInstructions(secondCharacter)

		//counterMove := firstMove.CounterMove()

		game := new(models.Game).
			FromMoveWithInstructions(firstMove)

		playerScore += game.PlayerScore()
		opponentScore += game.OpponentScore()

		//fmt.Printf("%v - %v\n", firstCharacter, secondCharacter)
		//fmt.Printf("%v\n", game)
		//fmt.Printf("%v : %v\n", game.PlayerScore(), game.OpponentScore())
		//fmt.Println("-------------")
	}

	fmt.Printf("Player: %v - Opponent: %v\n", playerScore, opponentScore)

	fmt.Println("-------------")

}
