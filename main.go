package main

const (
	boardWidth  = 8
	boardHeight = 8
)

const (
	nonPiece = iota
	pieceColorPurple
	pieceColorBrown
	pieceColorBlue
	pieceColorGreen
	pieceColorYellow
	pieceColorRed
	pieceColorPink
	pieceColorOrange
)

const (
	playerColorWhite = 0
	playerColorBlack = 1
)

type piece struct {
	color      int
	ownerColor int
}

type square struct {
	piece *piece
	color int
}

type Game struct {
	board [boardHeight][boardWidth]square
	turn  int
}

type GameData struct {
	Squares []int `json:"Squares"`
	Pieces  []int `json:"Pieces"`
	Turn    int   `json:"Turn"`
}

func (g *Game) Marshall() (result GameData) {
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			result.Squares = append(result.Squares, g.board[i][j].color)
			piece := *g.board[i][j].piece
			result.Pieces = append(result.Pieces, piece.color)
		}
	}
}

func main() {

}
