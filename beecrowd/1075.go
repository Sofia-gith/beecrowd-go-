package beecrowd

import (
	"fmt"
)



func Resto2(){
	var N int
	fmt.Scan(&N)
	for i := 0; i < 10000; i++ {
		if i % N == 2 {
			fmt.Println(i)
		}
	}
}