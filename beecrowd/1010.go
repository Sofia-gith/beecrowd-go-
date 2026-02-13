package beecrowd

import(
	"fmt"
)

type Peca struct{
	ID 			int
	Quantidade 	int
	Preco 		float64

	ID2 		int
	Quantidade2 int
	Preco2 		float64
}


func (p Peca) TotalPreco() float64{
	println("entrei aqui")
	return float64(p.Quantidade) * (p.Preco)
	
}

func (p Peca) TotalPreco2() float64{
	println("entrei aqui2")
	return float64(p.Quantidade2) * p.Preco2
}


//execultavel
//para executar no main adicionar:  var peca beecrowd.Peca
func Imprimir(p Peca ) {
    fmt.Scan(&p.ID, &p.Quantidade, &p.Preco)
    fmt.Scan(&p.ID2, &p.Quantidade2, &p.Preco2)

    TotalPrecoTotal := p.TotalPreco() + p.TotalPreco2()
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", TotalPrecoTotal)
}
