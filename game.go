package main

import (
	"bufio"
	"fmt"
	"gameoflife/game/game_logic"
	"log"
	"os"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	display_game(0, oldState)
	scanner := bufio.NewReader(os.Stdin)

	for {
		char, _, err := scanner.ReadRune()

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			panic(err)
		}

		display_game(char, oldState)
	}
}

func display_game(char rune, oldState *term.State) {
	if char == 3 {
		fmt.Println("\r\nFound some Ctrl+C\r")
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	} else {
		fmt.Print("\033[H\033[2J\r")
		game_logic.PrintBoard(5, 5)
		if char != 0 {
			fmt.Printf("Selected char: %v \r\n", char)
		}
		fmt.Print("Getting new char")
	}
}
