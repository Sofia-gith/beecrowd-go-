package beecrowd
import (
	"fmt"	
)

func AreaCircle(){
	var n, r float64
	n = 3.14159
	fmt.Scan(&r)
	var area = n * r * r
	fmt.Printf("A=%.4f\n", area)

	

}