package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	boardWidth    = 8
	boardHeight   = 8
	pieceColorNum = 8
	playerNum     = 2
)

const (
	gameModeStandard = iota
)

const (
	nonPiece = -1 + iota
	pieceColorPurple
	pieceColorBrown
	pieceColorBlue
	pieceColorGreen
	pieceColorYellow
	pieceColorRed
	pieceColorPink
	pieceColorOrange
)

const (
	nonPlayer = -1 + iota
	playerColorWhite
	playerColorBlack
)

const (
	playerWhiteInitRow = 7
	playerBlackInitRow = 0
)

const (
	cmdTypePulse = iota
	cmdTypeMove
	cmdTypeUndo
)

const (
	coorsSep = "@"
	coorSep  = ":"
)

var defaultBoardColor = [boardHeight][boardWidth]int{
	{pieceColorPurple, pieceColorBrown, pieceColorBlue, pieceColorGreen, pieceColorYellow, pieceColorRed, pieceColorPink, pieceColorOrange},
	{pieceColorRed, pieceColorPurple, pieceColorGreen, pieceColorPink, pieceColorBrown, pieceColorYellow, pieceColorOrange, pieceColorBlue},
	{pieceColorPink, pieceColorGreen, pieceColorPurple, pieceColorRed, pieceColorBlue, pieceColorOrange, pieceColorYellow, pieceColorBrown},
	{pieceColorGreen, pieceColorBlue, pieceColorBrown, pieceColorPurple, pieceColorOrange, pieceColorPink, pieceColorRed, pieceColorYellow},
	{pieceColorYellow, pieceColorRed, pieceColorPink, pieceColorOrange, pieceColorPurple, pieceColorBrown, pieceColorBlue, pieceColorGreen},
	{pieceColorBrown, pieceColorYellow, pieceColorOrange, pieceColorBlue, pieceColorRed, pieceColorPurple, pieceColorGreen, pieceColorPink},
	{pieceColorBlue, pieceColorOrange, pieceColorYellow, pieceColorBrown, pieceColorPink, pieceColorGreen, pieceColorPurple, pieceColorRed},
	{pieceColorOrange, pieceColorPink, pieceColorRed, pieceColorYellow, pieceColorGreen, pieceColorBlue, pieceColorBrown, pieceColorPurple},
}

// PlaygroundBackend - Playground structs implement this interface can be
// deployed by our engine
type PlaygroundBackend interface {
	InitNewGame() error
	GetGameData() (data string, err error)
	ExecuteCmd(command string) error
	SaveGame(filePath string) error
	LoadGame(filePath string) error
}

type piece struct {
	color      int
	ownerColor int
}

type square struct {
	piece *piece
	color int
}

// Coodinator is the data format of coodinator
type Coodinator struct {
	XPos int `json:"x_pos"`
	YPos int `json:"y_pos"`
}

// Record is the data format of game log
type Record struct {
	Index      int        `json:"index"`
	PlayerTurn int        `json:"player_turn"`
	SrcCoor    Coodinator `json:"src_coor"`
	DstCoor    Coodinator `json:"dst_coor"`
}

// Command is the data format of game command
type Command struct {
	Player  int    `json:"player"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// Playground is compelete Playground information
type Playground struct {
	board    [boardHeight][boardWidth]square
	pieceBox [playerNum][pieceColorNum]piece
	data     struct {
		GameMode     int       `json:"game_mode"`
		Records      []*Record `json:"log"`
		Turn         int       `json:"turn"`
		CurrentColor int       `json:"current_color"`
	}
}

// NewPlayground init a new Playground
func NewPlayground(gameMode int) *Playground {
	pg := &Playground{}
	pg.data.GameMode = gameMode
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			pg.board[i][j].color = defaultBoardColor[i][j]
		}
	}
	for j := 0; j < boardWidth; j++ {
		pg.pieceBox[playerColorWhite][j] = piece{
			color:      defaultBoardColor[playerWhiteInitRow][j],
			ownerColor: playerColorWhite,
		}
		pg.pieceBox[playerColorBlack][j] = piece{
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
			pg.board[i][j].piece = nil
		}
	}
	// place piece from pieceBox to board
	for j := 0; j < boardWidth; j++ {
		pg.board[playerWhiteInitRow][j].piece = &pg.pieceBox[playerColorWhite][boardWidth-1-j]
		pg.board[playerBlackInitRow][j].piece = &pg.pieceBox[playerColorBlack][j]
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

// checkMoveValid check
func (pg *Playground) checkMoveValid(record *Record) bool {
	// check player's piece exists in src square
	srcPiece := pg.board[record.SrcCoor.XPos][record.SrcCoor.YPos].piece
	if srcPiece == nil || srcPiece.ownerColor != record.PlayerTurn {
		return false
	}
	// check no pieces exists in dst square
	if pg.board[record.DstCoor.XPos][record.DstCoor.YPos].piece != nil {
		return false
	}
	return true
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

func (pg *Playground) executeMoveCmd(player int, cmd *Command) error {
	// check if game is still running
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
	pg.board[record.SrcCoor.XPos][record.SrcCoor.YPos].piece =
		pg.board[record.DstCoor.XPos][record.DstCoor.YPos].piece
	pg.board[record.DstCoor.XPos][record.DstCoor.YPos].piece = nil
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

func main() {
	// NewPlayground()
	playground := NewPlayground(gameModeStandard)
	playground.InitNewGame()
	fmt.Println(playground.GetGameData())
	// playground.Move
	cmdStr := `{

	}`
	playground.ExecuteCmd(cmdStr)
}
