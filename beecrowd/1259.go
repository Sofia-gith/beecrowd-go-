package beecrowd

import (
	"fmt"
	"sort"
)

func PareseImpares1259(){
	var EntradaInicial  int
	var List int
	var Pares []int
	var Impares []int

	fmt.Scan(&EntradaInicial)

	for i := 0; i < EntradaInicial; i++{
		
		fmt.Scan(&List)

		if List%2 == 0{
			Pares = append(Pares, List)
		}else{
			Impares = append(Impares, List)
		}
	}

	sort.Ints(Pares)
	sort.Sort(sort.Reverse(sort.IntSlice(Impares)))

	for i := 0; i < len(Pares); i++ {
		fmt.Println(Pares[i])
	}

	for i := 0; i < len(Impares); i++ {
		fmt.Println(Impares[i])
	}

}