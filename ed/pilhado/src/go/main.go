package main

import (
	//"bufio"
	"fmt"
	//"os"
)

type cord struct {
	l, c int
}

func Labirinto(grid *[][]rune, Ini cord, Fin cord) bool {
	if Ini.l < 0 || Ini.l >= len(*grid) || Ini.c < 0 || Ini.c >= len((*grid)[0]) {
		return false
	}
	if Ini.l == Fin.l && Ini.c == Fin.c {
		return true
	}
	if (*grid)[Ini.l][Ini.c] == ' ' {
		(*grid)[Ini.l][Ini.c] = '.'
	}

	if !Labirinto(grid, cord{Ini.l - 1, Ini.c}, Fin) &&
		!Labirinto(grid, cord{Ini.l + 1, Ini.c}, Fin) &&
		!Labirinto(grid, cord{Ini.l, Ini.c - 1}, Fin) &&
		!Labirinto(grid, cord{Ini.l, Ini.c + 1}, Fin) {
		(*grid)[Ini.l][Ini.c] = ' '
		return false
	}
	return true

	
}

func main() {
	var l, c int
	fmt.Scanf("%d %d", &l, &c)

	grid := make([][]rune, l)
	for i := range grid {
		grid[i] = make([]rune, c)
	}
	Ini := cord{-1, -1}
	Fin := cord{-1, -1}
	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			fmt.Scanf("%c", &grid[i][j])
			if grid[i][j] == 'I' {
				Ini = cord{i, j}
			}	
			if grid[i][j] == 'F' {
				Fin = cord{i, j}
			}
		}
		fmt.Scanf("\n")
	}
	Labirinto(&grid, Ini, Fin)

	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Printf("\n")
	}
	
}

