package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    row, exists := cb[file]
    count :=0
    if !exists{
        return 0
    } else {
        for _, hash:= range row{
            if hash{
                count +=1
            }
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0
    if rank < 1 || rank > 8{
        return 0
    } else{
        for _, value := range cb{
            if value[rank-1]{
                count++
            }
        }
	}
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    var rows int
    for _, row :=range cb{
        rows = len(row)
        break
    }
    cols := len(cb)
    return rows * cols
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    var count int 
    for _, rows := range cb {
        for _, cols := range rows{
            if cols{
                count ++
            }
        }
    }
    return count
}
