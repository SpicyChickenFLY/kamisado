package kamisado

import (
	"fmt"
	"math"
)

const (
	boardWidth    = 8
	boardHeight   = 8
	pieceColorNum = 8
)

// pieceColor
const (
	nonPieceColor = -1 + iota
	pieceColorMagenta
	pieceColorBrown
	pieceColorCyan
	pieceColorGreen
	pieceColorYellow
	pieceColorRed
	pieceColorPink
	pieceColorOrange
)

// Piece contain color & owner color
type Piece struct {
	Color      int `json:"color"`
	OwnerColor int `json:"owner_color"`
}

// Square of board
type Square struct {
	Piece *Piece `json:"piece"`
	Color int    `json:"color"`
}

type board [boardHeight][boardWidth]Square

func newBoard() *board {
	b := board{}
	b.init()
	return &b
}

func (b *board) init() {
	for c := 0; c < boardWidth; c++ {
		for x := 0; x < boardHeight; x++ {
			y := ((2*c+1)*x + c) % boardWidth
			b[x][y] = Square{Color: c, Piece: nil}
		}
	}
}

func (b *board) start() {
	// place pieces
	for j := 0; j < boardWidth; j++ {
		b[playerWhiteInitRow][j].Piece = &Piece{
			Color:      b[playerWhiteInitRow][j].Color,
			OwnerColor: playerWhite,
		}
		b[playerBlackInitRow][j].Piece = &Piece{
			Color:      b[playerBlackInitRow][j].Color,
			OwnerColor: playerBlack,
		}
	}
}

func (b *board) checkMoveValid(playerColor, gameColor int, from, to Coord) error {
	// check coordinateout of bound
	if from.X >= boardHeight || from.X < 0 || from.Y >= boardWidth || from.Y < 0 {
		return fmt.Errorf("source pos(%+v) is out of boud", from)
	}
	if to.X >= boardHeight || to.X < 0 || to.Y >= boardWidth || to.Y < 0 {
		return fmt.Errorf("target pos(%+v) is out of boud", to)
	}

	// check piece choice
	if b[from.X][from.Y].Piece == nil {
		return fmt.Errorf("no piece has been found in pos(%+v)", from)
	}
	if b[from.X][from.Y].Piece.OwnerColor != playerColor {
		return fmt.Errorf(
			"piece in pos(%+v) is not belongs to player(%d)", from, playerColor)
	}
	if gameColor != nonPieceColor && b[from.X][from.Y].Piece.Color != gameColor {
		return fmt.Errorf(
			"piece in pos(%+v) doesn't match turn color(%d)", from, gameColor)
	}

	// check move rule
	return b.checkMoveRuleValid(playerColor, from, to)
}

func (b *board) checkMoveRuleValid(playerColor int, from, to Coord) error {
	if from.X == to.X && from.Y == to.Y {
		return fmt.Errorf("source/target pos can not be the same pos(%+v)", from)
	}
	if to.Y > from.Y && playerColor == playerBlack || to.Y < from.Y && playerColor == playerWhite {
		return fmt.Errorf("piece can not move backward from pos(%+v) to pos(%+v)", from, to)
	}
	if (from.X == to.X) || (from.Y == to.Y) || (from.X-to.X == from.Y-to.Y) || (from.X-to.X == to.Y-from.Y) {
		return b.checkMovePathOccupied(from, to)
	}

	return fmt.Errorf("piece can only move sideways, vertically, and obliquely")
}

func (b *board) checkMovePathOccupied(from, to Coord) error {
	verticalDirection, horizontalDirection := 0, 0
	length := int(math.Max(math.Abs(float64(to.Y-from.Y)), math.Abs(float64(to.X-from.X))))
	if from.X == to.X {
		horizontalDirection = (to.Y - from.Y) / length
	} else if from.Y == to.Y {
		verticalDirection = (to.X - from.X) / length
	} else if (from.X-to.X == from.Y-to.Y) || (from.X-to.X == to.Y-from.Y) {
		verticalDirection = (to.X - from.X) / length
		horizontalDirection = (to.Y - from.Y) / length
	}

	for offset := 1; offset <= length; offset++ {
		if b[from.X+offset*verticalDirection][from.Y+offset*horizontalDirection].Piece != nil {
			return fmt.Errorf("path from pos(%+v) to pos(%+v) is occupied by other piece(s)", from, to)
		}
	}

	return nil
}

func (b *board) movePiece(playerColor, gameColor int, from, to Coord) (int, error) {
	if err := b.checkMoveValid(playerColor, gameColor, from, to); err != nil {
		return nonPieceColor, err
	}

	// move the piece
	b[to.X][to.Y].Piece = b[from.X][from.Y].Piece
	b[from.X][from.Y].Piece = nil

	return b[to.X][to.Y].Color, nil
}
