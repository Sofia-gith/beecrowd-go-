package beecrowd

// Conversão de tempo
import(
	"fmt"
)

type Tempo struct{
	horas 		int
	minutos 	int
	segundos 	int
}

func (t *Tempo) ConverterSegundos(){

	t.horas = t.segundos / 3600
	t.minutos = (t.segundos % 3600) / 60
	t.segundos = t.segundos % 60
}

func (t *Tempo) ImprimirTempo(){
	fmt.Scan( &t.segundos)
	t.ConverterSegundos()
	fmt.Printf("%d:%d:%d\n", t.horas, t.minutos, t.segundos)
}


