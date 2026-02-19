package beecrowd

// Distancia

import(
	"fmt"
)

type Carros struct{
	min int
	km int
}

func (c *Carros) CalculoDistancia(){
	fmt.Scan(&c.km)
	c.min = c.km * 2
	fmt.Println(c.min, " minutos")
}


