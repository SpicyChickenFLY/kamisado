package main

import (
	"fmt"
)

// PlaygroundBackend - Playground structs implement this interface can be
// deployed by our engine
type PlaygroundBackend interface {
	InitNewGame() error
	GetGameData() (data string, err error)
	ExecuteCmd(command string) error
	SaveGame(filePath string) error
	LoadGame(filePath string) error
}

func main() {
	// NewPlayground()
	playground := NewPlayground(gameModeStandard)
	playground.InitNewGame()
	fmt.Println(playground.GetGameData())
	// playground.Move
	cmdStr := `{

	}`
	playground.ExecuteCmd(cmdStr)
}
