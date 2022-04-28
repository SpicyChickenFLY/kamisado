package kamisado

import (
	"encoding/json"
	"errors"
)

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

type boardData struct {
	PiecesInfo [][]int `json:"pieces_info"`
	SquareInfo [][]int `json:"square_info"`
}

type board [boardHeight][boardWidth]square

func (b *board) marshalBoardJSON() ([]byte, error) {
	bd := boardData{
		PiecesInfo: make([][]int, 0),
		SquareInfo: make([][]int, boardHeight),
	}
	for h := 0; h < boardHeight; h++ {
		bd.SquareInfo[h] = make([]int, boardWidth)
		for w := 0; w < boardWidth; w++ {
			square := b[h][w]
			bd.SquareInfo[h][w] = square.color
			if square.piece != nil {
				bd.PiecesInfo = append(bd.PiecesInfo,
					[]int{square.piece.ownerColor, square.piece.color, h, w})
			}
		}
	}
	jsonStr, err := json.Marshal(bd)
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}

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
}

func (b *board) start() {
	// place pieces
	for j := 0; j < boardWidth; j++ {
		b[playerWhiteInitRow][j].piece = &piece{
			color:      b[playerWhiteInitRow][j].color,
			ownerColor: playerWhite,
		}
		b[playerBlackInitRow][j].piece = &piece{
			color:      b[playerBlackInitRow][j].color,
			ownerColor: playerBlack,
		}
	}
}

func (b *board) movePiece(playerColor, gameColor int, from, to Coodinator) (int, error) {
	// check out of bound
	if from.X >= boardHeight || from.X < 0 || from.Y >= boardWidth || from.Y < 0 {
		return nonPieceColor, errors.New("position(%d, %d) out of boud")
	}
	if to.X >= boardHeight || to.X < 0 || to.Y >= boardWidth || to.Y < 0 {
		return nonPieceColor, errors.New("position(%d, %d) out of boud")
	}
	// check source square
	if b[from.X][from.Y].piece == nil {
		return nonPieceColor, errors.New("No piece found in position(%d, %d)")
	}
	// check piece choice
	if b[from.X][from.Y].piece.ownerColor != playerColor {
		return nonPieceColor, errors.New("The piece(%d, %d) doesn't belongs to player(%d)")
	}
	if gameColor != nonPieceColor && b[from.X][from.Y].color != gameColor {
		return nonPieceColor, errors.New("This square(%d, %d) doesn't match color(%d)")
	}
	// check move rule
	if from.X == to.X && from.Y == to.Y {
		return nonPieceColor, errors.New("source and target position are the same")
	} else if to.Y > from.Y && playerColor == playerBlack || to.Y < from.Y && playerColor == playerWhite {
		return nonPieceColor, errors.New("piece cannot move backward")
	} else if from.X == to.X {
		for x := from.X; x <= to.X; x++ {
			if b[x][to.Y].piece != nil {
				return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
			}
		}
	} else if from.Y == to.Y {
		for y := from.Y; y <= to.Y; y++ {
			if b[to.X][y].piece != nil {
				return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
			}
		}
	} else if from.X-to.X == from.Y-to.Y {
		if from.X < to.X {
			for offset := 1; offset <= to.X-from.X; offset++ {
				if b[from.X+offset][from.Y+offset].piece != nil {
					return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
				}
			}
		} else {
			for offset := 1; offset <= from.X-to.X; offset++ {
				if b[from.X-offset][from.Y-offset].piece != nil {
					return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
				}
			}
		}
	} else if from.X-to.X == to.Y-from.Y {
		if from.X < to.X {
			for offset := 1; offset <= to.X-from.X; offset++ {
				if b[from.X+offset][from.Y-offset].piece != nil {
					return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
				}
			}
		} else {
			for offset := 1; offset <= from.X-to.X; offset++ {
				if b[from.X-offset][from.Y+offset].piece != nil {
					return nonPieceColor, errors.New("A piece occupied in position(%d, %d)")
				}
			}
		}
	} else {
		return nonPieceColor, errors.New("piece can only move sideways, vertically, and obliquely; and cannnot move backward")
	}

	// move the piece
	b[to.X][to.Y].piece = b[from.X][from.Y].piece
	b[from.X][from.Y].piece = nil

	return b[to.X][to.Y].color, nil
}
