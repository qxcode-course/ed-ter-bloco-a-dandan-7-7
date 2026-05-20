package main

import (
	"bufio"
	"fmt"
	"os"
)

func search(board [][]byte, i, j, bordai, bordaj int) {

	if i == bordai || j == bordaj || board[i][j] == 'X' || i < 0 || j < 0{
		return
	} 
	if board[i][j] == 'O'{
		board[i][j] = 'X'
	}
	 search(board, i+1, j, bordai, bordaj) 
	 search(board, i-1, j, bordai, bordaj) 
	 search(board, i, j+1, bordai, bordaj) 
	 search(board, i, j-1, bordai, bordaj)



}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	_ = board

	bordai := len(board)
	bordaj := len(board[0])

	search(board, 1, 1, bordai, bordaj)

	// for i := 0; i < bordai; i++ {
	// 	for j := 0; j < bordaj; j++ {
	// 		if board[i][j] == 'O' {
	// 			board[i][j] = 'X'
	// 		}
	// 	}
	// }
	
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
