package beecrowd

import "fmt"



func SeraTriangulo(){
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a+b>c && a+c>b && b+c>a{
		fmt.Printf("Perimetro = %.1f\n", a+b+c)
	}else{
		fmt.Printf("Area = %.1f\n", (a+b)*c/2)
	}
}