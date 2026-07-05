package main

import (
	"bufio"
	"fmt"
	"os"
)

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	//

	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}

	count := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {

				if !(i+1 < len(board)) && !(j+1 < len(board[0])) {
					count++
					board[i][j] = '.'
				} else if !(board[i+1][j] == 'X') && !(board[i][j+1] == 'X') {
					count++
					board[i][j] = '.'
				}
				if j+1 < len(board[0]) && i+1 < len(board) {
					if board[i][j+1] == 'X' {
						for j < len(board[0]) && board[i][j] == 'X' {
							board[i][j] = '.'
							j++
						}
					} else if board[i+1][j] == 'X' {
						for i < len(board) && board[i][j] == 'X' {
							board[i][j] = '.'
							i++
						}
					}
				}
				
			}
		}
	}


	
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
