package aprendendo

import (
	"fmt"
)

type No struct{
	valor int
	prox *No
}

func ListaCircularJosephus(){
	n, k := 6 ,2

	head:= &No{valor:1}
	dedo := head
	for i:=2; i<= n; i++{
		dedo.prox = &No{valor:i}
		dedo = dedo.prox
	}
	dedo.prox = head

	dedo = head
	for dedo.prox != dedo{ 
		//condição implica que ira continuar no loop até ter domente 1 elemento
		for i := 0; i<k-2; i++{
			dedo = dedo.prox
		}
		fmt.Println("Eliminado:", dedo.prox.valor)
		dedo.prox = dedo.prox.prox
	}
	fmt.Println("Sobrevivente:",dedo.valor)
}