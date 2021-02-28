package board

import (
    game2 "terrible_chess_computer/game"
    "testing"
)

func TestGame_pop_push(t *testing.T) {
	game := game2.Game{}

	board := GetStartBoard()
	game.push(board)
	if !board.Equal(game.pop()) {
		t.Error("popped board not equal to original")
	}
}

func Test_InitGame(t *testing.T) {
	game := game2.InitGame()
	board := GetStartBoard()
	if !board.Equal(game.pop()) {
		t.Error("game not initialized, position unequal to start position")
	}
}
