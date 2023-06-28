package kamisado

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCommand Command


func TestCommand_parseContentToMove(t *testing.T) {
		testCommand = Command{
			Type: cmdTypeMove,
			Content: `{"from":{"x":1,"y":2},"to":{"x":1,"y":2}}`,
		}
  _, err := testCommand.parseContentToMove()
	assert.Nil(t, err, "")
}
