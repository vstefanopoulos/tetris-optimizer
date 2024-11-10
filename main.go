package main

import (
	"fmt"
	"os"
	"tetris-optimizer/tetrisopt"
)

func main() {
	err := tetrisopt.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
