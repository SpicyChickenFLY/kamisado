package game

import (
	"encoding/json"
)

const (
	gameModeStandard = iota
)

// Game is compelete Game information
type Game struct {
	board board
	data  data
}

// NewGame return *Game
func NewGame(gameMode int) *Game {
	g := &Game{}
	g.Init(gameMode)
	return g
}

// Init board and data in game
func (g *Game) Init(gameMode int) {
	g.data.GameMode = gameMode
	g.data.Records = make([]*Record, 0)
	g.board.init()
}

// GetGameData encode Playground information to json struct
func (g *Game) GetGameData() (data string, err error) {
	jsonBytes, err := json.Marshal(g.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// func (g *Game) parseMoveCmdContent(cmd *Command) (*Record, error) {
// 	coorParts := strings.Split(cmd.Content, coorsSep)
// 	if len(coorParts) != 2 {
// 		return nil, errors.New("Wrong format of coodinators")
// 	}
// 	srcCoor := strings.Split(coorParts[0], coorSep)
// 	dstCoor := strings.Split(coorParts[1], coorSep)
// 	if len(srcCoor) != 2 || len(dstCoor) != 2 {
// 		return nil, errors.New("Wrong format of coodinators")
// 	}
// 	srcX, err := strconv.Atoi(srcCoor[0])
// 	if err != nil {
// 		return nil, err
// 	}
// 	srcY, err := strconv.Atoi(srcCoor[1])
// 	if err != nil {
// 		return nil, err
// 	}
// 	dstX, err := strconv.Atoi(dstCoor[0])
// 	if err != nil {
// 		return nil, err
// 	}
// 	dstY, err := strconv.Atoi(dstCoor[1])
// 	if err != nil {
// 		return nil, err
// 	}
// 	newRecord := &Record{
// 		Turn:  g.data.Records[len(g.data.Records)].Turn + 1,
// 		Color: cmd.Player,
// 		From:  Coodinator{X: srcX, Y: srcY},
// 		To:    Coodinator{X: dstX, Y: dstY},
// 	}
// 	return newRecord, nil
// }

// func (g *Game) executeMoveCmd(player int, cmd *Command) error {
// 	if player != g.data.CurrentPlayer {
// 		return errors.New("Not your turn")
// 	}
// 	record, err := g.parseMoveCmdContent(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	// move piece and write log
// 	err = g.board.movePiece(record.Color, record.From, record.To)
// 	if err != nil {
// 		return err
// 	}
// 	g.data.Records = append(g.data.Records, record)
// 	return nil
// }

// // ExecuteCmd execute command
// func (g *Game) ExecuteCmd(cmdStr string) error {
// 	command := Command{}
// 	if err := json.Unmarshal([]byte(cmdStr), &command); err != nil {
// 		return err
// 	}
// 	// TODO: execute command
// 	switch command.Type {
// 	case cmdTypeMove:
// 		return g.executeMoveCmd(command.Player, &command)
// 	}
// 	return nil
// }

// Save game data to file
func (g *Game) Save(filePath string) error {
	return nil
}

// Load game data from file
func (g *Game) Load(filePath string) error {
	return nil
}
