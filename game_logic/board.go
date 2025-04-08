package game_logic

import "fmt"

func PrintBoard(x, y int8) {
	var i int8 = 0
	var j int8 = 0

	for i = 0; i < x; i++ {
		for j = 0; j < y; j++ {
			fmt.Print("[ ] ")
		}

		fmt.Print("\r\n")
	}
}
