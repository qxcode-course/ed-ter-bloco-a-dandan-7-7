package main
import "fmt"

func ta(b []byte) bool{
   var pilha []byte
   for i:= 0; i < len(b); i++ {
    char := b[i]
    if  char == '(' || char == '[' || char == '{' {
        pilha = append(pilha, char)
    } else if char == ')' || char == ']' || char == '}' {
        if len(pilha) == 0 {
            return false
        }
        topo := pilha[len(pilha)-1]
        if (char == ')' && topo != '(') || (char == ']' && topo != '[') || (char == '}' && topo != '{') {
            return false
        }
        pilha = pilha[:len(pilha)-1]
    }
   }
   return len(pilha) == 0


}

func main() {
    bal := ""
    fmt.Scan(&bal)

    b := []byte(bal)

    if ta(b){
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}