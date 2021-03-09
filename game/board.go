package game

const (
	boardWidth    = 8
	boardHeight   = 8
	pieceColorNum = 8
)

const (
	nonPiece = -1 + iota
	pieceColorPurple
	pieceColorBrown
	pieceColorBlue
	pieceColorGreen
	pieceColorYellow
	pieceColorRed
	pieceColorPink
	pieceColorOrange
)

var defaultBoardColor = [boardHeight][boardWidth]int{
	{pieceColorPurple, pieceColorBrown, pieceColorBlue, pieceColorGreen, pieceColorYellow, pieceColorRed, pieceColorPink, pieceColorOrange},
	{pieceColorRed, pieceColorPurple, pieceColorGreen, pieceColorPink, pieceColorBrown, pieceColorYellow, pieceColorOrange, pieceColorBlue},
	{pieceColorPink, pieceColorGreen, pieceColorPurple, pieceColorRed, pieceColorBlue, pieceColorOrange, pieceColorYellow, pieceColorBrown},
	{pieceColorGreen, pieceColorBlue, pieceColorBrown, pieceColorPurple, pieceColorOrange, pieceColorPink, pieceColorRed, pieceColorYellow},
	{pieceColorYellow, pieceColorRed, pieceColorPink, pieceColorOrange, pieceColorPurple, pieceColorBrown, pieceColorBlue, pieceColorGreen},
	{pieceColorBrown, pieceColorYellow, pieceColorOrange, pieceColorBlue, pieceColorRed, pieceColorPurple, pieceColorGreen, pieceColorPink},
	{pieceColorBlue, pieceColorOrange, pieceColorYellow, pieceColorBrown, pieceColorPink, pieceColorGreen, pieceColorPurple, pieceColorRed},
	{pieceColorOrange, pieceColorPink, pieceColorRed, pieceColorYellow, pieceColorGreen, pieceColorBlue, pieceColorBrown, pieceColorPurple},
}

type piece struct {
	color      int
	ownerColor int
}

type square struct {
	piece *piece
	color int
}

type token struct {
	board    [boardHeight][boardWidth]square
	pieceBox [playerNum][pieceColorNum]piece
}
