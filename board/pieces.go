package board

type Piece int

const (
	BLACK_PAWN Piece = 0
	BLACK_ROOK Piece = 1
	BLACK_KNIGHT Piece = 2
	BLACK_BISHOP Piece = 3
	BLACK_QUEEN Piece = 4
	BLACK_KING Piece = 5
	WHITE_PAWN Piece = 6
	WHITE_ROOK Piece = 7
	WHITE_KNIGHT Piece = 8
	WHITE_BISHOP Piece = 9
	WHITE_QUEEN Piece = 10
	WHITE_KING Piece = 11
	NO_PIECE Piece = 12
)

func (piece Piece) IsBlack() bool {
	return piece < 6
}

func (piece Piece) IsWhite() bool {
	return piece >= 6 && piece != NO_PIECE
}

func (piece Piece) IsNone() bool {
	return piece == NO_PIECE
}

func (piece Piece) GetMovementMatrix(board *BitBoard, x, y int, allowCheck bool) BitBoard {
    movementMatrix := CreateEmptyBitBoard()
	switch piece {
	case BLACK_PAWN:
		movementMatrix = getPawnMatrix(board, x, y, false)
	case WHITE_PAWN:
		movementMatrix = getPawnMatrix(board, x, y, true)
	case BLACK_ROOK:
		movementMatrix = getRookMatrix(board, x, y, false)
	case WHITE_ROOK:
		movementMatrix = getRookMatrix(board, x, y, true)
	case BLACK_KNIGHT:
		movementMatrix = getKnightMatrix(board, x, y, false)
	case WHITE_KNIGHT:
		movementMatrix = getKnightMatrix(board, x, y, false)
	case BLACK_BISHOP:
		movementMatrix = getBishopMatrix(board, x, y, false)
	case WHITE_BISHOP:
		movementMatrix = getBishopMatrix(board, x, y, true)
	case BLACK_QUEEN:
		movementMatrix = getQueenMatrix(board, x, y, false)
	case WHITE_QUEEN:
		movementMatrix = getQueenMatrix(board, x, y, true)
	case BLACK_KING:
		movementMatrix = getKingMatrix(board, x, y, false)
	case WHITE_KING:
		movementMatrix = getKingMatrix(board, x, y, true)
	}

	if !allowCheck {
		removeInvalidMoves(&movementMatrix, board, x, y, piece.IsWhite())
	}

	return movementMatrix
}

func (piece Piece) GetNotation() string {
	symbols := [13]string{"p", "r", "n", "b", "q", "k", "P", "R", "N", "B", "Q", "K", "-"}
	return symbols[piece]
}

func GetPieceByNotation(s string) Piece {
	for i := 0; i < 13; i++ {
		if Piece(i).GetNotation() == s {
			return Piece(i)
		}
	}

	return NO_PIECE
}

func removeInvalidMoves(matrix *BitBoard, board *BitBoard, x, y int, white bool) {
	for i, row := range matrix.board {
		for j := range row {
			if !matrix.isFieldEmpty(i, j) {
				if board.doesMoveResultInCheck(x, y, i, j, white) {
					matrix.PlacePieceOnBoard(i, j, NO_PIECE)
				}
			}
		}
	}
}

func getPawnMatrix(board *BitBoard, x, y int, white bool) BitBoard {
    var piece Piece
    var direction int
    if white {
    	direction = 1
    	piece = WHITE_PAWN
	} else {
		direction = -1
		piece = BLACK_PAWN
	}
	matrix := CreateEmptyBitBoard()
	if board.isFieldAvailable(x, y + direction, white) {
		matrix.board[x][y + direction][piece] = true
		if (white && (y == 1)) || (!white && (y == 6)) {
			matrix.board[x][y + 2 * direction][piece] = board.isFieldAvailable(x, y + 2 * direction, white)
		}
	}

	if x > 0 {
		if board.isFieldAvailable(x - 1, y + direction, white) && !board.isFieldEmpty(x - 1, y + direction) ||
			(((white && y > 3) || (!white && y < 4)) && ((board.GetEnPassant()[0] == x - 1) &&
				(board.GetEnPassant()[1] == y + direction))) {
			matrix.PlacePieceOnBoard(x - 1, y + direction, piece)
		}
	}
	if x < 7 {
		if board.isFieldAvailable(x + 1, y + direction, white) && !board.isFieldEmpty(x + 1, y + direction) ||
			(((white && y > 3) || (!white && y < 4)) && ((board.GetEnPassant()[0] == x + 1) &&
				(board.GetEnPassant()[1] == y + direction))) {
			matrix.PlacePieceOnBoard(x + 1, y + direction, piece)
		}
	}


return matrix
}

func getRookMatrix(board *BitBoard, x, y int, white bool) BitBoard {
	var piece Piece
	if white {
		piece = WHITE_ROOK
	} else {
		piece = BLACK_ROOK
	}
	matrix := CreateEmptyBitBoard()

	for i := x + 1; i < 8; i++ {
		if !board.isFieldEmpty(i, y) {
			if board.isFieldAvailable(i, y, white) {
				matrix.PlacePieceOnBoard(i, y, piece)
			}
			break
		}
		matrix.PlacePieceOnBoard(i, y, piece)
	}

	for i := x - 1; i >= 0; i-- {
		if !board.isFieldEmpty(i, y) {
			if board.isFieldAvailable(i, y, white) {
				matrix.PlacePieceOnBoard(i, y, piece)
			}
			break
		}
		matrix.PlacePieceOnBoard(i, y, piece)
	}

	for j := y + 1; j < 8; j++ {
		if !board.isFieldEmpty(x, j) {
			if board.isFieldAvailable(x, j, white) {
				matrix.PlacePieceOnBoard(x, j, piece)
			}
			break
		}
		matrix.PlacePieceOnBoard(x, j, piece)
	}

	for j := y - 1; j >= 0; j-- {
		if !board.isFieldEmpty(x, j) {
			if board.isFieldAvailable(x, j, white) {
				matrix.PlacePieceOnBoard(x, j, piece)
			}
			break
		}
		matrix.PlacePieceOnBoard(x, j, piece)
	}

	return matrix
}

func getKnightMatrix(board *BitBoard, x, y int, white bool) BitBoard {
	var piece Piece
	if white {
		piece = WHITE_ROOK
	} else {
		piece = BLACK_ROOK
	}

	matrix := CreateEmptyBitBoard()

	if x + 2 < 8 {
		if y + 1 < 8 {
			matrix.board[x + 2][y + 1][piece] = board.isFieldAvailable(x + 2, y + 1, white)
		}
		if y - 1 >= 0 {
			matrix.board[x + 2][y - 1][piece] = board.isFieldAvailable(x + 2, y - 1, white)
		}
	}

	if x - 2 >= 0 {
		if y + 1 < 8 {
			matrix.board[x - 2][y + 1][piece] = board.isFieldAvailable(x - 2, y + 1, white)
		}
		if y - 1 >= 0 {
			matrix.board[x - 2][y - 1][piece] = board.isFieldAvailable(x - 2, y - 1, white)
		}
	}

	if y + 2 < 8 {
		if x + 1 < 8 {
			matrix.board[x + 1][y + 2][piece] = board.isFieldEmpty(x + 1, y + 2)
		}
		if x - 1 >= 0 {
			matrix.board[x - 1][y + 2][piece] = board.isFieldEmpty(x - 1, y + 2)
		}
	}

	if y - 2 >= 0 {
		if x + 1 < 8 {
			matrix.board[x + 1][y - 2][piece] = board.isFieldEmpty(x + 1, y - 2)
		}
		if x - 1 >= 0 {
			matrix.board[x - 1][y - 2][piece] = board.isFieldEmpty(x - 1, y - 2)
		}
	}

	return matrix
}

func getBishopMatrix(board *BitBoard, x, y int, white bool) BitBoard {
	var piece Piece
	if white {
		piece = WHITE_BISHOP
	} else {
		piece = BLACK_BISHOP
	}
	movementMatrix := CreateEmptyBitBoard()
	b1 := y - x
	b2 := y + x

	for i := x + 1; i < 8; i++ {
		yi := b1 + i
		if  yi >= 8 || yi < 0 {
			break
		}
		if !board.isFieldEmpty(i, yi) {
			if board.isFieldAvailable(i, yi, white) {
				movementMatrix.PlacePieceOnBoard(i, yi, piece)
			}
			break
		}
		movementMatrix.PlacePieceOnBoard(i, yi, piece)
	}

	for i := x + 1; i < 8; i++ {
		yi := b2 - i
		if  yi >= 8 || yi < 0 {
			break
		}
		if !board.isFieldEmpty(i, yi) {
			if board.isFieldAvailable(i, yi, white) {
				movementMatrix.PlacePieceOnBoard(i, yi, piece)
			}
			break
		}
		movementMatrix.PlacePieceOnBoard(i, yi, piece)
	}

	for i := x - 1; i >= 0; i-- {
		yi := b1 + i
		if  yi >= 8 || yi < 0 {
			break
		}
		if !board.isFieldEmpty(i, yi) {
			if board.isFieldAvailable(i, yi, white) {
				movementMatrix.PlacePieceOnBoard(i, yi, piece)
			}
			break
		}
		movementMatrix.PlacePieceOnBoard(i, yi, piece)
	}

	for i := x - 1; i >= 0; i-- {
		yi := b2 - i
		if  yi >= 8 || yi < 0 {
			break
		}
		if !board.isFieldEmpty(i, yi) {
			if board.isFieldAvailable(i, yi, white) {
				movementMatrix.PlacePieceOnBoard(i, yi, piece)
			}
			break
		}
		movementMatrix.PlacePieceOnBoard(i, yi, piece)
	}

	return movementMatrix
}

func getKingMatrix(board *BitBoard, x, y int, white bool) BitBoard {
	var piece Piece
	if white {
		piece = WHITE_KING
	} else {
		piece = BLACK_KING
	}
	movementMatrix := CreateEmptyBitBoard()

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			xi := x + i
			yi := y + j
			if !((i == x) && (j == y)) && xi < 8 && xi >= 0 && yi < 8 && yi >= 0 {
				if board.isFieldAvailable(xi, yi, white) {
					movementMatrix.PlacePieceOnBoard(xi, yi, piece)
				}
			}
		}
	}

	return movementMatrix
}

func getQueenMatrix(board *BitBoard, x, y int, white bool) BitBoard {
	var piece Piece
	var movementMatrixBishop BitBoard
	var movementMatrixRook BitBoard
	if white {
		piece = WHITE_QUEEN
		movementMatrixBishop = WHITE_BISHOP.GetMovementMatrix(board, x, y, true)
		movementMatrixRook = WHITE_ROOK.GetMovementMatrix(board, x, y, true)
	} else {
		movementMatrixBishop = BLACK_BISHOP.GetMovementMatrix(board, x, y, true)
		movementMatrixRook = BLACK_ROOK.GetMovementMatrix(board, x, y, true)
		piece = BLACK_QUEEN
	}

	movementMatrix := CreateEmptyBitBoard()
	for i, row := range movementMatrix.board {
		for j := range row {
		    if !movementMatrixRook.isFieldEmpty(i, j) || !movementMatrixBishop.isFieldEmpty(i, j) {
		    	movementMatrix.PlacePieceOnBoard(i, j, piece)
			}
		}
	}


	return movementMatrix
}
