package hangman

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"runtime"
)

// GetWord returns a random word from the words.txt file
func GetWord() string {
	Data, err := os.ReadFile("Files/" + os.Args[1])
	if err != nil {
		log.Fatal(fmt.Println("Error: ", err))
	}
	sep := "\n"
	if runtime.GOOS == "windows" {
		sep = string([]byte{13, 10})
	}
	Words := strings.Split(string(Data), string(sep))
	Word := strings.ToUpper(Words[rand.Intn(len(Words))])
	return Word

}

// Contains returns true if the given list contains the given letter
func Contains(list []string, n string) bool {
	for _, l := range list {
		if l == n {
			return true
		}
	}
	return false
}

// StartingLettersReveal Reveals (len(word) /2 - 1) letter(s) to the word at the beginning
func StartingLettersReveal() {
	game.RevealedLettres = nil
	NumberOfLetterRevealed := (len(game.Word) / 2) - 1

	if len(game.Word) >= 3 {
		for len(game.RevealedLettres) < NumberOfLetterRevealed {
			RandomLetter := string(game.Word[rand.Intn(len(game.Word))])
			if !Contains(game.RevealedLettres, RandomLetter) {
				game.RevealedLettres = append(game.RevealedLettres, RandomLetter)
			}
		}
	}
}
