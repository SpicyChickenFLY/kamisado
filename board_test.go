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

func TestBoard_marshalBoardJSON(t *testing.T) {
	testBoard.init()
	jsonBytes, err := testBoard.marshalBoardJSON()
	assert.Nil(t, err, "")
	assert.NotEqual(t, 0, len(jsonBytes))
}

// normal moves:
// abnormal moves:
//     no piece
//     square occupied
//     piece not belongs to player
//     color not match
//     wrong move rule(not \-|/ OR move backward OR cross blocked)
func TestBoard_movePiece(t *testing.T) {
	var nextColor int
	var err error
	// Should Pass:
	testBoard.init()
	nextColor, err = testBoard.movePiece(playerWhite, nonPieceColor, Coodinator{0, 0}, Coodinator{1, 1})
	assert.Nil(t, err, "")
	nextColor, err = testBoard.movePiece(playerBlack, nextColor, Coodinator{0, 0}, Coodinator{1, 1})
	assert.Nil(t, err, "")
	nextColor, err = testBoard.movePiece(playerWhite, nextColor, Coodinator{0, 0}, Coodinator{1, 1})
	assert.Nil(t, err, "")
	// You Shall Not Pass!
	testBoard.init()
	nextColor, err = testBoard.movePiece(playerWhite, nonPieceColor, Coodinator{0, 0}, Coodinator{1, 1})
	assert.NotNil(t, err, "")
}
