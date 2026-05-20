package main
import "fmt"

type jogador struct {
    pedra [2]int
}

func abs (a, b int) int{
    if a > b {
        return a - b
    } else {
        return b - a
    }
}
func main() {
   n := 0
   fmt.Scan(&n)
    vetor := make([]jogador, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i].pedra[0], &vetor[i].pedra[1])
    }

    
    

    menor := -1
    indice := -1
    for i := 0; i < n; i++ {
        if vetor[i].pedra[0] >= 10 && vetor[i].pedra[1] >= 10 {
            diff := abs(vetor[i].pedra[0], vetor[i].pedra[1])
            if menor == -1 || diff < menor {
                menor = diff
                indice = i
            }
        }
    }
    if indice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Printf("%d\n", indice)
	}
    
    

   
}
