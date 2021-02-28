package board

import "testing"

func TestCreateEmptyBitBoard(t *testing.T) {
	bitBoard := CreateEmptyBitBoard()
	board := bitBoard.board
	if len(board) != 8 {
		t.Error("Incorrect dimensions of board")
	}
	for _, row := range board {
		if len(row) != 8 {
			t.Error("Incorrect dimensions of board")
		}
		for _, field := range row {
			if len(field) != 12 {
				t.Error("Incorrect dimensions of board")
			}
			for _, value := range field {
				if value {
					t.Error("Board not empty")
				}
			}
		}
	}

	if !bitBoard.blackCastleQueen || !bitBoard.blackCastleKing ||
		!bitBoard.whiteCastleQueen || !bitBoard.whiteCastleKing {
		t.Error("castle rights are wrong")
	}
}

func TestBoardAnd(t *testing.T) {
	board1 := CreateEmptyBitBoard()
	board2 := CreateEmptyBitBoard()

	board1.board[0][0][0] = true
	board2.board[0][0][0] = true


	andBoard := And(board1, board2)
	if !andBoard.board[0][0][0] {
		t.Error("AND failed")
	}
}

func TestBoardNot(t *testing.T) {
	board1 := CreateEmptyBitBoard()

	board1.board[0][0][0] = true

	notBoard := Not(board1)
	if notBoard.board[0][0][0] || !notBoard.board[0][0][1] {
		t.Error("NOT failed")
	}
}

func TestBoardOr(t *testing.T) {
	board1 := CreateEmptyBitBoard()
	board2 := CreateEmptyBitBoard()

	board1.board[0][0][0] = true
	board2.board[0][0][0] = false

	board1.board[0][0][1] = true
	board2.board[0][0][1] = true

	board1.board[0][0][2] = false
	board2.board[0][0][2] = false

	orBoard := Or(board1, board2)
	if !orBoard.board[0][0][0] || !orBoard.board[0][0][1] || orBoard.board[0][0][2] {
		t.Error("OR failed")
	}
}

func TestPlacePieceOnBoard(t *testing.T) {
	board := CreateEmptyBitBoard()

	board.PlacePieceOnBoard(0, 0, BLACK_QUEEN)
	if !board.board[0][0][BLACK_QUEEN] {
		t.Error("Piece hasn't been placed")
	}
}

func TestBitBoard_IsFieldEmpty(t *testing.T) {
	board := CreateEmptyBitBoard()

	for i, row := range board.board {
		for j := range row {
			if !board.isFieldEmpty(i, j) {
				t.Error("not all fields are empty")
			}
		}
	}

	board.PlacePieceOnBoard(0, 0, BLACK_PAWN)

	if board.isFieldEmpty(0, 0) {
		t.Error("Field shouldn't be empty")
	}
}

func TestBitBoard_IsFieldBlack(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_ROOK)
	board.PlacePieceOnBoard(1, 1, BLACK_QUEEN)

	if board.isFieldBlack(0, 0) || !board.isFieldWhite(0, 0) {
		t.Error("white field classified as black")
	}
	if !board.isFieldBlack(1, 1) || board.isFieldWhite(1, 1) {
		t.Error("black field classified as white")
	}

	if board.isFieldBlack(2, 2) || board.isFieldWhite(2, 2) {
		t.Error("Empty field got classified as either black or white")
	}
}

func TestBitBoard_IsFieldAvailable(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_PAWN)
	board.PlacePieceOnBoard(1, 1, BLACK_PAWN)

	if board.isFieldAvailable(0, 0, true) {
	    t.Error("field should not be available for white")
	}
	if !board.isFieldAvailable(0, 0, false) {
		t.Error("field should be available for black")
	}

	if board.isFieldAvailable(1, 1, false) {
		t.Error("field should not be available for black")
	}
	if !board.isFieldAvailable(1, 1, true) {
		t.Error("field should be available for white")
	}
}

func TestBitBoard_GetPieceOnField(t *testing.T) {
	board := CreateEmptyBitBoard()

	board.PlacePieceOnBoard(0, 0, WHITE_PAWN)
	if WHITE_PAWN != board.GetPieceOnField(0, 0) {
		t.Error("got invalid piece")
	}
	board.PlacePieceOnBoard(1, 1, BLACK_QUEEN)
	if BLACK_QUEEN != board.GetPieceOnField(1, 1) {
		t.Error("got invalid piece")
	}
}

func TestBitBoard_IsCheck(t *testing.T) {
	board := GetStartBoard()

	if board.IsCheck(true) || board.IsCheck(false) {
		t.Error("check in starting pos is invalid")
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, WHITE_KING)
	board.PlacePieceOnBoard(0, 0, BLACK_QUEEN)

	if !board.IsCheck(true) {
		t.Error("white should be in check")
	}
	if board.IsCheck(false) {
		t.Error("black should not be in check")
	}

	board.PlacePieceOnBoard(3, 3, WHITE_PAWN)

	if board.IsCheck(true) {
		t.Error("white should not be in check")
	}
}

func TestBitBoard_Copy(t *testing.T) {
	board := GetStartBoard()
	copiedBoard := board.Copy()

	for i, row := range board.board {
		for j := range row {
			if board.GetPieceOnField(i, j) != copiedBoard.GetPieceOnField(i, j) {
				t.Error("invalid copy")
			}
		}
	}
}

/**
func TestBitBoard_MovePiece(t *testing.T) {
	board := GetStartBoard()
	movedBoard := board.MovePiece(1, 1, 1, 3)

	if board.isFieldEmpty(1, 1) {
		t.Error("old board should not have been changed")
	}

	if movedBoard.isFieldEmpty(1, 3) {
		t.Error("move hasn't been executed on the new board")
	}
}
*/

func TestBitBoard_ToFEN(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	board := GetStartBoard()

	if board.ToFEN() != fen {
		t.Errorf("starting board doesn't produce correct fen, produced fen:\n%s", board.ToFEN())
	}

	board = FromFEN(fen)
	if board.ToFEN() != fen {
		t.Errorf("fen doesn't produce correct board, produced fen:\n%s", board.ToFEN())
	}
}

func TestBitBoard_Equals(t *testing.T) {
	b1 := GetStartBoard()
	b2 := GetStartBoard()

	if !b1.Equal(b2) {
		t.Error("boards should be equal")
	}

	b2 = CreateEmptyBitBoard()

	if b1.Equal(b2) {
		t.Error("boards should not be equal")
	}
}

func TestBitBoard_IsCheckMate(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_KING)
	board.PlacePieceOnBoard(1, 1, BLACK_QUEEN)
	board.PlacePieceOnBoard(2, 2, BLACK_BISHOP)

	if !board.IsCheckMate(true) {
		t.Error("Position is check mate")
	}

	if board.IsCheckMate(false) {
		t.Error("for black position is no check mate")
	}
}

func TestBitBoard_IsMoveValid(t *testing.T) {
	board := GetStartBoard()

	if !board.IsMoveValid(5, 1, 5, 3) {
		t.Error("pawn move should be valid")
	}

	if !board.IsMoveValid(5, 1, 5, 2) {
		t.Error("pawn move should be valid")
	}

	if board.IsMoveValid(0, 3, 5, 3) {
		t.Error("invalid move, no piece on this square")
	}

	if board.IsMoveValid(0, 6, 4, 4) {
		t.Error("invalid move, cant move black piece as white")
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_QUEEN)
	board.PlacePieceOnBoard(4, 0, WHITE_KING)
	board.PlacePieceOnBoard(4, 1, WHITE_ROOK)

	if board.IsMoveValid(4, 1, 0, 1) {
		t.Error("error, moving rook would result in self check")
	}

	if !board.IsMoveValid(4, 1, 4, 4) {
		t.Error("error, capturing the queen is valid")
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_QUEEN)
	board.PlacePieceOnBoard(1, 7, BLACK_KING)

	if !board.IsMoveValid(0, 0, 1, 0) {
		t.Error("error, checking black should be valid")
	}


}