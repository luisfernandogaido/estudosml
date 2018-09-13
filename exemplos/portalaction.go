package exemplos

import (
	"fmt"
	"log"

	"github.com/luisfernandogaido/estudosml/dataframe"
)

func Ataques() {

	//exemplos-regressao.pdf, página 6
	df, err := dataframe.NewDataFrameFloat64CSV("./data/ataques.csv")
	if err != nil {
		log.Fatal(err)
	}
	r, err := df.NewRegressaoMultipla([]string{"ataques", "duracao"}, "indice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Formula(3))
}
