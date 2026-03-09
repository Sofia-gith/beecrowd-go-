package beecrowd

import(
	"fmt"
)

// Numero solitario

func NumeroSolitario(){
	var tamanho int
	var slice []int

	for{
		fmt.Scan(&tamanho)

		if tamanho == 0{
			break
		}


		for i := 0; i < tamanho; i++ {
			var numeros int
			fmt.Scan(&numeros)
			slice = append(slice, numeros)
			

			if i == tamanho-1{
				for j := 0; j < tamanho; j++ {
					if slice[j] == j{
						conta := 0
						conta++

						if conta%2 != 0{
							fmt.Println(slice[j])
						}
					
					}
					if tamanho == 0{
						break
					}
				}
			}

			if tamanho == 0{
				break
			}
		}

		

	
	}
	
}