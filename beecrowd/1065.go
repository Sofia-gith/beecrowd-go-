package beecrowd

import (
	"fmt"
)

func ParesEntreCinco(){
	var Par int
	for i:=0; i<5; i++{
		var num int
		
		fmt.Scan(&num)
		if num%2==0{
			Par++
		}
	}
	fmt.Println(Par, "valores pares")
}