package main

import (
	"bufio"
	"fmt"
	"os"
)

func afundarIlha(grid [][]byte, i, j, tami, tamj int) {
    
    if i < 0 || i >= tami || j < 0 || j >= tamj || grid[i][j] != '1' {
        return
    }
   
    grid[i][j] = '0'

    afundarIlha(grid, i+1, j, tami, tamj) 
    afundarIlha(grid, i-1, j, tami, tamj) 
    afundarIlha(grid, i, j+1, tami, tamj) 
    afundarIlha(grid, i, j-1, tami, tamj) 
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    cont := 0 
    tami := len(grid)
    tamj := len(grid[0])
    
    for i := 0; i < tami; i++ {
        for j := 0; j < tamj; j++ {
           
            if grid[i][j] == '1' {
                cont++
               
                afundarIlha(grid, i, j, tami, tamj)
            }
        }
    }

    return cont
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
