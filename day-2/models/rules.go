package models

var Rules = map[string]string{
	"rock loose":     "scissors",
	"rock draw":      "rock",
	"rock win":       "paper",
	"paper loose":    "rock",
	"paper draw":     "paper",
	"paper win":      "scissors",
	"scissors loose": "paper",
	"scissors draw":  "scissors",
	"scissors win":   "rock",
}
