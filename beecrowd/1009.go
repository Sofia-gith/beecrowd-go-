package beecrowd

import(
	"fmt"
)

type SobreVendedor struct{
	Nome string
	SalarioFixo float64
	TotalVendas float64
}

func (s SobreVendedor) TotalSalario() float64{
	return s.SalarioFixo + (s.TotalVendas * 0.15)
}

//executavel
func SobreVendedorTotalSalario(){
	var sobreVendedor SobreVendedor
	fmt.Scan(&sobreVendedor.Nome, &sobreVendedor.SalarioFixo, &sobreVendedor.TotalVendas)
	fmt.Printf("TOTAL = R$ %.2f\n", sobreVendedor.TotalSalario())
}
