package tetrisopt

func solve(inventory []Tetronimo, board [][]string) bool {
	side := len(board)
	piece, allUsed := selectPiece(inventory)
	if allUsed {
		return true
	}
	for row := 0; row < side; row++ {
		for col := 0; col < side; col++ {
			if itFits(inventory[piece], board, col, row, side) {
				placePiece(inventory[piece], board, col, row)
				inventory[piece].used = true
				if solve(inventory, board) {
					return true
				}
				removePiece(inventory[piece], board, col, row)
				inventory[piece].used = false
			}
		}
	}
	return false
}

func selectPiece(inventory []Tetronimo) (int, bool) {
	var largest, index int
	var allUsed bool = true
	for i, piece := range inventory {
		if !piece.used && piece.area > largest {
			index, largest = i, piece.area
			allUsed = false
		}
	}
	return index, allUsed
}

func itFits(piece Tetronimo, board [][]string, row, col, side int) bool {
	if side-row < piece.height || side-col < piece.width {
		return false
	}

	for _, index := range piece.hashIndxes {
		if board[row+index[0]][col+index[1]] != "" {
			return false
		}
	}
	return true
}

func removePiece(piece Tetronimo, board [][]string, row, col int) {
	for _, index := range piece.hashIndxes {
		board[row+index[0]][col+index[1]] = ""
	}
}

func placePiece(piece Tetronimo, board [][]string, row, col int) {
	for _, index := range piece.hashIndxes {
		board[row+index[0]][col+index[1]] = piece.name
	}
}
