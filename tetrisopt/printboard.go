package tetrisopt

import (
	"fmt"
	"time"
)

func printBoard(solution [][]string) {
	var result string
	for _, line := range solution {
		for _, char := range line {
			if char == "" {
				char = "."
			}
			result += string(char) + " "
		}
		result += "\n"
	}
	fmt.Print(result)
}

func printAnimated(board [][]string) {
	printBoard(board)
	time.Sleep(0 * time.Millisecond)
	for i := 0; i < len(board); i++ {
		fmt.Print("\033[A\033[K")
	}
}
