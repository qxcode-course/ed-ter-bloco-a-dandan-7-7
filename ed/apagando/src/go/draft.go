package main
import "fmt"

func remover(p, sai []int) []int {
	var res []int

	for i := 0; i < len(sai); i++ {
		for j := 0; j < len(p); j++ {
			if sai[i] == p[j] {
				p[j] = 0
			}
		}
	}

	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			res = append(res, p[i])
		}
	}

	return res


}



func main() {
	n:=0
	fmt.Scan(&n)
	var pessoas []int

	for i:=0; i<n; i++ {
		var p int
		fmt.Scan(&p)
		pessoas = append(pessoas, p)
	}
	sairam:= 0
	fmt.Scan(&sairam)
	var sairamPessoas []int
	for i:=0; i<sairam; i++ {
		var p int
		fmt.Scan(&p)
		sairamPessoas = append(sairamPessoas, p)
	}
	pessoas = remover(pessoas, sairamPessoas)

	for i := 0; i < len(pessoas); i++ {
		fmt.Print(pessoas[i]," " )
	}
	fmt.Println()
}
