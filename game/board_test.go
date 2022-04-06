package game

import (
	"testing"
)

var testBoard = newBoard()

func TestBoard_init(t *testing.T) {
	testBoard.init()
	// assert.Nil(t, err, fmt.Sprintf("test failed: %v", err))
	// assert.NotEqual(t, 0, len(data))
	// TODO: check every square
}

func TestBoard_movePiece(t *testing.T) {

}
