package beecrowd

import (
	"fmt"
)

type Cedulas struct{
	Dinheiro 		int

}

func (c *Cedulas) ImprimirCedulas(){
	fmt.Scan( &c.Dinheiro)
	fmt.Printf("%d\n", c.Dinheiro)
	fmt.Printf("%d nota(s) de R$ 100,00\n", c.Dinheiro / 100)
	c.Dinheiro %= 100
	fmt.Printf("%d nota(s) de R$ 50,00\n", c.Dinheiro / 50)
	c.Dinheiro %= 50
	fmt.Printf("%d nota(s) de R$ 20,00\n", c.Dinheiro / 20)
	c.Dinheiro %= 20
	fmt.Printf("%d nota(s) de R$ 10,00\n", c.Dinheiro / 10)
	c.Dinheiro %= 10
	fmt.Printf("%d nota(s) de R$ 5,00\n", c.Dinheiro / 5)
	c.Dinheiro %= 5
	fmt.Printf("%d nota(s) de R$ 2,00\n", c.Dinheiro / 2)
	c.Dinheiro %= 2
	fmt.Printf("%d nota(s) de R$ 1,00\n", c.Dinheiro)
}