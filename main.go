package main

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func main() {
	df, err := dataframe.NewFloat64CSV("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	df.Roda()
	tv := df.Series["radio"]
	fmt.Println(tv.Maximo, tv.Minimo, tv.Media)
}
