package beecrowd

import (
	"fmt"

)

//Perseguição em  uma linha perpendicular
// Codigo feito com auxilio de IA


func QuardaCosteira(){
	  var d, vf, vg int
    for {
        _, err := fmt.Scan(&d, &vf, &vg)
        if err != nil {
            break
        }

  
        if vg <= vf {
            fmt.Println("N")
            continue
        }


        if vf*vf*d*d <= 144*(vg*vg-vf*vf) {
            fmt.Println("S")
        } else {
            fmt.Println("N")
        }
    }
}