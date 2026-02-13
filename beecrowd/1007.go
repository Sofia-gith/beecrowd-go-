package beecrowd

import(
	"fmt"
)

type Not struct{
	A int
	B int
	C int
	D int
}


//executavel
func Diferenca(){
	var not Not
	fmt.Scan(&not.A, &not.B, &not.C, &not.D)
	fmt.Printf("DIFERENCA = %d\n", (not.A * not.B) - (not.C * not.D))
}