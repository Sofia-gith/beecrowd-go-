package beecrowd

import (
	"fmt"	
)
func BasicExtreme(){

	var A, B, X int
		 fmt.Scan( &A, &B)
		 X	= A+B
	 	fmt.Println("X =",X)
}

//Segunda forma de fazer: 

type Extreme struct {
    A int
    B int
}

func (e Extreme) Soma() int {
    return e.A + e.B
}

func BasicExtreme2() {
    var e Extreme

    fmt.Scan(&e.A, &e.B)

    fmt.Println("X =", e.Soma())
}