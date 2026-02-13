package beecrowd

import(
	"fmt"
)

type Prod struct{
	A int
	B int
}
func Produto(){
	var prod Prod
	fmt.Scan(&prod.A, &prod.B)
	fmt.Printf("PROD = %d\n", prod.A * prod.B)
}