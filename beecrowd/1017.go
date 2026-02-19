package beecrowd
// Gasto de combustivel
import(
	"fmt"
)

type ConsumoLitros struct {
    TempoGasto 		int 		
    VelocidadeMedia int 
}

func (c *ConsumoLitros) ConsumoTotal() float64 {
    return float64(c.TempoGasto) * float64(c.VelocidadeMedia) / 12.0
}

//Executavel
func (c *ConsumoLitros) ImprimirConsumo() {
    fmt.Scan(&c.TempoGasto, &c.VelocidadeMedia)
    fmt.Printf("%.3f\n", c.ConsumoTotal())
}
