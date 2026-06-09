package main

import "fmt"

type Pos struct {
	x int
	y int
}

func main() {
	q := 0
	var d rune
	fmt.Scanf("%d %c", &q, &d)

	p := make([]Pos, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&p[i].x, &p[i].y)

	}

	for i := q - 1; i >= 1; i-- {

		p[i] = p[i-1]

	}

	switch d {
	case 'L':
		p[0].x -= 1
	case 'R':
		p[0].x += 1
	case 'U':
		p[0].y -= 1
    case 'D':
		p[0].y += 1
	}

	for i := 0; i < q; i++ {
		fmt.Printf("%d %d\n", p[i].x, p[i].y)
	}

}
