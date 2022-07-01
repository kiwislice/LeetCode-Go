package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0130(t *testing.T) {
	assert := assert.New(t)

	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	solve(board)
	assert.Equal([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'}}, board)

	board = [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'}}
	printBoard(board)
	solve(board)
	printBoard(board)
	assert.Equal([][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'}}, board)

}

func printBoard(board [][]byte) {
	fmt.Print("\n")
	for i := range board {
		for j := range board[i] {
			fmt.Print(string(board[i][j]))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
