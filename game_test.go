package kamisado

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGame = NewGame(gameModeStandard)

func TestGame_init(t *testing.T) {
	testGame.init(gameModeStandard)
}

func TestGame_GetGameData(t *testing.T) {
	testGame.init(gameModeStandard)
	data, err := testGame.GetGameData()
	assert.Nil(t, err, fmt.Sprintf("test failed: %v", err))
	assert.NotEqual(t, 0, len(data))
}

func TestGame_ExecuteCmd(t *testing.T) {
	testGame.init(gameModeStandard)
	err := testGame.ExecuteCmd(
		playerWhite,
		&Command{
			Type: cmdTypeMove,
			Content: `{"from":{"x":4, "y":4}, "to":{"x":4, "y":4}}`,
		})
	assert.Nil(t, err, fmt.Sprintf("test failed: %v", err))
}
