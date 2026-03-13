package main
import (
    "fmt"
    "math"
)

func main() {
    var a, b, c, p, v1, fin float64
    fmt.Scanf("%f", &a)
    fmt.Scanf("%f", &b)
    fmt.Scanf("%f", &c)
    p = (a + b + c ) / 2
    v1 = p * (p - a) * (p - b) * (p - c)
    fin = math.Sqrt(v1)
    fmt.Printf("%.2f\n", fin)

}
