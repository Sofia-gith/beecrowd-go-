package beecrowd

import (
	"fmt"
	"math"
)

type Esfera struct{
	Raio float64
	Pi float64
	Volume float64
}

func (e *Esfera) CalculoVolume(){

	e.Pi = 3.14159
	e.Volume = (4.0/3.0) * e.Pi * math.Pow(e.Raio, 3)
}

//executavel
// para executar no main, ler essa função
// no main precisa adicionar "var esfera beecrowd.Esfera" antes de chamar a função
func (e *Esfera) ExecCalculoVolume(){
	fmt.Scan(&e.Raio)
	e.CalculoVolume()
	fmt.Printf("VOLUME = %.3f\n", e.Volume)

}
