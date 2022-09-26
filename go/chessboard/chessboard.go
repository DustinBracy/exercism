package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File[8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, v := range cb[file]{
		if v {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank > 8 || rank < 1 {
		return 0
	}
	for _, file := range cb {
		if file[rank - 1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for range cb{
		count++
	}
	return count * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for file, _ := range cb{
		count += CountInFile(cb, file)
	}
	return count
}
