package game

const (
	coorsSep = "@"
	coorSep  = ":"
)

// Record is the data format of game log
type Record struct {
	Turn   int        `json:"turn"`
	Player int        `json:"player"`
	From   Coodinator `json:"from"`
	To     Coodinator `json:"to"`
}

type data struct {
	GameMode int       `json:"game_mode"`
	Records  []*Record `json:"records"`
}
