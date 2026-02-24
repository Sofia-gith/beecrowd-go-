package beecrowd

import "fmt"

type Racional struct {
	N1, D1  int
	N2, D2  int
	operador string
}

func mdc(a, b int) int {
	if a < 0 {
		a = -a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (r *Racional) ContaComRacionais() (int, int) {
	var numRes, denRes int

	switch r.operador {
	case "+":
		numRes = r.N1*r.D2 + r.N2*r.D1
		denRes = r.D1 * r.D2
	case "-":
		numRes = r.N1*r.D2 - r.N2*r.D1
		denRes = r.D1 * r.D2
	case "*":
		numRes = r.N1 * r.N2
		denRes = r.D1 * r.D2
	case "/":
		numRes = r.N1 * r.D2
		denRes = r.N2 * r.D1
	}

	return numRes, denRes
}

//executavel

func (r *Racional) ImprimirRacionais() {

	fmt.Scanf("%d / %d %s %d / %d", &r.N1, &r.D1, &r.operador, &r.N2, &r.D2)

	numRes, denRes := r.ContaComRacionais()


	divisor := mdc(numRes, denRes)
	numSimp := numRes / divisor
	denSimp := denRes / divisor

	fmt.Printf("%d/%d = %d/%d\n", numRes, denRes, numSimp, denSimp)
}

func Exercicio1022() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		r := &Racional{}
		r.ImprimirRacionais()
	}
}