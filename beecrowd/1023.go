package beecrowd

//Estiagem


import (
	"fmt"
	"os"
	"sort"
)

type Cidade struct {
	casos 	int
	x 		[]int 
	y 		[]int
}

func (c *Cidade) Ordenar() {

	sort.Ints(c.x)

	for b:=0; b<c.casos; b++{
		fmt.Printf("%d/%d ", c.x[b], c.y[b])
	}
	
}

func Cidades(c *Cidade){

	i := 1
	for {
		fmt.Scanf("%d", &c.casos)
		switch {  
		case c.casos>1:
			c.x = make([]int, c.casos)
			c.y = make([]int, c.casos)
			
				for j:=0; j<c.casos; j++{
					fmt.Scan(&c.x[j], &c.y[j])
					c.y[j] = c.y[j] / c.x[j]
				}
				
	
			fmt.Printf("Cidade# %d: \n", i)


			fmt.Scanf("%d", &c.casos)
			i ++
		case c.casos == 0 : 
			os.Exit(0)
		}
	}
}