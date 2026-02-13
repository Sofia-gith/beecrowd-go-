package beecrowd

import (
	"fmt"
)

type Soma struct{
	A int
	B int
}

func Somar(){
	var soma Soma
	fmt.Scan(&soma.A, &soma.B)
	fmt.Printf("SOMA = %d\n", soma.A + soma.B)
}
