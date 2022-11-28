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
	Content, err := os.ReadFile("Files/" + os.Args[1])
	if err != nil {
		log.Fatal(fmt.Println("Error: ", err))
	}
	sep := "\n"
	if runtime.GOOS == "windows" {
		sep = string([]byte{13, 10})
	}
	Words := strings.Split(string(Content), string(sep))
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
	data.RevealedLettres = nil
	NumberOfLetterRevealed := (len(data.Word) / 2) - 1

	if len(data.Word) >= 3 {
		for len(data.RevealedLettres) < NumberOfLetterRevealed {
			RandomLetter := string(data.Word[rand.Intn(len(data.Word))])
			if !Contains(data.RevealedLettres, RandomLetter) {
				data.RevealedLettres = append(data.RevealedLettres, RandomLetter)
			}
		}
	}
}
