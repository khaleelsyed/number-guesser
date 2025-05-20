package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

type GameSession struct {
	Difficulty    DifficultyLevel
	NumberToGuess Guess
}

func (g *GameSession) setSession(difficulty string) error {
	n, err := strconv.Atoi(difficulty)
	if err != nil {
		return err
	}

	level := DifficultyLevel(n)
	if level.IsValid() {
		g.Difficulty = level
		return nil
	}

	return ErrInvalidGameDifficulty
}

func (g *GameSession) checkGuess(userGuess int) (bool, string) {
	guess := Guess(userGuess)
	if !guess.IsValid() {
		return false, "invalid guess - you must choose between 1 and 100 (inclusive)"
	}

	if guess == g.NumberToGuess {
		return true, ""
	} else if g.NumberToGuess < guess {
		return false, fmt.Sprintf("Incorrect! The number is less than %d.", guess)
	} else {
		return false, fmt.Sprintf("Incorrect! The number is greater than %d.", guess)
	}
}

func NewGameSession() *GameSession {
	randomNumber := rand.IntN(100) + 1

	return &GameSession{
		Difficulty:    -1,
		NumberToGuess: Guess(randomNumber),
	}
}
