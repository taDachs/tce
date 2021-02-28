package board

import (
	"strconv"
	"strings"
)

type BitBoard struct {
	board            [][][]bool
	blackCastleQueen bool
	blackCastleKing  bool
	whiteCastleQueen bool
	whiteCastleKing  bool
	whitesTurn       bool
	turn             int
	halfmove         int
	enPassant        []int
}

func (board *BitBoard) GetTurn() int {
	return board.turn
}

func (board *BitBoard) SetTurn(turn int) {
	board.turn = turn
}

func (board *BitBoard) SetWhitesTurn(whitesTurn bool) {
	board.whitesTurn = whitesTurn
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
			if board.isFieldEmpty(i, j) {
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
	if board.whiteCastleKing {
		output += "K"
	}
	if board.whiteCastleQueen {
		output += "Q"
	}
	if board.blackCastleKing {
		output += "k"
	}
	if board.blackCastleQueen {
		output += "q"
	}

	if !(board.whiteCastleKing || board.whiteCastleQueen || board.blackCastleKing || board.blackCastleQueen) {
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
			board.whiteCastleKing = true
		case "Q":
			board.whiteCastleQueen = true
		case "k":
			board.blackCastleKing = true
		case "q":
			board.blackCastleQueen = true
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

func (board BitBoard) Equal(other BitBoard) bool {
	for i, row := range board.board {
		for j, field := range row {
			for k, value := range field {
				if value != other.board[i][j][k] {
					return false
				}
			}
		}
	}

	if board.blackCastleQueen != other.blackCastleQueen || board.blackCastleKing != other.blackCastleKing ||
		board.whiteCastleQueen != other.whiteCastleQueen || board.whiteCastleKing != other.whiteCastleKing ||
		board.whitesTurn != other.whitesTurn || board.turn != other.turn || board.halfmove != other.halfmove ||
		board.enPassant[0] != other.enPassant[0] || board.enPassant[1] != other.enPassant[1] {
		return false
	}

	return true
}

func (board BitBoard) IsWhitesTurn() bool {
	return board.whitesTurn
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

func (board *BitBoard) isFieldEmpty(x, y int) bool {
	for _, value := range board.board[x][y] {
		if value {
			return false
		}
	}
	return true
}

func (board *BitBoard) isFieldBlack(x, y int) bool {
	if board.isFieldEmpty(x, y) {
		return false
	}

	for i := 0; i < 6; i++ {
		if board.board[x][y][i] {
			return true
		}
	}

	return false
}

func (board *BitBoard) isFieldWhite(x, y int) bool {
	if board.isFieldEmpty(x, y) {
		return false
	}

	for i := 6; i < 12; i++ {
		if board.board[x][y][i] {
			return true
		}
	}

	return false
}

func (board *BitBoard) isFieldAvailable(x, y int, white bool) bool {
	if board.isFieldEmpty(x, y) {
		return true
	}
	if white {
		return board.isFieldBlack(x, y)
	} else {
		return board.isFieldWhite(x, y)
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
	if board.isFieldEmpty(x, y) {
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

			if !combinedMatrix.isFieldEmpty(x, y) {
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

func (board *BitBoard) IsCheckMate(white bool) bool {
	if !board.IsCheck(white) {
		return false
	}
	for i, row := range board.board {
		for j := range row {
			if (white && board.isFieldWhite(i, j)) || (!white && board.isFieldBlack(i, j)) {
				piece := board.GetPieceOnField(i, j)
				movementMatrix := piece.GetMovementMatrix(board, i, j, false)
				for k, rowMovement := range movementMatrix.board {
					for l := range rowMovement {
						if !movementMatrix.isFieldEmpty(k, l) {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

func (board *BitBoard) IsMoveValid(x1, y1, x2, y2 int) bool {
	piece := board.GetPieceOnField(x1, y1)

	if piece.IsNone() || (piece.IsWhite() != board.whitesTurn) || (piece.IsBlack() != !board.whitesTurn) {
		return false
	}

	movementMatrix := piece.GetMovementMatrix(board, x1, y1, false)
	if movementMatrix.isFieldEmpty(x2, y2) {
		return false
	}

	return true
}

func (board *BitBoard) GetCastleRightsQueenSideAfterPieceMove(x, y int) bool {
	if (!board.whiteCastleQueen && board.whitesTurn) || (!board.blackCastleQueen && !board.whitesTurn) {
	    return false
	}
	piece := board.GetPieceOnField(x, y)
	if (piece == WHITE_KING) && board.whitesTurn {
	    return false
	}

	if (piece == BLACK_KING) && !board.whitesTurn {
		return false
	}

	if (piece == WHITE_ROOK) && board.whitesTurn && ((x == 0) && (y == 0)) {
		return false
	}

	if (piece == BLACK_ROOK) && !board.whitesTurn && ((x == 0) && (y == 7)) {
		return false
	}

	return true
}

func (board *BitBoard) GetCastleRightsKingSideAfterPieceMove(x, y int) bool {
	if (!board.whiteCastleKing && board.whitesTurn) || (!board.blackCastleKing && !board.whitesTurn) {
		return false
	}
	piece := board.GetPieceOnField(x, y)
	if (piece == WHITE_KING) && board.whitesTurn {
		return false
	}

	if (piece == BLACK_KING) && !board.whitesTurn {
		return false
	}

	if (piece == WHITE_ROOK) && board.whitesTurn && ((x == 7) && (y == 0)) {
		return false
	}

	if (piece == BLACK_ROOK) && !board.whitesTurn && ((x == 7) && (y == 7)) {
		return false
	}

	return true
}

func (board *BitBoard) SetEnPassant(x, y int) {
	board.enPassant = []int{x, y}
}

func (board *BitBoard) GetEnPassant() []int {
	return board.enPassant
}