package models

type Game struct {
	playerMove   Move
	opponentMove Move
}

func (g Game) FromMoveWithInstructions(m Move) Game {
	return g.FromMoves(m, m.CounterMove())
}

func (g Game) FromMoves(m Move, c Move) Game {
	return Game{playerMove: m, opponentMove: c}
}

func (g Game) PlayerScore() int {
	return g.playerMove.shapeScore() +
		GameScore[g.playerMove.Name+" "+g.opponentMove.Name]
}

func (g Game) OpponentScore() int {
	return g.opponentMove.shapeScore() +
		GameScore[g.opponentMove.Name+" "+g.playerMove.Name]
}

func (g Game) Winner() string {
	if g.PlayerScore() > g.OpponentScore() {
		return "player"
	}

	return "opponent"
}
