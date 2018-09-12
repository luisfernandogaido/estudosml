package main

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func main() {
	df := dataframe.NewDataFrameFloat64()
	df.NewSerie("x1", []float64{-1, 1, -1, 1, 0, 0, 0, -1, 1, 0, 0, 0, 0, 0.1667})
	df.NewSerie("x2", []float64{-1, -1, 0.6667, 0.6667, -0.4444, -0.7222, 0.6667, -0.1667, -0.1667, -1, 0.9444, -0.1667, 1, -0.1667})
	df.NewSerie("y", []float64{1004, 1636, 852, 1506, 1272, 1270, 1269, 903, 1555, 1260, 1146, 1276, 1225, 1321})
	rm, err := df.NewRegressaoMultipla([]string{"x1", "x2"}, "y")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rm)
}
