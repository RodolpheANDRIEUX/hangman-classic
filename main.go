package main

import (
	"flag"
	hangman "hangman-classic-for-web/Fonctions"
	"math/rand"
	"time"
)

var Save string

// init random seed and declare flags
func init() {
	rand.Seed(time.Now().UnixNano())

	const (
		defaultSave = ""
		usage       = "the name of the save file"
	)
	flag.StringVar(&Save, "startWith", defaultSave, usage)
	flag.StringVar(&Save, "s", defaultSave, usage+" (shorthand)")
}

func main() {
	flag.Parse()
	hangman.Run(Save)
}
