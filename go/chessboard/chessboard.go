package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if _, ok := cb[file]; !ok {
		return 0
	}

	var cnt int
	for _, v := range cb[file] {
		if v {
			cnt++
		}
	}

	return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > len(cb) || rank <= 0 {
		return 0
	}

	var cnt int
	for _, v := range cb {
		if v[rank-1] {
			cnt++
		}
	}

	return cnt
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var cnt int
	for _, v := range cb {
		cnt += len(v)
	}

	return cnt
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var cnt int

	for k := range cb {
		cnt += CountInFile(cb, k)
	}

	return cnt
}
