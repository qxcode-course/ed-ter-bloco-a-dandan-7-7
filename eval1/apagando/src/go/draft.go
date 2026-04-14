package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var pessoas = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &pessoas[i])
	}
	var m int
	fmt.Scanf("%d", &m)
	var deixaram = make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &deixaram[i])
		
	}
	var fin = make([]int, n-m)
	k := 0
	
	
	for i := 0; i < m; i++{
		
		for j:= 0; j < n; j++{
			if pessoas[j] == deixaram[i]{
				pessoas[j] = -1
			
				break
			}
		}
		
	}
	
	for i := 0; i < n - m; i++ {
			if(pessoas[i] != -1){
				fin[k] = pessoas[i]
				k++
			}
	}
	for i := 0; i < n - m; i++ {
		fmt.Printf("%d ", fin[i])
	}
		
	fmt.Println()
}
