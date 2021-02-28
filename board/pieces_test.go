package board

import (
	"math"
	"testing"
)

func TestGetMovementMatrixPawn(t *testing.T) {
	board := GetStartBoard()

	var movementMatrix BitBoard

	movementMatrix = WHITE_PAWN.GetMovementMatrix(&board, 0, 1, false)

	for i, row := range movementMatrix.board {
		for j, field := range row {
			for k, value := range field {
				if value != ((i == 0) && (k == int(WHITE_PAWN)) && ((j == 2) || (j == 3))) {
					t.Error("movement matrix is wrong")
				}
			}
		}
	}

	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6, false)

	for i, row := range movementMatrix.board {
		for j, field := range row {
			for k, value := range field {
				if value != ((i == 0) && (k == int(BLACK_PAWN)) && ((j == 5) || (j == 4))) {
					t.Error("movement matrix is wrong")
				}
			}
		}
	}

	board.PlacePieceOnBoard(0, 5, BLACK_PAWN)
	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6, false)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix should be empty")
			}
		}
	}

	board.PlacePieceOnBoard(1, 5, WHITE_PAWN)
	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6, false)
	for i, row := range movementMatrix.board {
		for j := range row {
			if (i == 1 && j == 5) && movementMatrix.isFieldEmpty(i, j) {
				t.Error("invalid movement matrix, pawn should be able to move and capture")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, WHITE_PAWN)
	movementMatrix = WHITE_PAWN.GetMovementMatrix(&board, 4, 4, true)
	if !movementMatrix.isFieldEmpty(4, 6) {
		t.Error("pawn not allowed to move 2 squares")
	}
}

func TestPiece_GetMovementMatrixRook(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_ROOK.GetMovementMatrix(&board, 7, 7, false)

	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix wrong, rook is blocked in every direction")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_ROOK)
	movementMatrix = WHITE_ROOK.GetMovementMatrix(&board, 0, 0, false)

	for i := 1; i < 8; i++ {
		if !movementMatrix.board[0][i][WHITE_ROOK] || !movementMatrix.board[i][0][WHITE_ROOK] {
			t.Error("movement matrix wrong for rook in corner")
		}
	}

	board.PlacePieceOnBoard(0, 4, WHITE_ROOK)
	movementMatrix = WHITE_ROOK.GetMovementMatrix(&board, 0, 0, false)
	for i := 1; i < 8; i++ {
		if i < 4 {
			if !movementMatrix.board[0][i][WHITE_ROOK] {
				t.Error("movement matrix wrong for rook in corner blocked by piece")
			}
		} else {
			if movementMatrix.board[0][i][WHITE_ROOK] {
				t.Error("movement matrix wrong for rook in corner blocked by piece")
			}

		}
	}
}

func TestPiece_GetMovementMatrixKnight(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_KNIGHT.GetMovementMatrix(&board, 1, 7, false)

	for i, row := range movementMatrix.board {
		for j := range row {
			if !(movementMatrix.isFieldEmpty(i, j) == !((j == 5) && ((i == 0) || (i == 2)))) {
				t.Error("movement matrix for knight wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_KNIGHT)

	movementMatrix = BLACK_KNIGHT.GetMovementMatrix(&board, 4, 4, false)
	if movementMatrix.isFieldEmpty(6, 3) || movementMatrix.isFieldEmpty(6, 5) ||
		movementMatrix.isFieldEmpty(2, 3) || movementMatrix.isFieldEmpty(2, 5) ||
		movementMatrix.isFieldEmpty(3, 6) || movementMatrix.isFieldEmpty(5, 6) ||
		movementMatrix.isFieldEmpty(3, 2) || movementMatrix.isFieldEmpty(5, 2) {
		t.Error("Movement matrix wrong for knight in mid")
	}
}

func TestPiece_GetMovementMatrixBishop(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_BISHOP.GetMovementMatrix(&board, 2, 7, false)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix for bishop in start pos wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_BISHOP)
	movementMatrix = BLACK_BISHOP.GetMovementMatrix(&board, 4, 4, false)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if ((i == j) || (j == (8 - i))) && (i != 4) && (j != 4) {
		    	if movementMatrix.isFieldEmpty(i, j) {
					t.Error("movement matrix invalid, piece should be able to move here")
				}
			} else if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix invalid, piece should not be able to move here")
			}
		}
	}

	board.PlacePieceOnBoard(3, 3, BLACK_BISHOP)
	movementMatrix = BLACK_BISHOP.GetMovementMatrix(&board, 4, 4, false)

	for i := 0; i < 8; i ++ {
		if i < 5 {
			if !movementMatrix.isFieldEmpty(i, i) {
				t.Error("movement matrix invalid, blocked piece should not be able to move here")
			}
		} else if movementMatrix.isFieldEmpty(i, i) {
			t.Error("movement matrix invalid, blocked piece should be able to move here")
		}
	}
}
func TestPiece_GetMovementMatrixQueen(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_QUEEN.GetMovementMatrix(&board, 3, 7, false)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix for queen in start pos wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_QUEEN)
	movementMatrix = BLACK_QUEEN.GetMovementMatrix(&board, 4, 4, false)

	for i, row := range movementMatrix.board {
		for j := range row {
			rookField := (i == 4) || (j == 4)
			bishopField := ((i == j) || (j == (8 - i))) && (i != 4) && (j != 4)
			if (rookField || bishopField) && !((i == 4) && (j == 4)) {
				if movementMatrix.isFieldEmpty(i, j) {
					t.Error("invalid movement matrix, piece should be able to move here")
				}
			} else if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("invalid movement matrix, piece should not be able to move here")
			}
		}
	}
}

func TestPiece_GetMovementMatrixKing(t *testing.T) {
    board := GetStartBoard()
    movementMatrix := BLACK_KING.GetMovementMatrix(&board, 4, 7, false)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("movement matrix should be empty")
		    }
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_KING)
	movementMatrix = BLACK_KING.GetMovementMatrix(&board, 4, 4, false)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if int(math.Abs(float64(4 - i)) - 1) <= 0 && int(math.Abs(float64(4 - j)) - 1) <= 0 { // go needs floats for abs?
		    	if movementMatrix.isFieldEmpty(i, j) && !((i == 4) && (j == 4)) {
		    		t.Error("invalid movement matrix, piece should be able to move here")
				}
			} else if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("invalid movement matrix, piece should not be able to move here")
			}
		}
	}

	board.PlacePieceOnBoard(4, 2, WHITE_KING)
	movementMatrix = WHITE_KING.GetMovementMatrix(&board, 4, 2, false)
	if !(movementMatrix.isFieldEmpty(3, 3) && movementMatrix.isFieldEmpty(4, 3) &&
		movementMatrix.isFieldEmpty(5, 3)) {
		t.Error("Kings have to stay one square apart")
	}

}

func TestPiece_IsColor(t *testing.T) {
	if !BLACK_KING.IsBlack() {
		t.Error("Black king incorrectly classified")
	}
	if BLACK_KING.IsWhite() {
		t.Error("Black king incorrectly classified")
	}
	if !WHITE_KING.IsWhite() {
		t.Error("White king incorrectly classified")
	}
	if WHITE_KING.IsBlack() {
		t.Error("White king incorrectly classified")
	}
}

func TestPiece_MoveInCheck(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, WHITE_KING)
	board.PlacePieceOnBoard(0, 0, BLACK_QUEEN)

	movementMatrix := WHITE_KING.GetMovementMatrix(&board, 4, 4, false)

	if !movementMatrix.isFieldEmpty(3, 3) || !movementMatrix.isFieldEmpty(5, 5) {
		t.Error("king should not be able to move into check")
	}

	if movementMatrix.isFieldEmpty(3, 4) || movementMatrix.isFieldEmpty(3, 5) {
		t.Error("king should be able to move here")
	}

	board.PlacePieceOnBoard(1, 0, WHITE_BISHOP)

	movementMatrix = WHITE_BISHOP.GetMovementMatrix(&board, 1, 0, false)

	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.isFieldEmpty(i, j) {
				t.Error("bishop should not be able to move because king is in check")
			}
		}
	}
}

func TestPiece_MovePawnCheckKing(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_QUEEN)
	board.PlacePieceOnBoard(4, 0, WHITE_KING)
	board.PlacePieceOnBoard(4, 1, WHITE_PAWN)
	board.PlacePieceOnBoard(3, 2, BLACK_PAWN)


	movementMatrix := WHITE_PAWN.GetMovementMatrix(&board, 4, 1, false)
	if !movementMatrix.isFieldEmpty(3, 2) {
		t.Error("this move would result in self check")
	}
}

func TestPiece_MovePawnEnPassant(t *testing.T) {
	board := CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, WHITE_PAWN)
	board.PlacePieceOnBoard(3, 4, BLACK_PAWN)
	board.SetEnPassant(3, 5)


	movementMatrix := WHITE_PAWN.GetMovementMatrix(&board, 4, 4, false)
	if movementMatrix.isFieldEmpty(3, 5) {
		t.Error("Error, this move should be legal due to en passant")
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 3, WHITE_PAWN)
	board.PlacePieceOnBoard(3, 3, BLACK_PAWN)
	board.SetEnPassant(4, 2)


	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 3, 3, false)
	if movementMatrix.isFieldEmpty(4, 2) {
		t.Error("Error, this move should be legal due to en passant")
	}

	board.PlacePieceOnBoard(3, 1, WHITE_PAWN)
	movementMatrix = WHITE_PAWN.GetMovementMatrix(&board, 3, 1, false)
	if !movementMatrix.isFieldEmpty(4, 2) {
		t.Error("Error, this move should not be legal due to en passant")

	}
}
