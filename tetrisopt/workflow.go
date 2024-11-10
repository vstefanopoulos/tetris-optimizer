package tetrisopt

import (
	"fmt"
)

type Tetronimo struct {
	name       string
	height     int
	width      int
	hashIndxes [][]int
	hasDouble  bool
	used       bool
	area       int
}

func Run(args []string) error {
	var tetronimos []Tetronimo
	var sqrs [][]string

	// Check args
	fileName, err := handleArgs(args)
	if err != nil {
		fmt.Println("ERROR")
		return err
	}

	// read the file
	lines, err := readFile(fileName)
	if err != nil {
		return err
	}

	// make 2D grid and check for obvious formating errors
	sqrs, err = make2DGrid(*lines)
	if err != nil {
		return err
	}

	// validate each square and add to Invetory
	for _, sqr := range sqrs {
		t := &Tetronimo{}
		err = validate(sqr, t)
		if err != nil {
			return err
		}
		t.name = string(rune(65 + len(tetronimos)))

		/*
			commented out a call to a function that names the same shape
			that is given more than one time with the same letter
			addToInventory(t, tetronimos)
		*/

		tetronimos = append(tetronimos, *t)
	}

	// init smallest possible board
	var enlarge int
	minArea := len(tetronimos) * 4
	board := initBoard(minArea, enlarge)

	// call solution with enlarged board until it returns a board
	for {
		solution := solve(tetronimos, board)
		if solution {
			break
		}
		enlarge++
		board = initBoard(minArea, enlarge)
	}

	printBoard(board)

	return nil
}
