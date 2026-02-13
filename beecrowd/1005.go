package beecrowd

import(
	"fmt"
)

type Nota struct{
	A float64
	B float64
}
func Media(){
	var nota Nota
	fmt.Scan(&nota.A, &nota.B)
	fmt.Printf("MEDIA = %.5f\n", ((nota.A * 3.5) + (nota.B * 7.5)) / 11)
}
