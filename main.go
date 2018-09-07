//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
//https://slideplayer.com.br/slide/348945/
package main

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/matriz"
)

func main() {
	a := matriz.New([][]float64{
		{2, 5, 9},
		{3, 6, 8},
	})
	b := matriz.New([][]float64{
		{2, 7},
		{4, 3},
		{5, 2},
	})
	c, err := a.Multiplica(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
