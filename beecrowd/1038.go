package beecrowd

import(
	"fmt"
)

func Lanche(){
	precos := map[int]float64{
		1: 4.00,
		2: 4.50,
		3: 5.00,
		4: 2.00,
		5: 1.50,
	}

	var codigo, quantidade int
	fmt.Scan(&codigo, &quantidade)

	total:= precos[codigo] * float64(quantidade)
	fmt.Printf("Total: R$ %.2f\n", total)
}
