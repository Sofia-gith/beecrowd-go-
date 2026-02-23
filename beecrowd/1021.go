package beecrowd

import "fmt"

type Note struct {
    Value int
}

func (n *Note) ImprimirNota() {
    var input float64
    fmt.Scanf("%f", &input)
    n.Value = int(input * 100) 

    remainder := n.Value

    fmt.Printf("NOTAS:\n")
    fmt.Printf("%d nota(s) de R$ 100.00\n", remainder/10000); remainder %= 10000
    fmt.Printf("%d nota(s) de R$ 50.00\n",  remainder/5000);  remainder %= 5000
    fmt.Printf("%d nota(s) de R$ 20.00\n",  remainder/2000);  remainder %= 2000
    fmt.Printf("%d nota(s) de R$ 10.00\n",  remainder/1000);  remainder %= 1000
    fmt.Printf("%d nota(s) de R$ 5.00\n",   remainder/500);   remainder %= 500
    fmt.Printf("%d nota(s) de R$ 2.00\n",   remainder/200);   remainder %= 200

    fmt.Printf("MOEDAS:\n")
    fmt.Printf("%d moeda(s) de R$ 1.00\n",  remainder/100);   remainder %= 100
    fmt.Printf("%d moeda(s) de R$ 0.50\n",  remainder/50);    remainder %= 50
    fmt.Printf("%d moeda(s) de R$ 0.25\n",  remainder/25);    remainder %= 25
    fmt.Printf("%d moeda(s) de R$ 0.10\n",  remainder/10);    remainder %= 10
    fmt.Printf("%d moeda(s) de R$ 0.05\n",  remainder/5);     remainder %= 5
    fmt.Printf("%d moeda(s) de R$ 0.01\n",  remainder/1)
}