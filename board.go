package kamisado

import "errors"

const (
	boardWidth    = 8
	boardHeight   = 8
	pieceColorNum = 8
)

// pieceColor
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
			ownerColor: playerWhite,
		}
		squareForPlayerBlack := b[playerBlackInitRow][j]
		squareForPlayerBlack.piece = &piece{
			color:      squareForPlayerBlack.color,
			ownerColor: playerBlack,
		}
	}
}

func (b *board) movePiece(playerColor, gameColor int, from, to Coodinator) (int, error) {
	// check source square
	if b[from.X][from.Y].piece == nil {
		return nonPieceColor, errors.New("No piece found in position(%d, %d)")
	}
	// check target square
	if b[to.X][to.Y].piece == nil {
		return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
	}
	// check piece choice
	if b[from.X][from.Y].piece.ownerColor != playerColor {
		return nonPieceColor, errors.New("The piece(%d, %d) doesn't belongs to player(%d)")
	}
	if b[from.X][from.Y].color != nonPieceColor && b[from.X][from.Y].color != gameColor {
		return nonPieceColor, errors.New("This square(%d, %d) doesn't match color(%d)")
	}
	// TODO: check move rule
	if false {
		return nonPieceColor, errors.New("piece can only move sideways, vertically, and obliquely,; and cannnot move backward")
	}

	// move the piece
	b[to.X][to.Y].piece = b[from.X][from.Y].piece
	b[from.X][from.Y].piece = nil

	return b[to.X][to.Y].color, nil
}
