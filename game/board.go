package game

import "errors"

const (
	boardWidth    = 8
	boardHeight   = 8
	pieceColorNum = 8
)

const (
	nonPieceColor = -1 + iota
	pieceColorPurple
	pieceColorBrown
	pieceColorBlue
	pieceColorGreen
	pieceColorYellow
	pieceColorRed
	pieceColorPink
	pieceColorOrange
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

// Coodinator is the data format of coodinator
type Coodinator struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type piece struct {
	color      int
	ownerColor int
}

type square struct {
	piece *piece
	color int
}

type board [boardHeight][boardWidth]square

func newBoard() *board {
	b := board{}
	b.init()
	return &b
}

func (b *board) init() {
	// init board
	for x := 0; x < boardHeight; x++ {
		for y := 0; y < boardWidth; y++ {
			b[x][y] = square{
				color: defaultBoardColor[x][y],
				piece: nil,
			}
		}
	}
	// place pieces
	for j := 0; j < boardWidth; j++ {
		squareForPlayerWhite := b[playerWhiteInitRow][j]
		squareForPlayerWhite.piece = &piece{
			color:      squareForPlayerWhite.color,
			ownerColor: playerColorWhite,
		}
		squareForPlayerBlack := b[playerBlackInitRow][j]
		squareForPlayerBlack.piece = &piece{
			color:      squareForPlayerBlack.color,
			ownerColor: playerColorBlack,
		}
	}
}

func (b *board) movePiece(playerColor int, from, to Coodinator) error {
	// check source square
	if b[from.X][from.Y].piece == nil {
		return errors.New("No piece found in position(%d, %d)")
	}
	// check target square
	if b[to.X][to.Y].piece == nil {
		return errors.New("A piece occupied in position(%d, %d)")
	}
	// move the piece
	b[to.X][to.Y].piece = b[from.X][from.Y].piece
	b[from.X][from.Y].piece = nil

	return nil
}
