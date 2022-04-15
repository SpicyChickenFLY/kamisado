package kamisado

import (
	"encoding/json"
	"errors"
)

const (
	gameModeStandard = iota
)

// Game is compelete Game information
type Game struct {
	board      board
	data       Data
	nextPlayer int
	nextColor  int
}

// NewGame return *Game
func NewGame(gameMode int) *Game {
	g := &Game{}
	g.init(gameMode)
	return g
}

// init board and data in game
func (g *Game) init(gameMode int) {
	g.data.GameMode = gameMode
	g.data.Records = make([]*Record, 0)
	g.board.init()
	g.nextPlayer = firstPlayer
	g.nextColor = nonPieceColor
}

// GetGameData encode Playground information to json struct
func (g *Game) GetGameData() (data string, err error) {
	jsonBytes, err := json.Marshal(g.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// ExecuteCmd by player
func (g *Game) ExecuteCmd(player int, cmd *Command) error {
	if player != g.nextPlayer {
		return errors.New("Not your turn")
	}
	switch cmd.Type {
	case cmdTypeMove:
		move, err := cmd.parseContentToMove()
		if err != nil {
			return err
		}
		// move piece and write log
		nextColor, err := g.board.movePiece(player, g.nextColor, move.From, move.To)
		if err != nil {
			return err
		}
		newRecord := &Record{
			Turn:   len(g.data.Records) + 1,
			Player: player,
			Move:   *move,
		}
		g.nextColor = nextColor
		g.nextPlayer = getNextPlayer(player)
		g.data.Records = append(g.data.Records, newRecord)
	}
	return nil
}

// Save game data to file
func (g *Game) Save(filePath string) error {
	return nil
}

// Load game data from file
func (g *Game) Load(filePath string) error {
	return nil
}
