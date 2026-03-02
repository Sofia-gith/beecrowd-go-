package beecrowd

import(
	"fmt"
)

func Intervalo2(){

	var num int	
	var teste int
	var in int
	var out int


	fmt.Scan(&teste)

	for i:=0; i<teste; i++{
		fmt.Scan(&num)
		if num>=10 && num<=20{
			in++
		}else{
			out++
		}
	}

	fmt.Println(in, "in")
	fmt.Println(out, "out")
}