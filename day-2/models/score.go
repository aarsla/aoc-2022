package models

var ShapeScore = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var GameScore = map[string]int{
	"scissors rock":  0,
	"rock paper":     0,
	"paper scissors": 0,

	"rock rock":         3,
	"scissors scissors": 3,
	"paper paper":       3,

	"rock scissors":  6,
	"paper rock":     6,
	"scissors paper": 6,
}
