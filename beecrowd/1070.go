package beecrowd

import(
	"fmt"
)

func Impares(){
	var num int
	fmt.Scan(&num)

	for i:=0; i<6; i++{

		if num%2!=0{
			fmt.Println(num)
		}else{
			i--
		}
		num++
	}

}