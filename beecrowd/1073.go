package beecrowd

import (
	"fmt"
)




func QuadradoPares(){
	var num int
	fmt.Scan(&num)

	if num>5 && num <2000{

		for i:=1; i<=num; i++{
			if i%2==0{
				fmt.Printf("%d^2 = %d\n",i,i*i)
			}
		}
	}
}