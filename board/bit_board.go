package board

import (
	"strconv"
	"strings"
)

type BitBoard struct {
	board            [][][]bool
	blackCastleLeft  bool
	blackCastleRight bool
	whiteCastleLeft  bool
	whiteCastleRight bool
	whitesTurn       bool
	turn             int
	halfmove		 int
	enPassant		 []int
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

	return BitBoard{board, true, true, true, true, true, 0, 0, []int{-1, -1}}
}

func GetStartBoard() BitBoard {
	return FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
}

// TODO: halfmove clock and en passant field
// [board (as usual)] [current side] [castle rights] [en passant field] [halfmove clock] [turn]
func (board *BitBoard) ToFEN() string {
	output := ""

	// board
	for j := 7; j >= 0; j-- {
		emptySquare := 0
		for i := 0; i < 8; i++ {
			if board.IsFieldEmpty(i, j) {
				emptySquare++
				continue
			}

			if emptySquare != 0 {
				output += strconv.Itoa(emptySquare)
				emptySquare = 0
			}

			piece := board.GetPieceOnField(i, j)
			output += piece.GetNotation()
		}

		if emptySquare != 0 {
			output += strconv.Itoa(emptySquare)
		}

		if j != 0 {
			output += "/"
		}
	}

	output += " "

	// active color
	if board.turn % 2 == 1 {
		output += "w"
	} else {
		output += "b"
	}

	output += " "

	// castle rights
	if board.whiteCastleRight {
		output += "K"
	}
	if board.whiteCastleLeft {
		output += "Q"
	}
	if board.blackCastleRight {
		output += "k"
	}
	if board.blackCastleLeft {
		output += "q"
	}

	if !(board.whiteCastleRight || board.whiteCastleLeft || board.blackCastleRight || board.blackCastleLeft) {
		output += "-"
	}

	// en passant
	if board.enPassant[0] == -1 || board.enPassant[1] == -1 {
		output += " -"
	} else {
		output += " " + RowColToAlgebra(board.enPassant[0], board.enPassant[1])
	}

	// halfmove clock
	output += " " + strconv.Itoa(board.halfmove)

	// turn
	output += " " + strconv.Itoa(board.turn)

	return output
}

func FromFEN(fen string) BitBoard {
    board := CreateEmptyBitBoard()

    fenSplitBySpace := strings.Split(fen, " ")
    boardString := fenSplitBySpace[0]
    currentSide := fenSplitBySpace[1]
    castleRights := fenSplitBySpace[2]
    enPassant := fenSplitBySpace[3]
	halfmove, _ := strconv.Atoi(fenSplitBySpace[4])
	turn, _ := strconv.Atoi(fenSplitBySpace[5])

    rows := strings.Split(boardString, "/")

    for j := 7; j >= 0; j-- {
    	row := rows[7 - j]
    	for i, notation := range strings.Split(row, "") {
			piece := GetPieceByNotation(notation)
    		board.PlacePieceOnBoard(i, j, piece)
		}
	}

	board.whitesTurn = currentSide == "w"

	for _, c := range strings.Split(castleRights, "") {
		switch c {
		case "K":
			board.whiteCastleRight = true
		case "Q":
			board.whiteCastleLeft = true
		case "k":
			board.blackCastleRight = true
		case "q":
			board.blackCastleLeft = true
		}
	}

	if enPassant == "-" {
		board.enPassant = []int{-1, -1}
	} else {
		board.enPassant = AlgebraToRowCol(enPassant)
	}

	board.halfmove = halfmove
	board.turn = turn

	return board
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
	output := ""
	for i, row := range board.board {
		for j := range row {
			piece := board.GetPieceOnField(i, j)
			output += " " + piece.GetNotation()
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
	for i := range board.board[x][y] {
		board.board[x][y][i] = false
	}
	if piece == NO_PIECE {
		return
	}
	board.board[x][y][piece] = true
}

func (board *BitBoard) GetPieceOnField(x, y int) Piece {
	if board.IsFieldEmpty(x, y) {
		return NO_PIECE
	}

	for i, piece := range board.board[x][y] {
		if piece {
			return Piece(i)
		}
	}

	return NO_PIECE
}

func (board *BitBoard) findKing(white bool) (int, int) {
	var piece Piece
	if white {
		piece = WHITE_KING
	} else {
		piece = BLACK_KING
	}

	for i, row := range board.board {
		for j := range row {
			if board.GetPieceOnField(i, j) == piece {
				return i, j
			}
		}
	}

	return -1, -1
}

func (board *BitBoard) IsCheck(white bool) bool {
	checkMatrix := CreateEmptyBitBoard()
	x, y := board.findKing(white)

	if x == -1 || y == -1 {
		return false
	}
	for i := 0; i < 12; i++ {
		checkMatrix.board[x][y][i] = true
	}

	for i, row := range board.board {
		for j := range row {
			piece := board.GetPieceOnField(i, j)
			if piece == NO_PIECE {
				continue
			}

			if (white && piece.IsWhite()) || (!white && piece.IsBlack()) {
				continue
			}

			movementMatrix := piece.GetMovementMatrix(board, i, j, true)
			combinedMatrix := And(checkMatrix, movementMatrix)

			if !combinedMatrix.IsFieldEmpty(x, y) {
				return true
			}
		}
	}
	return false
}

func (board *BitBoard) Copy() BitBoard {
	copy := CreateEmptyBitBoard()

	for i, row := range board.board {
		for j := range row {
			copy.PlacePieceOnBoard(i, j, board.GetPieceOnField(i, j))
		}
	}

	return copy
}

func (board *BitBoard) doesMoveResultInCheck(x1, y1, x2, y2 int, white bool) bool {
	movedBoard := board.Copy()

	piece := movedBoard.GetPieceOnField(x1, y1)
	movedBoard.PlacePieceOnBoard(x1, y1, NO_PIECE)
	movedBoard.PlacePieceOnBoard(x2, y2, piece)

	return movedBoard.IsCheck(white)
}
