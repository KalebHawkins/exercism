package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankedCount := 0

	for k, v := range cb {
		if rank == k {
			for _, r := range v {
				if r {
					rankedCount++
				}
			}
		}
	}

	return rankedCount
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file <= 0 || file > 8 {
		return 0
	}

	fileCount := 0

	for _, v := range cb {
		for i, r := range v {
			if i == file-1 && r {
				fileCount++
			}
		}
	}

	return fileCount
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	squareCount := 0

	for _, v := range cb {
		for range v {
			squareCount++
		}
	}

	return squareCount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	occupiedCount := 0

	for _, v := range cb {
		for _, r := range v {
			if r {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}
