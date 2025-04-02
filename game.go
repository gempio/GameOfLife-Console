package main

import (
	"fmt"
	"gameoflife/game/game_logic"
	"log"

	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	fmt.Println("Hello, this is an implementation of game of life. You can setup the board before pressing \"Enter\" To start.")

	r := 0

	for {
		fmt.Print("\033[H\033[2J")
		game_logic.PrintBoard(5, 5)

		if r == 'q' {
			fmt.Println("Quitting game")
			break
		}

		if r != 0 {
			fmt.Printf("You pressed: %d", r)
		}

		char, err := tty.ReadRune()

		r = int(char)

		if err != nil {
			log.Fatal(err)
		}
	}
}
