package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	nums, boards := SetupGame(in)

	score := PlayGame(nums, boards)

	fmt.Printf("Final score of the winning board is %d\n", score)
}

func PlayGame(nums []int, boards [][5][5]int) int {
	for _, num := range nums {
		for n := 0; n < len(boards); n++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if boards[n][i][j] == num {
						// mark the number
						boards[n][i][j] = -1
					}
				}
			}

			if CheckBoardForWin(boards[n]) {
				if len(boards) == 1 {
					// this is the last board, so get its score
					return GetBoardScore(num, boards[n])
				} else {
					// remove the board from the game, because we want to
					// continue playing until the last board wins
					boards = append(boards[:n], boards[n+1:]...)
				}
			}
		}
	}

	return -1
}

func SetupGame(input []string) ([]int, [][5][5]int) {
	randNumsRaw := strings.Split(input[0], ",")

	randNums := []int{}

	for _, r := range randNumsRaw {
		num, _ := strconv.Atoi(r)

		randNums = append(randNums, num)
	}

	boards := ParseBoards(input[2:])

	return randNums, boards
}
