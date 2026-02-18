package beecrowd

import "fmt"

const Pi = 3.14159

type Metricas struct {
    A float64  // base
    B float64  // Segunda base do trapezio
    C float64 // altura / é o raio tambem
}

func (m *Metricas) CalculaAreas() {
    triangulo := (m.A * m.C) / 2
    circulo := Pi * (m.C * m.C)
    trapezio := ((m.A + m.B) * m.C) / 2
    quadrado := m.B * m.B
    retangulo := m.A * m.B

    fmt.Printf("TRIANGULO: %.3f\n", triangulo)
    fmt.Printf("CIRCULO: %.3f\n", circulo)
    fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
    fmt.Printf("QUADRADO: %.3f\n", quadrado)
    fmt.Printf("RETANGULO: %.3f\n", retangulo)
}

func (m *Metricas) ExecMetricas() {
    fmt.Scan(&m.A, &m.B, &m.C)
    m.CalculaAreas()
}