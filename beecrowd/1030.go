package beecrowd

import (
	"fmt"
)


type No struct{
	valor int
	prox *No
}

func Josephus(){
	var teste, n, k, contar int
	fmt.Scan(&teste)
	contar = 1

	for i := 0; i < teste; i++ {
		
		fmt.Scan(&n, &k)

		cabeca := &No{valor:1}
		dedo := cabeca
		for i:=2; i<= n; i++{
			dedo.prox = &No{valor:i}
			dedo = dedo.prox
		}
		
		dedo.prox = cabeca

		for dedo.prox != dedo{
			
			for i := 0; i<k-1; i++{
				
				dedo = dedo.prox
			}
			dedo.prox = dedo.prox.prox
		}
		 fmt.Printf("Case %d: %d\n", contar, dedo.valor)
		contar++
	}
}