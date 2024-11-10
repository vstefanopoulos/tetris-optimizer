package tetrisopt

import (
	"errors"
	"fmt"
)

// Make 2D grid and early checks for formating errors
func make2DGrid(lines []string) ([][]string, error) {
	err := errors.New("ERROR")
	var sqrs [][]string
	var sqrComplete bool
	var count int
	for i, line := range lines {
		if line != "" {
			if len(line) != 4 {
				fmt.Println("Input file err: Unexpected line width")
				return nil, err
			}
			sqrComplete = false
			count++
		}

		if line == "" && sqrComplete {
			continue
		}

		if count != 4 && line == "" {
			fmt.Println("Input file err: Too many lines")
			return nil, err
		}

		if line == "" {
			sqrs = append(sqrs, lines[i-4:i])
			sqrComplete = true
			count = 0
		}
	}
	return sqrs, nil
}
