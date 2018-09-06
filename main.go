package main

//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
import (
	"fmt"
)

type Matriz struct {
	elementos [][]float64
}

func NewMatriz(v [][]float64) Matriz {
	return Matriz{v}
}

func (m Matriz) Em(i, j int) float64 {
	return m.elementos[i-1][j-1]
}

func (m Matriz) String() string {
	var out string
	for i := 0; i < len(m.elementos); i++ {
		out += "|"
		for j := 0; j < len(m.elementos[i]); j++ {
			out += fmt.Sprintf("%6.2f", m.elementos[i][j]) + " "
		}
		out += "|\n"
	}
	return out
}

func main() {
	valores := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matriz := NewMatriz(valores)
	fmt.Println(matriz)
}
