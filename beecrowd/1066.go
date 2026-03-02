package beecrowd

import(
	"fmt"
)


func ParesImparesPosNega(){


	var Par 	int
	var Impar 	int
	var Pos 	int
	var Nega 	int


	for i:=0; i<5; i++{
		var num int
		fmt.Scan(&num)
		if num%2==0{
			Par++
		}  
		if num%2!=0{
			Impar++
		}  
		if num>0{
			Pos++
		}  
		if num<0{
			Nega++
		}
	}
	fmt.Println(Par, "valor(es) par(es)")
	fmt.Println(Impar, "valor(es) impar(es)")
	fmt.Println(Pos, "valor(es) positivo(s)")
	fmt.Println(Nega, "valor(es) negativo(s)")
}
