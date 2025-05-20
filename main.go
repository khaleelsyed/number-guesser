package main

import (
	"fmt"
)

const welcomeMessage string = `
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.`

func main() {
	var difficultyUserInput string
	gameSession := NewGameSession()
	fmt.Println(welcomeMessage)

	fmt.Println("\nPlease select the difficulty level:")
	for i, value := range difficultyLabel {
		fmt.Printf("%d: %s\n", i, value)
	}

	fmt.Print("\nChosen Difficulty: ")
	fmt.Scan(&difficultyUserInput)
	err := gameSession.setSession(difficultyUserInput)
	if err != nil {
		panic(err)
	}

	var guess int
	for try := range difficultyTries[gameSession.Difficulty] {
		// TODO: Countdown retries

		fmt.Printf("Guess %d: ", try+1)
		fmt.Scan(&guess)

		pass, message := gameSession.checkGuess(guess)
		if pass {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", try)
			break
		}
		fmt.Println(message)
	}

}
