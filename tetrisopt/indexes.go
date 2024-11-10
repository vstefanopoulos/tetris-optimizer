package tetrisopt

func returnIndexes(sqr []string, indexLeft, indexRight, indexTop, indexBottom int, t *Tetronimo) {
	// crop sqr to fit tetronimo
	var row, col int
	for i := indexTop; i <= indexBottom; i++ {
		for j := indexLeft; j <= indexRight; j++ {
			if sqr[i][j] == '#' {
				index := []int{row, col}
				t.hashIndxes = append(t.hashIndxes, index)
			}
			col++
		}
		row++
		col = 0
	}
	t.height = indexBottom - indexTop + 1
	t.width = indexRight - indexLeft + 1
	t.area = t.height * t.width
}
