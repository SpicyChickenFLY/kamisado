package game

const (
	coorsSep = "@"
	coorSep  = ":"
)

// Coodinator is the data format of coodinator
type Coodinator struct {
	XPos int `json:"x_pos"`
	YPos int `json:"y_pos"`
}

// Record is the data format of game log
type Record struct {
	Index      int        `json:"index"`
	PlayerTurn int        `json:"player_turn"`
	SrcCoor    Coodinator `json:"src_coor"`
	DstCoor    Coodinator `json:"dst_coor"`
}

type data struct {
	GameMode     int       `json:"game_mode"`
	Records      []*Record `json:"log"`
	Turn         int       `json:"turn"`
	CurrentColor int       `json:"current_color"`
}
