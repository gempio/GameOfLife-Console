package game_logic

import "fmt"

func PrintBoard(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print("[ ] ")
		}

		fmt.Print("\r\n")
	}
}
