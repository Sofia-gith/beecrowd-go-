package beecrowd

import (
	"fmt"
	"math"
)

// Distancia Entre Dois Pontos

type Eixos struct{

	distancia float64

	x1 float64
	y1 float64
// -------------------------------------- //
	x2 float64
	y2 float64
}

func (e *Eixos) CalculoDistancia(){
	e.distancia = math.Sqrt(math.Pow(e.x2-e.x1, 2) + math.Pow(e.y2-e.y1, 2))
}

func (e *Eixos) ImprimirDistancia(){
	fmt.Scan(&e.x1, &e.y1, &e.x2, &e.y2)
	e.CalculoDistancia()
	fmt.Printf("%.4f\n", e.distancia)
}