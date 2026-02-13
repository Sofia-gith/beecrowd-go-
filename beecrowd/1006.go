package beecrowd

import(
	"fmt"
)

type Media2 struct{
	A float64
	B float64
	C float64
}
func MediaPonderada(){
	var media Media2
	fmt.Scan(&media.A, &media.B, &media.C)
	fmt.Printf("MEDIA = %.1f\n", ((media.A * 2.0) + (media.B * 3.0) + (media.C * 5.0)) / 10.0)
}