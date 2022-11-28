package hangman

import (
	"fmt"
)

type Game struct {
	Word                   string
	Tries                  int
	Guess, RevealedLettres []string
}

var game Game

func NewGame() {
	game.Word = GetWord()
	game.Tries = 0
	StartingLettersReveal()
}

// WordIsCompleted returns true if the word is fully completed
func WordIsCompleted() bool {
	for _, EveryLetter := range game.Word {
		if !Contains(game.RevealedLettres, string(EveryLetter)) {
			return false
		}
	}
	return true
}

// IsGoodAnswer returns true if the word contains the letter guessed
func IsGoodAnswer(guess string) bool {
	for _, l := range game.Word {
		if string(l) == guess || len(guess) > 1 && guess == game.Word {
			return true
		}
	}
	return false
}

// ProcessLetter manages the guess to increment or not the tries
func ProcessLetter(guess string) {
	if IsGoodAnswer(guess) {
		if len(guess) > 1 {
			for _, i := range guess {
				game.RevealedLettres = append(game.RevealedLettres, string(i))
			}
		} else {
			game.RevealedLettres = append(game.RevealedLettres, guess)
		}
	} else {
		if len(guess) > 1 && game.Tries < 9 {
			game.Tries++
		} else {
			game.Guess = append(game.Guess, guess)
		}
		game.Tries++
		PrintHangman()
	}
}

// Run is a call function that runs all the game
func Run(Save string) {
	if Save != "" {
		ReadJSON("Saves/" + Save + ".json")
	} else {
		NewGame()
	}

	for game.Tries < 10 && !WordIsCompleted() {
		PrintWord()
		guess := AskForLetter()
		ProcessLetter(guess)
	}

	if WordIsCompleted() {
		fmt.Printf("GG!")
	} else {
		fmt.Printf("GAME OVER")
	}
}
