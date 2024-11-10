package tetrisopt

import (
	"errors"
	"fmt"
)

// Accepts a 4x4 shape of '.' and '#' and checks if tetromino is valid
// if it is it updates the cropped shape, height, width and area to the t instace of Tetromino
func validate(sqr []string, t *Tetronimo) error {
	err := errors.New("ERROR")
	var coordinates [][]int
	var emptyLines int
	var indexLeft, indexRight, indexTop, indexBottom int
	var setStartRow, setEndRow, setStartCol, setEndCol bool

	// checking for empty lines and registering coordinates
	// setting up the frame of tetronimo corners with inclusive borders
	// on left and top and exclusive on right and bottom
	for a := 0; a < 4; a++ {
		rowEmpty, colEmpty := true, true
		for b := 0; b < 4; b++ {
			if sqr[a][b] == '#' {
				rowEmpty = false
				coordinates = append(coordinates, []int{a, b})
				if !setStartRow {
					indexTop = a
					setStartRow = true
				}
				if setStartRow && a > indexBottom {
					indexBottom = a
					setEndRow = true
				}
			}
			if setEndRow {
				indexBottom = a
				setEndRow = false
			}

			if sqr[b][a] == '#' {
				colEmpty = false
				if !setStartCol {
					indexLeft = a
					setStartCol = true
				}
				if setStartCol && a > indexRight {
					indexRight = a
					setEndCol = true
				}
			}
			if setEndCol {
				indexRight = a
				setEndCol = false
			}
		}

		if rowEmpty {
			emptyLines++
		}

		if colEmpty {
			emptyLines++
		}
	}
	// early return by checking number of empty lines
	if emptyLines < 3 || emptyLines > 4 || len(coordinates) != 4 {
		fmt.Println("Input file err: Too many empty lines")
		return err
	}

	// checking for edge case .##.
	//                        #.#.
	if emptyLines == 3 {
		for _, cur := range coordinates {
			if !hasTangent(cur, coordinates) {
				fmt.Println("Input file err: Tangent not found")
				return err
			}
		}
	}

	returnIndexes(sqr, indexLeft, indexRight, indexTop, indexBottom, t)
	return nil
}

// checks if the top, bottom, left and right neighbour is '#'
func hasTangent(cur []int, coordinates [][]int) bool {
	for _, other := range coordinates {
		if other[0] == cur[0] && other[1] == cur[1] {
			continue
		}
		if (cur[0] == other[0] && (cur[1]-1 == other[1] || cur[1]+1 == other[1])) ||
			(cur[1] == other[1] && (cur[0] == other[0]-1 || cur[0] == other[0]+1)) {
			return true
		}
	}
	return false
}
