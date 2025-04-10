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
	var cursorX, cursorY int8
	var arrowUp, arrowDown, arrowRight, arrowLeft rune = 65, 66, 67, 68
	var boardSize int8 = 8
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	display_game(0, cursorX, cursorY, boardSize, oldState)
	scanner := bufio.NewReader(os.Stdin)

	for {
		char, _, err := scanner.ReadRune()

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			panic(err)
		}

		switch char {
		case arrowUp:
			cursorY -= 1
			break
		case arrowDown:
			cursorY += 1
			break
		case arrowRight:
			cursorX += 1
			break
		case arrowLeft:
			cursorX -= 1
			break
		}

		if cursorY < 0 {
			cursorY = boardSize - 1
		}
		if cursorY >= boardSize {
			cursorY = 0
		}
		if cursorX < 0 {
			cursorX = boardSize - 1
		}
		if cursorX >= boardSize {
			cursorX = 0
		}

		display_game(char, cursorX, cursorY, boardSize, oldState)
	}
}

func display_game(char rune, cursorX int8, cursorY int8, boardSize int8, oldState *term.State) {
	if char == 3 {
		fmt.Print("\033[H\033[2J")
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Welcome to game of life. For explanation how the game functions visit: https://en.wikipedia.org/wiki/Conway's_Game_of_Life\r")
		fmt.Println("To start, move the cursor around the board using arrow keys and pressing \"space\" to change the state of a tile on a board.\r")
		fmt.Println("To animate the board, press \"Enter\".\r")
		fmt.Println()
		game_logic.PrintBoard(boardSize, boardSize)

		command := "\x1b["
		command += string(5+cursorY) + ";"
		command += string(2 + cursorX*4)
		command += "H"

		fmt.Fprint(os.Stdout, "\x1b[")
		fmt.Fprint(os.Stdout, 5+cursorY)
		fmt.Fprint(os.Stdout, ";")
		fmt.Fprint(os.Stdout, 2+cursorX*4)
		fmt.Fprint(os.Stdout, "H")
	}
}
