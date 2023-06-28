package kamisado

import (
	"encoding/json"
)

// Data of this game can be exported as json
type Data struct {
}

// Record is the data format of game log
type Record struct {
	Turn   int `json:"turn"`
	Player int `json:"player"`
	Move
}

// Move tell piece from where and to where
type Move struct {
	From Coord `json:"from"`
	To   Coord `json:"to"`
}

// Coord is the data format of coodinator
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

const (
	cmdTypeMove = iota
)

// Command to interactive with game
type Command struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

func (c *Command) parseContentToMove() (move *Move, err error) {
	err = json.Unmarshal([]byte(c.Content), &move)
	if err != nil {
		return nil, err
	}
	return move, nil
}
