package main

func isValidSudoku(board [][]byte) bool {
	var box [3][3][9]bool
	// row & col
	for i := range 9 {
		var r, c [9]bool
		for ii := range 9 {
			br := board[i][ii]
			bc := board[ii][i]

			// row
			if br != 46 {
				// box
				if box[i/3][ii/3][br-49] {
					return false
				} else {
					box[i/3][ii/3][br-49] = true
				}

				if r[br-49] {
					return false
				} else {
					r[br-49] = true
				}
			}
			// col
			if bc != 46 {
				if c[bc-49] {
					return false
				}
				c[bc-49] = true
			}
		}
	}

	return true

}
