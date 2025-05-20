package main

import "errors"

type DifficultyLevel int

const (
	Easy DifficultyLevel = iota + 1
	Medium
	Hard
)

var ErrInvalidGameDifficulty error = errors.New("invalid difficulty selected")

var difficultyLabel = map[DifficultyLevel]string{
	Easy:   "Easy (10 chances)",
	Medium: "Medium (5 chances)",
	Hard:   "Hard (3 chances)",
}

var difficultyTries = map[DifficultyLevel]int{
	Easy:   10,
	Medium: 5,
	Hard:   3,
}

func (d DifficultyLevel) IsValid() bool {
	switch d {
	case Easy, Medium, Hard:
		return true
	}
	return false
}
