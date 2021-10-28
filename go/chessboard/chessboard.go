package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	for _, value := range cb[rank] {
		if value {
			counter += 1
		}

	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counter := 0
	// lenRank := len(cb["A"])
	// fmt.Println(lenRank)
	for _, value := range cb {
		if file > len(value) {
			return 0
		}
		// fmt.Println(key, value[file])
		if value[file-1] {
			counter += 1
		}
	}
	// fmt.Println(counter)
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	counterAll := 0

	//for key := range cb {
	//	counterRank := CountInRank(cb, key)
	//	counterAll += counterRank
	//}
	// using no index/key and value
	for range cb {
		for range cb {
			counterAll += 1
		}
	}

	return counterAll
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counterAll := 0
	for key := range cb {
		counterRank := CountInRank(cb, key)
		counterAll += counterRank
	}
	return counterAll
}
