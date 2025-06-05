package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool


// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
/*
Chessboard := {
	"A" = [true, false, true, false, false, false, false, true]
}
*/

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
	for i, x := range cb {
        if i == file {
            for _, b := range x {
                if b == true {
                    count ++
                }
            }
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0
    if rank >= 1 && rank <= 8 {
    	for _, x := range cb {
            if x[rank - 1] == true {
                count ++
            }
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
    for i := 1; i <= len(cb); i++ {
    	for range cb {
    		count ++
    	}
	}
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for _, x := range cb {
        for _, b := range x {
            if b == true {
                count ++
            }
        }
    }
    return count
}
