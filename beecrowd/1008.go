package beecrowd

import(
	"fmt"
)

type Horas struct{
	A int
	B int
	C float64
}

//função de Conta separada
func (H Horas) ContaHtrabalhadas() float64{
	return H.C * float64(H.B)
}


//executavel
func HorasTrabalhadas(){
	var horas Horas
	fmt.Scan(&horas.A, &horas.B, &horas.C)

	fmt.Printf("NUMBER = %d\n", horas.A)
	fmt.Printf("SALARY = U$ %.2f\n", horas.ContaHtrabalhadas())
}