//https://pt.wikipedia.org/wiki/M%C3%A9todo_dos_m%C3%ADnimos_quadrados
//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
package main

import (
	"fmt"

	"github.com/luisfernandogaido/estudosml/matriz"
)

func main() {
	m := matriz.New([][]float64{
		{1, 139, 0.115},
		{1, 126, 0.12},
		{1, 90, 0.105},
		{1, 144, 0.09},
		{1, 163, 0.1},
		{1, 136, 0.12},
		{1, 61, 0.105},
		{1, 62, 0.08},
		{1, 41, 10},
		{1, 120, 0.115},
	})
	fmt.Println(m)
	m.TrocaLinhas(0, 1)
	fmt.Println(m)
}
