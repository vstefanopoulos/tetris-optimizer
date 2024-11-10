package tetrisopt

func sqr(n int) int {
	for i := 2; i <= n/2; i++ {
		if i*i == n {
			return i
		} else if i*i > n {
			return i
		}
	}
	return 0
}

func initBoard(minArea, enlarge int) [][]string {
	side := sqr(minArea) + enlarge
	board := make([][]string, side)
	for i := range board {
		board[i] = make([]string, side)
	}
	return board
}
