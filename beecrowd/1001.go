package beecrowd

import (
	"fmt"
)

type BasicExtreme struct {
    A int
    B int
}


func basicExtreme(){

	var basicExtreme BasicExtreme
	fmt.Scanf("%d %d", &basicExtreme.A, &basicExtreme.B)
	fmt.Println(basicExtreme.A + basicExtreme.B)

}