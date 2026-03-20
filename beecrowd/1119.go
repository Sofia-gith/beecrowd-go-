package beecrowd

import "fmt"

type Circular struct {
	valor int
	prox  *Circular
	ant   *Circular
}

func Rinocerontes() {
	var N, k, m int
	for {
		fmt.Scan(&N, &k, &m)
		if N == 0 && k == 0 && m == 0 {
			break
		}

		cabeca := &Circular{valor: 1}
		dedo := cabeca
		for i := 2; i <= N; i++ {
			dedo.prox = &Circular{valor: i, ant: dedo}
			dedo = dedo.prox
		}
		dedo.prox = cabeca
		cabeca.ant = dedo

		dedoH := dedo   // nó N — predecessor clockwise do nó 1
		dedoA := cabeca // nó 1 — predecessor anti-clockwise do nó N

		tamanho := N
		primeiro := true

		for tamanho > 0 {
			for i := 0; i < k-1; i++ {
				dedoH = dedoH.prox
			}
			vitH := dedoH.prox

			for i := 0; i < m-1; i++ {
				dedoA = dedoA.ant
			}
			vitA := dedoA.ant

			if !primeiro {
				fmt.Print(",")
			}
			primeiro = false

			if vitH == vitA {
				fmt.Printf("%3d", vitH.valor)
				vitH.ant.prox = vitH.prox
				vitH.prox.ant = vitH.ant
				dedoH = vitH.ant
				dedoA = vitH.prox
				tamanho--
			} else {
				fmt.Printf("%3d%3d", vitH.valor, vitA.valor)
				if dedoA == vitH {
					dedoA = vitH.prox
				}
				vitH.ant.prox = vitH.prox
				vitH.prox.ant = vitH.ant
				dedoH = vitH.ant
				tamanho--
				if dedoH == vitA {
					dedoH = vitA.ant
				}
				vitA.ant.prox = vitA.prox
				vitA.prox.ant = vitA.ant
				dedoA = vitA.prox
				tamanho--
			}
		}
		fmt.Println()
	}
}