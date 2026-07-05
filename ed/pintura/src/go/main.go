package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	//
	corOri:= image[sr][sc]
	if corOri == color{
		return image
	}

	var pintar func(sr int, sc int)
	pintar = func(sr int, sc int) {
		if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
			return
		}
		if image[sr][sc] != corOri {
			return
		}
		image[sr][sc] = color
		pintar(sr+1, sc)
		pintar(sr-1, sc)
		pintar(sr, sc+1)
		pintar(sr, sc-1)
	}

	pintar(sr, sc)
	return image
}
	

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
