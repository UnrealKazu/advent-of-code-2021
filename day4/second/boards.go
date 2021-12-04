package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// ParseBoards takes an input consisting only of raw boards, and converts them into matrices
func ParseBoards(input []string) [][5][5]int {
	low := 0
	up := 5

	boards := [][5][5]int{}

	for up <= len(input) {
		b := input[low:up]

		boards = append(boards, ParseBoard(b))
		low += 6
		up += 6
	}

	return boards
}

// ParseBoard will take a slice of 5 rows and parse it into a multidimensional array
func ParseBoard(input []string) [5][5]int {

	board := [5][5]int{}

	r := regexp.MustCompile(` *([0-9]+)`)

	for i := 0; i < 5; i++ {
		ents := r.FindAllStringSubmatch(input[i], -1)

		for j, v := range ents {
			var err error
			num, err := strconv.Atoi(v[1])

			if err != nil {
				fmt.Printf("Unexpected board entry: %v", err)
				panic("Cannot parse board. Exiting")
			}

			board[i][j] = num
		}
	}

	return board
}

// CheckBoardForWin traverses the given board, and checks if any row or column has a bingo
func CheckBoardForWin(board [5][5]int) bool {
	for i := 0; i < 5; i++ {
		winRow := true

		for j := 0; j < 5; j++ {
			if board[i][j] > -1 {
				winRow = false
				break
			}
		}

		if winRow {
			return true
		}
	}

	for j := 0; j < 5; j++ {
		winColumn := true

		for i := 0; i < 5; i++ {
			if board[i][j] > -1 {
				winColumn = false
				break
			}
		}

		if winColumn {
			return true
		}
	}

	return false
}

// GetBoardScore returns the score for the given board
func GetBoardScore(winNum int, board [5][5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] > -1 {
				sum += board[i][j]
			}
		}
	}

	return winNum * sum
}
