package beecrowd

import(
	"fmt"
)

// Numero solitario

func NumeroSolitario() {
	for {
		var tamanho int
		contagem := make(map[int]int)
		var num int


		fmt.Scan(&tamanho)

		if tamanho == 0 {
			break
		}

		for i := 0; i < tamanho; i++ {
			
			fmt.Scan(&num)
			contagem[num]++
		}

		for num, count := range contagem {
			if count%2 != 0 {
				fmt.Println(num)
			}
		}
	}
}


//outra versão mais eficiente, utilizando a propriedade XOR
func NumeroSolitario2() {
	for {
		var tamanho int
		fmt.Scan(&tamanho)

		if tamanho == 0 {
			break
		}

		resultado := 0
		for i := 0; i < tamanho; i++ {
			var num int
			fmt.Scan(&num)
			resultado ^= num
		}

		fmt.Println(resultado)
	}
}
