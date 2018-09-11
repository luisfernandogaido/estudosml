//https://matrix.reshish.com/ptBr/gauss-jordanElimination.php
//https://www.khanacademy.org/math/precalculus/precalc-matrices/elementary-matrix-row-operations/a/matrix-row-operations
//https://pt.wikipedia.org/wiki/M%C3%A9todo_dos_m%C3%ADnimos_quadrados
//http://www.portalaction.com.br/analise-de-regressao/22-estimacao-dos-parametros-do-modelo
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/luisfernandogaido/estudosml/matriz"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	valores := make([][]float64, 0)
	lado := 4
	for i := 0; i < lado; i++ {
		linha := make([]float64, 0)
		for j := 0; j < lado; j++ {
			linha = append(linha, float64(rand.Intn(20))-10)
		}
		valores = append(valores, linha)
	}
	m := matriz.New(valores)
	inversa, err := m.Inversa()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	fmt.Println(inversa)
	multiplicada, err := m.Multiplica(inversa)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(multiplicada)
}
