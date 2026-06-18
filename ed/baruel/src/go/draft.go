package main

import "fmt"

func main() {
	qtdttl := 0
	fmt.Scan(&qtdttl)
	qtdbar := 0
	fmt.Scan(&qtdbar)
	fig := make([]int, qtdbar)

	cnt := make(map[int]int)
	for i := 0; i < qtdbar; i++ {
		fmt.Scan((&fig[i]))
		cnt[fig[i]]++
	}

	ver := false
	j := 0
	guar := make([]int, qtdbar)
	for i := 1; i < qtdbar; i++ {
		if fig[i] == fig[i-1] {
			guar[j] = fig[i]

			j++
			ver = true
		}

	}
	for i := 0; i < j; i++ {
		fmt.Print(guar[i])
		if i < j-1 {
			fmt.Print(" ")
		}
	}
	if !ver {
		fmt.Print("N")
	}
	fmt.Println()

	ver = false
	primeiroPrint := true
	for i := 1; i <= qtdttl; i++ {
		if cnt[i] == 0 {
			if !primeiroPrint {
				fmt.Print(" ")
			}
            primeiroPrint = false
			fmt.Print(i)

			ver = true
		}
	}

	if !ver {
		fmt.Print("N")
	}
	fmt.Println()
}
