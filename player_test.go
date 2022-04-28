package kamisado

import (
	"fmt"
	"testing"
)

func TestPlayer_getNextPlayer(t *testing.T) {
	testPlayer := 1
	for i := 0; i < 10; i++ {
		testPlayer = getNextPlayer(testPlayer)
		fmt.Println(testPlayer)
	}
}
