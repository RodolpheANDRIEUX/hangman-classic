package hangman

import "fmt"

// PrintWord Display the word to find without the missing letters
func PrintWord() {
	Display := ""
	for _, l := range game.Word {
		if Contains(game.RevealedLettres, string(l)) {
			Display += " " + string(l)
		} else {
			Display += " _"
		}
	}
	fmt.Println(Display)
}

// PrintHangman Supposed to print hangman here
func PrintHangman() {
	fmt.Printf("%v tries left\n", 10-game.Tries)
}
