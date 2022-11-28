package hangman

import (
	"fmt"
)

type Game struct {
	Word                   string
	Tries                  int
	Guess, RevealedLettres []string
}

var data Game

func NewGame() {
	data.Word = GetWord()
	data.Tries = 0
	StartingLettersReveal()
}

// WordIsCompleted returns true if the word is fully completed
func WordIsCompleted() bool {
	for _, EveryLetter := range data.Word {
		if !Contains(data.RevealedLettres, string(EveryLetter)) {
			return false
		}
	}
	return true
}

// IsGoodAnswer returns true if the word contains the letter guessed
func IsGoodAnswer(guess string) bool {
	for _, l := range data.Word {
		if string(l) == guess || len(guess) > 1 && guess == data.Word {
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
				data.RevealedLettres = append(data.RevealedLettres, string(i))
			}
		} else {
			data.RevealedLettres = append(data.RevealedLettres, guess)
		}
	} else {
		if len(guess) > 1 && data.Tries < 9 {
			data.Tries++
		} else {
			data.Guess = append(data.Guess, guess)
		}
		data.Tries++
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

	for data.Tries < 10 && !WordIsCompleted() {
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
