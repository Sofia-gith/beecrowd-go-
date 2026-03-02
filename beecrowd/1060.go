package beecrowd

import(
	"fmt"
)


func NumPositivo(){
	var positivos int

	for i:=0; i<6; i++{
		var num float64
		fmt.Scan(&num)
		if num>0{
			positivos++
		}
	}
	fmt.Println(positivos, "valores positivos")
}