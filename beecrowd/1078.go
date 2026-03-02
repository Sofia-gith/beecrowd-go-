package beecrowd

import(
	"fmt"
)


func Tabuada(){
	var num int
	fmt.Scan(&num)

	for i:=1; i<=10; i++{
		fmt.Printf("%d x %d = %d\n",i,num,num*i)
	}

}