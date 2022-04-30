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
	Board      board     `json:"board"`
	Mode       int       `json:"mode"`
	Records    []*Record `json:"records"`
	NextPlayer int       `json:"next_player"`
	NextColor  int       `json:"next_color"`
}

// NewGame return *Game
func NewGame(gameMode int) *Game {
	g := &Game{}
	g.init(gameMode)
	return g
}

// init board and data in game
func (g *Game) init(gameMode int) {
	g.Mode = gameMode
	g.Records = make([]*Record, 0)
	g.Board.init()
}

func (g *Game) start() {
	g.Board.start()
	g.NextPlayer = firstPlayer
	g.NextColor = nonPieceColor
}

// GetGameData encode Playground information to json struct
func (g *Game) GetGameData() (data string, err error) {
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// GetBoard data as json
func (g *Game) getBoard() (boardData string, err error) {
	jsonBytes, err := g.Board.marshalBoardJSON()
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// ExecuteCmd by player
func (g *Game) ExecuteCmd(player int, cmd *Command) error {
	if player != g.NextPlayer {
		return errors.New("Not your turn")
	}
	switch cmd.Type {
	case cmdTypeMove:
		move, err := cmd.parseContentToMove()
		if err != nil {
			return err
		}
		// move piece and write log
		nextColor, err := g.Board.movePiece(player, g.NextColor, move.From, move.To)
		if err != nil {
			return err
		}
		newRecord := &Record{
			Turn:   len(g.Records) + 1,
			Player: player,
			Move:   *move,
		}
		g.NextColor = nextColor
		g.NextPlayer = getNextPlayer(player)
		g.Records = append(g.Records, newRecord)
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
