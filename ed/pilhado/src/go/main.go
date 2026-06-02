package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var l, c int
	fmt.Scanf("%d %d", &l, &c)

	grid := make([][]rune, l)
	for i := range grid {
		grid[i] = make([]rune, c)
	}
	
}

