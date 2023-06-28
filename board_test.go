package kamisado

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBoard = newBoard()

func TestBoard_init(t *testing.T) {
	testBoard.init()
	assert.Equal(t, boardHeight, len(testBoard), "init failed: wrong row number")
	for _, boardRow := range testBoard {
		assert.Equal(t, boardWidth, len(boardRow), "init failed: wrong column number")
	}
}

// |---------------------------------------|
// |m   |b   |c   |g   |y   |r   |p   |o   |
// |---------------------------------------|
// |r   |m bO|g   |p bO|b   |y bO|o   |c   |
// |---------------------------------------|
// |p   |g   |m   |r   |c   |o   |y   |b   |
// |---------------------------------------|
// |g   |c bO| b  |m wC|o   |p bO|r   |y   |
// |---------------------------------------|
// |y   |r   |p   |o   |m   |b   |c   |g   |
// |---------------------------------------|
// |b   |y bO|o   |c bO|r   |m bO|g   |p   |
// |---------------------------------------|
// |c   |o   |y   |b   |p   |g   |m   |r   |
// |---------------------------------------|
// |o   |p   |r   |y   |g   |c   |b   |m   |
// |---------------------------------------|

// abnormal moves:
//     path occupied

func TestBoard_checkMovePathOccupied(t *testing.T) {
	testBoard.init()
	testBoard[1][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[1][3].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[1][5].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[3][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[3][3].Piece = &Piece{pieceColorOrange, playerWhite}
	testBoard[3][5].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][3].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][5].Piece = &Piece{pieceColorOrange, playerBlack}

	var err error
	// normal case
	err = testBoard.checkMovePathOccupied(Coord{0,0}, Coord{0,2})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{0,2}, Coord{2,2})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{2,2}, Coord{2,0})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{2,0}, Coord{0,0})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{1,0}, Coord{3,2})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{3,2}, Coord{1,0})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{1,2}, Coord{3,0})
	assert.Nil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{3,0}, Coord{1,2})
	assert.Nil(t, err, "")

	// abnormal case
	err = testBoard.checkMovePathOccupied(Coord{0,1}, Coord{7,1})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{7,1}, Coord{0,1})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{1,0}, Coord{1,7})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{1,7}, Coord{1,0})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{0,0}, Coord{7,7})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{7,7}, Coord{0,0})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{6,0}, Coord{0,6})
	assert.NotNil(t, err, "")
	err = testBoard.checkMovePathOccupied(Coord{0,6}, Coord{6,0})
	assert.NotNil(t, err, "")
}

// abnormal moves:
//     not move
//     backward
//     wrong move rule(not \-|/)
func TestBoard_checkMoveRuleValid(t *testing.T) {
	testBoard.init()
	testBoard[1][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[1][3].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[1][5].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[3][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[3][3].Piece = &Piece{pieceColorOrange, playerWhite}
	testBoard[3][5].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][1].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][3].Piece = &Piece{pieceColorOrange, playerBlack}
	testBoard[5][5].Piece = &Piece{pieceColorOrange, playerBlack}

	var err error
	// normal case:
	err = testBoard.checkMoveRuleValid(playerWhite, Coord{0,0}, Coord{1,1})
	assert.Nil(t, err, "")
	err = testBoard.checkMoveRuleValid(playerBlack, Coord{0,0}, Coord{1,1})
	assert.Nil(t, err, "")

}


// abnormal moves:
//     no piece
//     piece not belongs to player
//     color not match
func TestBoard_checkMoveValid(t *testing.T) {
	// nextColor, err = testBoard.movePiece(playerBlack, nextColor, Coord{7, 7}, Coord{6, 7})
	// assert.Nil(t, err, "")
	// nextColor, err = testBoard.movePiece(playerWhite, nextColor, Coord{0, 5}, Coord{2, 5})
	// assert.Nil(t, err, "")
	// // You Shall Not Pass!
	// testBoard.init()
	// nextColor, err = testBoard.movePiece(playerWhite, nonPieceColor, Coord{4, 4}, Coord{1, 1})
	// assert.NotNil(t, err, "")
}
