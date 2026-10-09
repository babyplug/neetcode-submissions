func isValidSudoku(board [][]byte) bool {
	// var valid bool

	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	boxes := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
    	rows[i] = make(map[byte]struct{})
    	cols[i] = make(map[byte]struct{})
		boxes[i] = make(map[byte]struct{})
	}

	findBoxIdx := func(i, j int) int {
		// [0, 1, 2]
		// [3, 4, 5]
		// [6, 7, 8]
		// (0, 0), (0, 2), (2, 2) -> 0
		// (3, 3), (3, 5), (5, 5) -> 4
		// (6, 6), (6, 8), (8, 8) -> 8
		// Old Formula: (i / 3) + (j / 3) -> (8, 8) = 2 + 2 = 4
		// New Formula: ((i / 3) * 3) + (j / 3) ->(8, 8) = (2 * 3) + 2 = 8
		return ((i / 3) * 3) + (j / 3)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			key := board[i][j]
			if key == '.' {
				continue
			}
			// 1) Check rows
			if _, seen := rows[i][key]; seen {
				return false
			}
			// 2) Check cols
			if _, seen := cols[j][key]; seen {
				return false
			}
			boxIdx := findBoxIdx(i, j)
			if _, seen := boxes[boxIdx][key]; seen {
				return false
			}
			rows[i][key] = struct{}{}
			cols[j][key] = struct{}{}
			boxes[boxIdx][key] = struct{}{}
		}
	}
	
	return true
}
