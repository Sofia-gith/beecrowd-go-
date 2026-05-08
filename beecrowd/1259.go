package beecrowd

import (
	"fmt"
	"sort"
)

func PareseImpares1259(){

	// Função desenvolvida por mim

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


func PareseImpares1259Melhorado(){

	// Função mais eficiente

	var n int
	var numero int

	var numeros []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Scan(&numero)

		numeros = append(numeros, numero)
	}

	sort.Slice(numeros, func(i, j int) bool {

		// Se o primeiro for par e o segundo ímpar
		if numeros[i]%2 == 0 && numeros[j]%2 != 0 {
			return true
		}

		// Se o primeiro for ímpar e o segundo par
		if numeros[i]%2 != 0 && numeros[j]%2 == 0 {
			return false
		}

		// Se ambos forem pares -> crescente
		if numeros[i]%2 == 0 {
			return numeros[i] < numeros[j]
		}

		// Se ambos forem ímpares -> decrescente
		return numeros[i] > numeros[j]
	})

	for _, valor := range numeros {
		fmt.Println(valor)
	}
}
