package main

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func main() {
	df, err := dataframe.NewDataFrameFloat64CSV("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	df.Roda()
	for _, s := range df.Series {
		fmt.Println(s.Nome, s.Q1, s.Q2, s.Q3)
	}
}
