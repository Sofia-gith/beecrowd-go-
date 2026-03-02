package beecrowd

import(
	"fmt"
)

func TesteSelecao1(){
	var a, b, c, d int
	fmt.Scan(&a,&b,&c,&d)

	if b>c && d>a && c+d > a+b && c>0 && d>0 && a%2==0{
		fmt.Println("Valores aceitos")
	} else{
		fmt.Println("Valores nao aceitos")
	}
}
