package beecrowd

import (
	"fmt"
	"sort"
)


func SortSimples(){
	var num[3]int


	for i:=0; i<3; i++{
		fmt.Scan(&num[i])
	}
	
	ordenado := num
	sort.Ints(ordenado[:])

	
	for _, n := range ordenado{
		fmt.Println(n)
	}

	fmt.Println() // pulando uma linha

	for _, n := range num{
		fmt.Println(n)
	}

}
