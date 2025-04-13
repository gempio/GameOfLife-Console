package game_logic

import (
	"fmt"
)

func PrintBoard(gameSquares [][]bool) {
	for y, row := range gameSquares {
		for x, _ := range row {
			if gameSquares[x][y] {
				fmt.Print("[X] ")
			} else {
				fmt.Print("[ ] ")
			}
		}

		fmt.Print("\r\n")
	}
}

func FlipBoardSquare(gameSquares [][]bool, x int8, y int8) {
	gameSquares[x][y] = !gameSquares[x][y]
}

func InitializeGameState(boardSize int8) [][]bool {
	var gameSquares = [][]bool{}
	var i int8

	//Initialize rows
	for i = 0; i < boardSize; i++ {
		row := make([]bool, boardSize)

		fmt.Println(len(row))

		gameSquares = append(gameSquares, row)
	}

	return gameSquares
}
