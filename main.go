//https://www.khanacademy.org/math/precalculus/precalc-matrices/elementary-matrix-row-operations/a/matrix-row-operations
//https://pt.wikipedia.org/wiki/M%C3%A9todo_dos_m%C3%ADnimos_quadrados
//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
package main

import (
	"fmt"

	"github.com/luisfernandogaido/estudosml/matriz"
)

func main() {
	m := matriz.New([][]float64{
		{-5, 7, 3},
		{-2, -1, 4},
		{8, 8, -6},
	})
	m.MultiplicaLinha(3, 2)
	m.SomaLinhas(1, 1, 3)
	fmt.Println(m)
}
