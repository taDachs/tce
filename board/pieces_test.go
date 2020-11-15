package board

import (
	"math"
	"testing"
)

func TestGetMovementMatrixPawn(t *testing.T) {
	board := GetStartBoard()

	var movementMatrix BitBoard

	movementMatrix = WHITE_PAWN.GetMovementMatrix(&board, 0, 1)

	for i, row := range movementMatrix.board {
		for j, field := range row {
			for k, value := range field {
				if value != ((i == 0) && (k == int(WHITE_PAWN)) && ((j == 2) || (j == 3))) {
					t.Error("movement matrix is wrong")
				}
			}
		}
	}

	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6)

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
	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix should be empty")
			}
		}
	}

	board.PlacePieceOnBoard(1, 5, WHITE_PAWN)
	movementMatrix = BLACK_PAWN.GetMovementMatrix(&board, 0, 6)
	for i, row := range movementMatrix.board {
		for j := range row {
			if (i == 1 && j == 5) && movementMatrix.IsFieldEmpty(i, j) {
				t.Error("invalid movement matrix, pawn should be able to move and capture")
			}
		}
	}
}

func TestPiece_GetMovementMatrixRook(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_ROOK.GetMovementMatrix(&board, 7, 7)

	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix wrong, rook is blocked in every direction")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(0, 0, WHITE_ROOK)
	movementMatrix = WHITE_ROOK.GetMovementMatrix(&board, 0, 0)

	for i := 1; i < 8; i++ {
		if !movementMatrix.board[0][i][WHITE_ROOK] || !movementMatrix.board[i][0][WHITE_ROOK] {
			t.Error("movement matrix wrong for rook in corner")
		}
	}

	board.PlacePieceOnBoard(0, 4, WHITE_ROOK)
	movementMatrix = WHITE_ROOK.GetMovementMatrix(&board, 0, 0)
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

	movementMatrix := BLACK_KNIGHT.GetMovementMatrix(&board, 1, 7)

	for i, row := range movementMatrix.board {
		for j := range row {
			if !(movementMatrix.IsFieldEmpty(i, j) == !((j == 5) && ((i == 0) || (i == 2)))) {
				t.Error("movement matrix for knight wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_KNIGHT)

	movementMatrix = BLACK_KNIGHT.GetMovementMatrix(&board, 4, 4)
	if movementMatrix.IsFieldEmpty(6, 3) || movementMatrix.IsFieldEmpty(6, 5) ||
		movementMatrix.IsFieldEmpty(2, 3) || movementMatrix.IsFieldEmpty(2, 5) ||
		movementMatrix.IsFieldEmpty(3, 6) || movementMatrix.IsFieldEmpty(5, 6) ||
		movementMatrix.IsFieldEmpty(3, 2) || movementMatrix.IsFieldEmpty(5, 2) {
		t.Error("Movement matrix wrong for knight in mid")
	}
}

func TestPiece_GetMovementMatrixBishop(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_BISHOP.GetMovementMatrix(&board, 2, 7)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix for bishop in start pos wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_BISHOP)
	movementMatrix = BLACK_BISHOP.GetMovementMatrix(&board, 4, 4)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if ((i == j) || (j == (8 - i))) && (i != 4) && (j != 4) {
		    	if movementMatrix.IsFieldEmpty(i, j) {
					t.Error("movement matrix invalid, piece should be able to move here")
				}
			} else if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix invalid, piece should not be able to move here")
			}
		}
	}

	board.PlacePieceOnBoard(3, 3, BLACK_BISHOP)
	movementMatrix = BLACK_BISHOP.GetMovementMatrix(&board, 4, 4)

	for i := 0; i < 8; i ++ {
		if i < 5 {
			if !movementMatrix.IsFieldEmpty(i, i) {
				t.Error("movement matrix invalid, blocked piece should not be able to move here")
			}
		} else if movementMatrix.IsFieldEmpty(i, i) {
			t.Error("movement matrix invalid, blocked piece should be able to move here")
		}
	}
}
func TestPiece_GetMovementMatrixQueen(t *testing.T) {
	board := GetStartBoard()

	movementMatrix := BLACK_QUEEN.GetMovementMatrix(&board, 3, 7)
	for i, row := range movementMatrix.board {
		for j := range row {
			if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix for queen in start pos wrong")
			}
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_QUEEN)
	movementMatrix = BLACK_QUEEN.GetMovementMatrix(&board, 4, 4)

	for i, row := range movementMatrix.board {
		for j := range row {
			rookField := (i == 4) || (j == 4)
			bishopField := ((i == j) || (j == (8 - i))) && (i != 4) && (j != 4)
			if (rookField || bishopField) && !((i == 4) && (j == 4)) {
				if movementMatrix.IsFieldEmpty(i, j) {
					t.Error("invalid movement matrix, piece should be able to move here")
				}
			} else if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("invalid movement matrix, piece should not be able to move here")
			}
		}
	}
}

func TestPiece_GetMovementMatrixKing(t *testing.T) {
    board := GetStartBoard()
    movementMatrix := BLACK_KING.GetMovementMatrix(&board, 4, 7)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("movement matrix should be empty")
		    }
		}
	}

	board = CreateEmptyBitBoard()
	board.PlacePieceOnBoard(4, 4, BLACK_KING)
	movementMatrix = BLACK_KING.GetMovementMatrix(&board, 4, 4)

	for i, row := range movementMatrix.board {
		for j := range row {
		    if (math.Abs(float64(4 - i)) - 1) < 0.001 && (math.Abs(float64(4 - j)) - 1) < 0.001 {
		    	if movementMatrix.IsFieldEmpty(i, j) && !((i == 4) && (j == 4)) {
		    		t.Error("invalid movement matrix, piece should be able to move here")
				}
			} else if !movementMatrix.IsFieldEmpty(i, j) {
				t.Error("invalid movement matrix, piece should not be able to move here")
			}
		}
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
