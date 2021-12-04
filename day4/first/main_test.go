package main

import (
	"testing"

	"github.com/advent-of-code-2021/input"
)

func TestParseBoard_With5by5Numbers_ShouldReturnValidBoard(t *testing.T) {
	in := input.Read("test_input.txt")

	// get a single board
	in = in[2:7]

	board := ParseBoard(in)

	exp := [5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != exp[i][j] {
				t.Errorf("Unexpected board entry at %d,%d. Got %d, expected %d", i, j, board[i][j], exp[i][j])
			}
		}
	}
}

func TestParseBoards_WithMultipleBoards_ShouldReturnMultipleValidBoards(t *testing.T) {
	in := input.Read("test_input.txt")

	// skip over the first 2 lines, because they consist of the random numbers and a newline
	in = in[2:]

	boards := ParseBoards(in)

	exp := [][5][5]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}

	for bi := 0; bi < 3; bi++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if boards[bi][i][j] != exp[bi][i][j] {
					t.Errorf("Unexpected board entry at %d,%d. Got %d, expected %d", i, j, boards[bi][i][j], exp[bi][i][j])
				}
			}
		}
	}
}

func TestCheckBoardForWin_WithWinningRow_ShouldReturnTrue(t *testing.T) {
	board := [5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{-1, -1, -1, -1, -1},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	result := CheckBoardForWin(board)

	if result == false {
		t.Errorf("Wrong board win returned. Expected true, got false")
	}
}

func TestCheckBoardForWin_WithWinningColumn_ShouldReturnTrue(t *testing.T) {
	board := [5][5]int{
		{22, 13, -1, 11, 0},
		{8, 2, -1, 4, 24},
		{21, 9, -1, 16, 7},
		{6, 10, -1, 18, 5},
		{1, 12, -1, 15, 19},
	}

	result := CheckBoardForWin(board)

	if result == false {
		t.Errorf("Wrong board win returned. Expected true, got false")
	}
}

func TestCheckBoardForWin_WithNoWin_ShouldReturnFalse(t *testing.T) {
	board := [5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	result := CheckBoardForWin(board)

	if result == true {
		t.Errorf("Wrong board win returned. Expected false, got true")
	}
}

func TestCheckBoardForWin_WithPartialWin_ShouldReturnFalse(t *testing.T) {
	board := [5][5]int{
		{22, -1, 17, 11, 0},
		{8, -1, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{-1, -1, 20, 15, 19},
	}

	result := CheckBoardForWin(board)

	if result == true {
		t.Errorf("Wrong board win returned. Expected false, got true")
	}
}

func TestGetBoardScore_WithValidInput_ShouldReturnCorrectScore(t *testing.T) {
	board := [5][5]int{
		{-1, -1, -1, -1, -1},
		{10, 16, 15, -1, 19},
		{18, 8, -1, 26, 20},
		{22, -1, 13, 6, -1},
		{-1, -1, 12, 3, -1},
	}

	score := GetBoardScore(24, board)

	exp := 4512

	if score != exp {
		t.Errorf("Unexpected winning board score. Got %d, expected %d", score, exp)
	}
}

func TestCompleteGame_WithTestInput_ShouldReturnCorrectScore(t *testing.T) {
	in := input.Read("test_input.txt")

	nums, boards := SetupGame(in)

	score := PlayGame(nums, boards)

	exp := 4512

	if score != exp {
		t.Errorf("Unexpected endscore. Got %d, expected %d", score, exp)
	}
}
