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

// normal moves:
// abnormal moves:
//     no piece
//     piece not belongs to player
//     color not match
//     backward
//     cross blocked
//     wrong move rule(not \-|/)
func TestBoard_movePiece(t *testing.T) {
	var nextColor int
	var err error
	// Should Pass:
	testBoard.init()
	testBoard.start()
	nextColor, err = testBoard.movePiece(playerWhite, nonPieceColor, Coodinator{0, 0}, Coodinator{1, 1})
	assert.Nil(t, err, "")
	nextColor, err = testBoard.movePiece(playerBlack, nextColor, Coodinator{7, 7}, Coodinator{6, 7})
	assert.Nil(t, err, "")
	nextColor, err = testBoard.movePiece(playerWhite, nextColor, Coodinator{0, 5}, Coodinator{2, 5})
	assert.Nil(t, err, "")
	// You Shall Not Pass!
	testBoard.init()
	nextColor, err = testBoard.movePiece(playerWhite, nonPieceColor, Coodinator{4, 4}, Coodinator{1, 1})
	assert.NotNil(t, err, "")
}
