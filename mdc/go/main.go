package main

import (
	"fmt"
)

func mdc(a, b, div, ult int) int {
	
	var fin int
	if(div < a && div < b){
		if(a % div == 0 && b % div == 0){
			ult = div
			
			fin = mdc(a, b, div+1, ult)
		} else{
			fin = mdc(a, b, div+1, ult)
		}
	} else{
		return ult
	}
	return fin
	
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b, 1, 1))
}
