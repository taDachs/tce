package board

type BitBoard struct {
	board            [][][]bool
	blackCastleLeft  bool
	blackCastleRight bool
	whiteCastleLeft  bool
	whiteCastleRight bool
	move             int
}

func CreateEmptyBitBoard() BitBoard {
	board := make([][][]bool, 8)
	for i := 0; i < 8; i++ {
		row := make([][]bool, 8)
		for j := 0; j < 8; j++ {
			field := make([]bool, 12)
			row[j] = field
		}
		board[i] = row
	}

	return BitBoard{board, true, true, true, true, 0}
}

func And(a, b BitBoard) BitBoard {
	result := CreateEmptyBitBoard()

	for i, row := range result.board {
		for j, field := range row {
			for k := range field {
				result.board[i][j][k] = a.board[i][j][k] && b.board[i][j][k]
			}
		}
	}

	return result
}

func Or(a, b BitBoard) BitBoard {
	result := CreateEmptyBitBoard()

	for i, row := range result.board {
		for j, field := range row {
			for k := range field {
				result.board[i][j][k] = a.board[i][j][k] || b.board[i][j][k]
			}
		}
	}

	return result
}

func Not(a BitBoard) BitBoard {
	result := CreateEmptyBitBoard()

	for i, row := range result.board {
		for j, field := range row {
			for k := range field {
				result.board[i][j][k] = !a.board[i][j][k]
			}
		}
	}

	return result
}

func (board *BitBoard) String() string {
	symbols := [12]string{"p", "r", "n", "b", "q", "k", "P", "R", "N", "B", "Q", "K"}
	output := ""
	for _, row := range board.board {
		for _, field := range row {
			piece := "-"
			for k, value := range field {
				if value {
					piece = symbols[k]
				}
			}
			output += " " + piece
		}
		output += "\n"
	}

	return output
}

func (board *BitBoard) IsFieldEmpty(x, y int) bool {
	for _, value := range board.board[x][y] {
		if value {
			return false
		}
	}
	return true
}

func (board *BitBoard) IsFieldBlack(x, y int) bool {
	if board.IsFieldEmpty(x, y) {
		return false
	}

	for i := 0; i < 6; i++ {
		if board.board[x][y][i] {
			return true
		}
	}

	return false
}

func (board *BitBoard) IsFieldWhite(x, y int) bool {
	if board.IsFieldEmpty(x, y) {
		return false
	}

	for i := 6; i < 12; i++ {
		if board.board[x][y][i] {
			return true
		}
	}

	return false
}

func (board *BitBoard) IsFieldAvailable(x, y int, white bool) bool {
	if board.IsFieldEmpty(x, y) {
		return true
	}
    if white {
    	return board.IsFieldBlack(x, y)
	} else {
		return board.IsFieldWhite(x, y)
	}
}

func (board *BitBoard) PlacePieceOnBoard(x, y int, piece Piece) {
	board.board[x][y][piece] = true
}

func GetStartBoard() BitBoard {
	board := CreateEmptyBitBoard()
	for i := 0; i < 8; i++ {
		board.PlacePieceOnBoard(i, 1, WHITE_PAWN)
		board.PlacePieceOnBoard(i, 6, BLACK_PAWN)
	}

	board.PlacePieceOnBoard(0, 0, WHITE_ROOK)
	board.PlacePieceOnBoard(7, 0, WHITE_ROOK)
	board.PlacePieceOnBoard(0, 7, BLACK_ROOK)
	board.PlacePieceOnBoard(7, 7, BLACK_ROOK)

	board.PlacePieceOnBoard(1, 0, WHITE_KNIGHT)
	board.PlacePieceOnBoard(6, 0, WHITE_KNIGHT)
	board.PlacePieceOnBoard(1, 7, BLACK_KNIGHT)
	board.PlacePieceOnBoard(6, 7, BLACK_KNIGHT)

	board.PlacePieceOnBoard(2, 0, WHITE_BISHOP)
	board.PlacePieceOnBoard(5, 0, WHITE_BISHOP)
	board.PlacePieceOnBoard(2, 7, BLACK_BISHOP)
	board.PlacePieceOnBoard(5, 7, BLACK_BISHOP)

	board.PlacePieceOnBoard(3, 0, WHITE_QUEEN)
	board.PlacePieceOnBoard(4, 0, WHITE_KING)
	board.PlacePieceOnBoard(3, 7, BLACK_QUEEN)
	board.PlacePieceOnBoard(4, 7, BLACK_KING)
	return board
}