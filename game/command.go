package game

const (
	cmdTypePulse = iota
	cmdTypeMove
	cmdTypeUndo
)

// Command is the data format of game command
type Command struct {
	Player  int    `json:"player"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}
