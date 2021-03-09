package service

// PlaygroundBackend - Playground structs implement this interface can be
// deployed by our engine
type PlaygroundBackend interface {
	InitNewGame() error
	GetGameData() (data string, err error)
	ExecuteCmd(command string) error
	SaveGame(filePath string) error
	LoadGame(filePath string) error
}
