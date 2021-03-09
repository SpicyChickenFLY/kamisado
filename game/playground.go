package game

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

const (
	gameModeStandard = iota
)

// Playground is compelete Playground information
type Playground struct {
	token token
	data  data
}

// NewPlayground init a new Playground
func NewPlayground(gameMode int) *Playground {
	pg := &Playground{}
	pg.data.GameMode = gameMode
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			pg.token.board[i][j].color = defaultBoardColor[i][j]
		}
	}
	for j := 0; j < boardWidth; j++ {
		pg.token.pieceBox[playerColorWhite][j] = piece{
			color:      defaultBoardColor[playerWhiteInitRow][j],
			ownerColor: playerColorWhite,
		}
		pg.token.pieceBox[playerColorBlack][j] = piece{
			color:      defaultBoardColor[playerBlackInitRow][j],
			ownerColor: playerColorBlack,
		}
	}
	pg.data.Turn = playerColorWhite
	pg.data.CurrentColor = nonPiece
	return pg
}

// InitNewGame init new Game in Playground
func (pg *Playground) InitNewGame() {
	// remove all piece from board to pieceBox
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			pg.token.board[i][j].piece = nil
		}
	}
	// place piece from pieceBox to board
	for j := 0; j < boardWidth; j++ {
		pg.token.board[playerWhiteInitRow][j].piece = &pg.token.pieceBox[playerColorWhite][boardWidth-1-j]
		pg.token.board[playerBlackInitRow][j].piece = &pg.token.pieceBox[playerColorBlack][j]
	}
}

// GetGameData encode Playground information to json struct
func (pg *Playground) GetGameData() (data string, err error) {
	jsonBytes, err := json.Marshal(pg.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (pg *Playground) parseMoveCmdContent(cmd *Command) (*Record, error) {
	coorParts := strings.Split(cmd.Content, coorsSep)
	if len(coorParts) != 2 {
		return nil, errors.New("Wrong format of coodinators")
	}
	srcCoor := strings.Split(coorParts[0], coorSep)
	dstCoor := strings.Split(coorParts[1], coorSep)
	if len(srcCoor) != 2 || len(dstCoor) != 2 {
		return nil, errors.New("Wrong format of coodinators")
	}
	srcXPos, err := strconv.Atoi(srcCoor[0])
	if err != nil {
		return nil, err
	}
	srcYPos, err := strconv.Atoi(srcCoor[1])
	if err != nil {
		return nil, err
	}
	dstXPos, err := strconv.Atoi(dstCoor[0])
	if err != nil {
		return nil, err
	}
	dstYPos, err := strconv.Atoi(dstCoor[1])
	if err != nil {
		return nil, err
	}

	newRecord := &Record{
		Index:      pg.data.Records[len(pg.data.Records)].Index + 1,
		PlayerTurn: cmd.Player,
		SrcCoor:    Coodinator{XPos: srcXPos, YPos: srcYPos},
		DstCoor:    Coodinator{XPos: dstXPos, YPos: dstYPos},
	}
	return newRecord, nil
}

// checkMoveValid check
func (pg *Playground) checkMoveValid(record *Record) bool {
	// check player's piece exists in src square
	srcPiece := pg.token.board[record.SrcCoor.XPos][record.SrcCoor.YPos].piece
	if srcPiece == nil || srcPiece.ownerColor != record.PlayerTurn {
		return false
	}
	// check no pieces exists in dst square
	if pg.token.board[record.DstCoor.XPos][record.DstCoor.YPos].piece != nil {
		return false
	}
	return true
}

func (pg *Playground) executeMoveCmd(player int, cmd *Command) error {
	if player != pg.data.Turn {
		return errors.New("Not your turn")
	}
	record, err := pg.parseMoveCmdContent(cmd)
	if err != nil {
		return err
	}
	if !pg.checkMoveValid(record) {
		return errors.New("Invalid move")
	}
	// move piece and write log
	pg.token.board[record.SrcCoor.XPos][record.SrcCoor.YPos].piece =
		pg.token.board[record.DstCoor.XPos][record.DstCoor.YPos].piece
	pg.token.board[record.DstCoor.XPos][record.DstCoor.YPos].piece = nil
	pg.data.Records = append(pg.data.Records, record)
	return nil
}

// ExecuteCmd execute command
func (pg *Playground) ExecuteCmd(cmdStr string) error {
	command := Command{}
	if err := json.Unmarshal([]byte(cmdStr), &command); err != nil {
		return err
	}
	// TODO: execute command
	switch command.Type {
	case cmdTypeMove:
		return pg.executeMoveCmd(command.Player, &command)
	}
	return nil
}

// SaveGame save game data to json file
func (pg *Playground) SaveGame(filePath string) error {
	return nil
}

// LoadGame load game data from json file
func (pg *Playground) LoadGame(filePath string) error {
	return nil
}
