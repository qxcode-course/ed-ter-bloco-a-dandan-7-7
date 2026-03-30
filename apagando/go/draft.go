package main

import "fmt"

func main() {
	var qini, qdeix int
	fmt.Scan(&qini)
	pessoas := make([]int, qini)

	for i := 0; i < qini; i++ {
		fmt.Scan(&pessoas[i])

	}
	fmt.Scan(&qdeix)
	sairam := make([]int, qdeix)
	for i := 0; i < qdeix; i++ {
		fmt.Scan(&sairam[i])
	}

	for i := 0; i < qdeix; i++ {
		for j := 0; j < qini; j++ {
			if pessoas[j] == sairam[i] {
				pessoas[j] = 0
			}
		}
	}
	for i := 0; i < qini; i++ {
		if pessoas[i] != 0 {
			fmt.Print(pessoas[i], " ")
		}
	}
    fmt.Println()
}
