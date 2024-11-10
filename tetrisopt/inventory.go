package tetrisopt

func addToInventory(t *Tetronimo, tetronimos []Tetronimo) {
	for _, submited := range tetronimos {
		if t.hasDouble = compareShapes(submited.hashIndxes, t.hashIndxes); t.hasDouble {
			t.name = submited.name
			t.hasDouble = true
			break
		}
	}
	if !t.hasDouble && tetronimos != nil {
		var count int
		for i := range tetronimos {
			if tetronimos[i].hasDouble {
				continue
			}
			count++
		}
		t.name = string(rune(65 + count))
	}
}

func compareShapes(a, b [][]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
