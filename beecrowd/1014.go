package beecrowd

import (
	"fmt"
)

type Consumo struct {
    X int 		// distância percorrida
    Y float64 	// combustível gasto
}

func (c Consumo) ConsumoMedio() float64 {
    return float64(c.X) / c.Y
}
func (c *Consumo) ImprimirConsumo() {

	fmt.Scan(&c.X, &c.Y)
    fmt.Printf("%.3f km/l\n", c.ConsumoMedio())
}