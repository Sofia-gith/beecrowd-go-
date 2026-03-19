package beecrowd

import (
	"fmt"
)

type Circular struct{
	valor int
	prox *Circular
}

func Rinocerontes(){
	var Roda, Horario, AntiHorario int
	fmt.Scan(&Roda)


	fmt.Scan(&Horario, &AntiHorario)

	cabeca:= &Circular{valor:1}
	dedo := cabeca

	for i:=2; i<=Roda; i++{
		dedo.prox = &Circular{valor:i}
		dedo = dedo.prox
		
	}
	dedo.prox= cabeca

	for dedo.prox != dedo{
		for i:=0; i<Horario-1; i++{
			dedo = dedo.prox
			fmt.Println("Eliminado sentido horario:", dedo.prox.valor)
		}	
		fmt.Println("Eliminado sentido horario definitivo:", dedo.prox.valor)

		dedo.prox = dedo.prox.prox

		
		for j:=0; j<AntiHorario+1; j++{
			dedo = dedo.prox
			fmt.Println("Eliminado sentido antiHorario:", dedo.prox.valor)
		}	
		fmt.Println("Eliminado sentido horario definitivo:", dedo.prox.valor)
		dedo.prox = dedo.prox.prox

	}
	fmt.Println("fim:", dedo.prox.valor)
	
}