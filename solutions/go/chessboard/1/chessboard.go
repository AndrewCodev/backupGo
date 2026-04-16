package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok{
        return 0
    }

    count := 0
    for _, occupied := range f {
        if occupied {
			count++
		}
    }
    return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }

	count := 0
    for _, file := range cb {
        if file[rank-1] {
            count++
        }
    }
    return count
}

func CountAll(board Chessboard) int {
	count := 0
	for _, file := range board {
		count += len(file)
	}
	return count
}

func CountOccupied(cb Chessboard) int {
    count := 0
	for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                count++
            }
        }
    }
    return count
}
