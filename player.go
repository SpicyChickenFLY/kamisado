package kamisado

const (
	playerNum   = 2
	firstPlayer = playerWhite
)

const (
	nonPlayer = -1 + iota
	playerWhite
	playerBlack
)

const (
	playerWhiteInitRow = 0
	playerBlackInitRow = 7
)

func getNextPlayer(currentPlayer int) int {
	nextPlayer := currentPlayer + 1
	if nextPlayer >= playerNum {
		nextPlayer = firstPlayer
	}
	return nextPlayer
}
