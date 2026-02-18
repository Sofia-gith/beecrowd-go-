package beecrowd

import "fmt"

type Omaior struct {
    A, B, C int
}

// deixa o valor positivo se for negativo
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func maiorEntreDois(a, b int) int {
    return (a + b + abs(a-b)) / 2
}

func (o *Omaior) ImprimirOmaior() {
    fmt.Scan(&o.A, &o.B, &o.C)

    maiorAB := maiorEntreDois(o.A, o.B)   // maior entre A e B
    resultado := maiorEntreDois(maiorAB, o.C) // maior entre o resultado anterior e C

    fmt.Printf("%d eh o maior\n", resultado)
}