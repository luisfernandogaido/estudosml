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
	amplitude := 9
	lado := 5
	valores := make([][]float64, 0)
	for i := 0; i < lado; i++ {
		linha := make([]float64, 0)
		for j := 0; j < lado; j++ {
			n := rand.Intn(2*amplitude-1) - amplitude
			linha = append(linha, float64(n))
		}
		valores = append(valores, linha)
	}
	ma := matriz.New(valores)
	fmt.Println(ma)
	t0 := time.Now()
	det, err := ma.Det()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.0f\n", det)
	fmt.Println(time.Since(t0))
}
