package beecrowd

import (

	"fmt"
)

// Idade em dias

type Idade struct{
	dias int
}

func (i *Idade) ImprimirIdade(){
	fmt.Scan( &i.dias)
	fmt.Printf("%d ano(s)\n", i.dias / 365)
	fmt.Printf("%d mes(es)\n", (i.dias % 365)/30)
	fmt.Printf("%d dia(s)\n", (i.dias % 365) % 30)
	
}