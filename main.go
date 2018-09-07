//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
//https://slideplayer.com.br/slide/348945/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/luisfernandogaido/estudosml/matriz"
)

func main() {
	t0 := time.Now()
	m := matriz.New([][]float64{
		{1, 2, 2, -1, 2, 1},
		{3, 1, -2, 4, -2, 2},
		{-5, 0, 7, 1, 8, 1},
		{1, 2, 3, 4, -6, -3},
		{2, 2, -1, 3, 5, 6},
		{9, -7, 2, 1, 0, -3},
	})
	det, err := m.Det()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(det)
	fmt.Println(time.Since(t0))
}
