package beecrowd

import (
	"fmt"
)

//Outro metodo utilizando matematica modular, mais elegante ainda

func TempoJogoModular(){
		var inicio, fim int
	fmt.Scan(&inicio, &fim)

	duracao := (fim - inicio + 24) % 24

	if duracao == 0 {
		duracao = 24
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duracao)
}

//outro meio de resolução bem mais elegante
func Outro() {
	var inicio, fim int
	fmt.Scan(&inicio, &fim)

	var duracao int

	if inicio < fim {
		duracao = fim - inicio
	} else if inicio > fim {
		duracao = (24 - inicio) + fim
	} else {
		duracao = 24
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duracao)
}


//desenvolvido por mim
func JogoDoTempo(){

	dia := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
		12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
		22, 23,
	}

	var inicio, fim int
	fmt.Scan(&inicio, &fim)

	if inicio == fim {
		fmt.Println("O JOGO DUROU 24 HORA(S)")
		return
	}

	duracao := 0
	pos := inicio

	for {
		
		pos = (pos + 1) % 24
		duracao++

		if dia[pos] == fim {
			break
		}
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duracao)
}
